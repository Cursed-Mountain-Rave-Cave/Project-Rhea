package main

import (
	"log"
	"net"
	"os"

	"./storage"
	"./web"
)

var users = storage.NewUsers()

func handleLogin(c *web.Connection, login web.Login) error {
	err := users.Login(c, login)
	if err != nil {
		return c.SendResponse(web.NewError(err.Error()))
	}
	return c.SendResponse(web.NewInfo("You have been successfully logged in"))
}

func handleRegister(c *web.Connection, register web.Register) error {
	err := users.Register(c, register)
	if err != nil {
		return c.SendResponse(web.NewError(err.Error()))
	}
	return c.SendResponse(web.NewInfo("You have been successfully registered"))
}

func handleSendAll(c *web.Connection, sendAll web.SendAll) error {
	err := users.SendAll(c, sendAll)
	if err != nil {
		return c.SendResponse(web.NewError(err.Error()))
	}
	return nil
}

func handleRequest(c *web.Connection, r web.Request) error {
	switch r.Type {
	case "login":
		{
			login, err := web.UnwrapLogin(r.Data)
			if err != nil {
				c.SendResponse(web.NewError(err.Error()))
				return err
			}
			return handleLogin(c, login)
		}
	case "register":
		{
			register, err := web.UnwrapRegister(r.Data)
			if err != nil {
				c.SendResponse(web.NewError(err.Error()))
				return err
			}
			return handleRegister(c, register)
		}
	case "send_all":
		{
			sendAll, err := web.UnwrapSendAll(r.Data)
			if err != nil {
				c.SendResponse(web.NewError(err.Error()))
				return err
			}
			return handleSendAll(c, sendAll)
		}
	}
	return nil
}

func handleConnection(c *web.Connection) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	defer log.Printf("Stop serving %s\n", c.RemoteAddr().String())
	defer func() {
		users.Logout(c.GetLogin())
		c.Close()
	}()
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

	log.Println("Opyat' rabota(")

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(web.NewConnection(c))
	}
}
