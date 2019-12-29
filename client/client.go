package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

//Message ss
type Message struct {
	MessageType string
	Message     []byte
}

//Login ss
type Login struct {
	Login    string
	Password string
}

func encode(login, password string) []byte {
	user := Login{login, password}
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(userJSON))
	msg := Message{"login", userJSON}
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(msgJSON))
	return msgJSON
}

func decode(userJSON []byte) {
	var msg Message
	err := json.Unmarshal(userJSON, &msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msg)

	var user Login
	err = json.Unmarshal(msg.Message, &user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user.Login)

}

func main() {
	login := ""
	password := ""
	c, err := net.Dial("tcp4", "127.0.0.1:25565")
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	for {
		fmt.Scan(&login, &password)
		_, err := c.Write(encode(login, password))
		if err != nil {
			log.Fatalln(err)
		}
		_, err = c.Write([]byte("\n"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Successfully send", Login{login, password}, "to", c.RemoteAddr().String())
	}
}
