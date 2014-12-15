package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Spot Go"
	app.Usage = "Index your usenet with Go"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "login",
			Value: "",
			Usage: "username for your news server",
		},
		cli.StringFlag{
			Name: "password",
			Value: "",
			Usage: "password for your news server",
		},
		cli.StringFlag{
			Name: "ip",
			Value: "news.astraweb.com",
			Usage: "address of your news server",
		},
		cli.StringFlag{
			Name: "port",
			Value: "119",
			Usage: "port number of your news server",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "login",
			Usage: "login to your news server then quit",
			Action: func(c *cli.Context) {
				conn := login(c)
				defer conn.Close()
			},
		},
	}
	app.Run(os.Args)
}

func login(c *cli.Context) (conn *net.TCPConn) {
	ip, port := c.GlobalString("ip"), c.GlobalString("port")

	newsAddr, err := net.ResolveTCPAddr("tcp", ip + ":" + port)
	if err != nil {
		println("Could not resolve", ip, "on", port)
		os.Exit(1)
	}

	conn, err = net.DialTCP("tcp", nil, newsAddr)
	defer conn.Close()
	if err != nil {
		println("Could not connect to", ip, "on", port)
		os.Exit(1)
	}
	fmt.Println("Connected")
	connbuf := bufio.NewReader(conn)
	str, err := connbuf.ReadString('\n')
	login, password := c.GlobalString("login"), c.GlobalString("password")

	fmt.Println("Sending authinfo")
	conn.Write([]byte("authinfo\n"))
	str, err = connbuf.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Sending username: ", login)
	conn.Write([]byte("authinfo user " + login + "\n"))
	str, err = connbuf.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Sending password")
	conn.Write([]byte("authinfo pass " + password + "\n"))
	str, err = connbuf.ReadString('\n')
	fmt.Println(str)

	return conn
}
