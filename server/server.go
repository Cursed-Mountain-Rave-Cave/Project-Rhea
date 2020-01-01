package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"

	"./utils"
	"./web"
)

func encode(login, password string) []byte {
	user := web.Login{Login: login, Password: password}
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(userJSON))
	msg := web.Message{Type: "login", Data: string(userJSON)}
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(msgJSON))
	return msgJSON
}

func decode(userJSON []byte) {
	var msg web.Message
	err := json.Unmarshal(userJSON, &msg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(msg)

	var user web.Login
	err = json.Unmarshal([]byte(msg.Data), &user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(user)

}

func handleConnection(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()
	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)
	for {
		netData, err := utils.ReadJSON(reader)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Get", string(netData), "from", c.RemoteAddr().String())
		decode(netData)

		var msg = encode("hello", "there")
		log.Println("Send", string(msg), "to", c.RemoteAddr().String())

		_, err = writer.Write(msg)
		err = writer.Flush()
		if err != nil {
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
		go handleConnection(c)
	}
}
