package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
)

var (
	random = rand.New(rand.NewSource(0))
)

//Message ss
type Message struct {
	MessageType string
	Message     string
}

//Login ss
type Login struct {
	Login    string
	Password string
}

/*
func encode(login, password string) []byte {
	user := Login{login, password}
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(userJSON))
	msg := Message{"login", userJSON}
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(msgJSON))
	return msgJSON
}*/

func decode(userJSON []byte) {
	var msg Message
	err := json.Unmarshal(userJSON, &msg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(msg)

	var user Login
	err = json.Unmarshal([]byte(msg.Message), &user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(user)

}

func handleConnection(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		netData := []byte{}
		balance := 0
		b, err := reader.ReadByte()
		if err != nil && b != byte('{') {
			log.Println(err)
			return
		}

		netData = append(netData, b)
		balance++
		for balance > 0 {
			b, err := reader.ReadByte()
			if err != nil {
				log.Println(err)
				return
			}
			netData = append(netData, b)

			switch b {
			case byte('{'):
				balance++
			case byte('}'):
				balance--
			}
		}

		temp := strings.TrimSpace(string(netData))
		log.Println("Get", temp, "from", c.RemoteAddr().String())
		decode([]byte(temp))
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
		go handleConnection(c)
	}
}
