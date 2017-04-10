// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hobeone/nntp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gospotgo",
	Short: "Usenet indexer written in go",
	Long: `gospotgo aims to be a spotweb compatible (or better) usenet indexer
to allow search and downloading of nzb files.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, port := viper.GetString("host"), viper.GetString("port")
		login, password := viper.GetString("login"), viper.GetString("password")

		// connect to news server
		conn, err := nntp.NewTLS("tcp", host+":"+port, nil)
		if err != nil {
			log.Fatalf("connection failed: %v", err)
		}

		// auth
		if err := conn.Authenticate(login, password); err != nil {
			log.Fatalf("Could not authenticate")
		}

		// connect to a news group
		grpName := "alt.binaries.pictures"
		group, err := conn.Group(grpName)
		if err != nil {
			log.Fatalf("Could not connect to group %s: %v", grpName, err)
		}

		fmt.Printf("%+v", group)

		// fetch an article
		// id := "<4c1c18ec$0$8490$c3e8da3@news.astraweb.com>"
		// article, err := conn.Article(id)
		// if err != nil {
		// 	log.Fatalf("Could not fetch article %s: %v", id, err)
		// }

		// read the article contents
		//body := strings.Join(article.Body, "")

		// newsAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
		// if err != nil {
		// 	println("Could not resolve", host, "on", port)
		// 	os.Exit(1)
		// }
		// conn, err := net.DialTCP("tcp", nil, newsAddr)
		// if err != nil {
		// 	println("Could not connect to", host, "on", port)
		// 	os.Exit(1)
		// }
		// defer conn.Close()
		// fmt.Println("Connected")
		// connbuf := bufio.NewReader(conn)
		// str, err := connbuf.ReadString('\n')

		// fmt.Println("Sending authinfo")
		// conn.Write([]byte("authinfo\n"))
		// str, err = connbuf.ReadString('\n')
		// fmt.Println(str)

		// fmt.Println("Sending username: ", login)
		// conn.Write([]byte("authinfo user " + login + "\n"))
		// str, err = connbuf.ReadString('\n')
		// fmt.Println(str)

		// fmt.Println("Sending password")
		// conn.Write([]byte("authinfo pass " + password + "\n"))
		// str, err = connbuf.ReadString('\n')
		// fmt.Println(str)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gospotgo.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gospotgo") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error using config file:", viper.ConfigFileUsed(), err)
	}
}
