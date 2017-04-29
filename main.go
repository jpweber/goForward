package main

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

type netParams struct {
	listenPort string
	destHost   string
	destPort   string
}

func main() {

	app := cli.NewApp()
	app.Name = "goProxy"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "listen, l",
			Usage:  "Posrt to listen on",
			EnvVar: "LISTEN_PORT",
		},
		cli.StringFlag{
			Name:   "dest-host, dh",
			Usage:  "Host to forward traffic to",
			EnvVar: "DEST_HOST",
		},
		cli.StringFlag{
			Name:   "dest-port, dp",
			Usage:  "Port to forward traffit to",
			EnvVar: "DEST_PORT",
		},
	}

	app.Action = func(c *cli.Context) error {
		log.Println("Port Forwarder Started")
		p := netParams{
			listenPort: c.String("listen"),
			destHost:   c.String("dest-host"),
			destPort:   c.String("dest-port"),
		}

		ln, err := net.Listen("tcp", ":"+p.listenPort)
		if err != nil {
			panic(err)
		}

		for {
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}

			go handleRequest(conn, &p)
		}
	}

	app.Run(os.Args)

}

func handleRequest(conn net.Conn, params *netParams) {
	log.Println("Initiating New Connection")

	proxy, err := net.Dial("tcp", params.destHost+":"+params.destPort)
	if err != nil {
		panic(err)
	}

	log.Println("Connected Successfully")
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
