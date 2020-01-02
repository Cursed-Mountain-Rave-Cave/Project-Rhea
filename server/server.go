package main

import (
	"encoding/json"
	"log"
	"net"
	"os"

	"./web"
)

func handleLogin(c *web.Connection, login web.Login) error {
	return c.SendResponse(web.NewInfo("You have been successfully logged in"))
}

func handleRegister(c *web.Connection, register web.Register) error {
	return c.SendResponse(web.NewInfo("You have been successfully registered"))
}

func handleSendAll(c *web.Connection, sendAll web.SendAll) error {
	return nil
}

func handleRequest(c *web.Connection, r web.Request) error {
	switch r.Type {
	case "login":
		{
			var login web.Login
			err := json.Unmarshal([]byte(r.Data), &login)
			if err != nil {
				c.SendResponse(web.NewError(err.Error()))
				return err
			}
			return handleLogin(c, login)
		}
	case "register":
		{
			var register web.Register
			err := json.Unmarshal([]byte(r.Data), &register)
			if err != nil {
				c.SendResponse(web.NewError(err.Error()))
				return err
			}
			return handleRegister(c, register)
		}
	case "send_all":
		{

		}
	}
	return nil
}

func handleConnection(c *web.Connection) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()
	defer log.Printf("Stop serving %s\n", c.RemoteAddr().String())
	for {
		request, err := c.ReceiveRequest()
		if err != nil {
			log.Println(err)
			return
		}
		if handleRequest(c, request) != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Plese provide a port number!")
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(web.NewConnection(c))
	}
}
