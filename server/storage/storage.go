package storage

import (
	"fmt"
	"log"
	"sync"

	"../web"
)

//User struct
type User struct {
	login      string
	password   string
	connection *web.Connection
}

//Users struct
type Users struct {
	users   map[string]*User
	rwmutex sync.RWMutex
}

//NewUsers creates new Users structure
func NewUsers() *Users {
	return &Users{users: make(map[string]*User)}
}

//Login logs in user to users
func (users *Users) Login(c *web.Connection, form web.Login) error {
	users.rwmutex.Lock()
	defer users.rwmutex.Unlock()
	login, password := form.Login, form.Password
	user, registered := users.users[login]
	if !registered {
		return fmt.Errorf("User with login \"%s\" do not exist", login)
	}
	if user.password != password {
		return fmt.Errorf("Wrong password fo user with login \"%s\"", login)
	}
	if user.connection != nil {
		return fmt.Errorf("User with login \"%s\" already online", login)
	}

	if c.GetLogin() != "" {
		users.Logout(c.GetLogin())
	}
	user.connection = c
	c.SetLogin(login)

	return nil
}

//Register adds new user to users
func (users *Users) Register(c *web.Connection, form web.Register) error {
	users.rwmutex.Lock()
	defer users.rwmutex.Unlock()
	login, password := form.Login, form.Password
	_, registered := users.users[login]
	if registered {
		return fmt.Errorf("User with login \"%s\" already exist", login)
	}
	users.users[login] = &User{login, password, nil}
	return nil
}

//SendAll sends message to all connected users
func (users *Users) SendAll(c *web.Connection, msg web.SendAll) error {
	users.rwmutex.Lock()
	defer users.rwmutex.Unlock()

	if c.GetLogin() == "" {
		return fmt.Errorf("You must log in to send messages")
	}
	receiveAll := web.ReceiveAll{Login: c.GetLogin(), Message: msg.Message}
	response := web.Response{Type: "receive_all", Data: receiveAll.String()}
	for _, user := range users.users {
		if user.connection != nil {
			user.connection.SendResponse(response)
		}
	}
	return nil
}

//Logout makes user logged out from server
func (users *Users) Logout(login string) {
	oldUser, registered := users.users[login]
	if registered {
		log.Printf("\"%s\" logs out", login)
		oldUser.connection.SetLogin("")
		oldUser.connection = nil
	}
}
