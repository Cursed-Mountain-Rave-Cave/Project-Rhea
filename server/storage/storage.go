package storage

import (
	"fmt"
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

func NewUsers() *Users {
	return &Users{users: make(map[string]*User)}
}

/*
//Login logs in user to users
func (users *Users) Login(c *web.Connection, form web.Login) error {
	users.rwmutex.Lock()
	defer users.rwmutex.Unlock()
	login, password := form.Login, form.Password
	user, registered := users.users[login]
	//make good login handler
	if !registered {
		return fmt.Errorf("User with login \"%s\" do not exist.", login)
	}
	if user.connection != nil {
		return fmt.Errorf("User with login \"%s\" already online", login)
	}
	return nil
}*/

//Register adds new user to users
func (users *Users) Register(c *web.Connection, form web.Register) error {
	users.rwmutex.Lock()
	defer users.rwmutex.Unlock()
	login, password := form.Login, form.Password
	_, registered := users.users[login]
	if registered {
		return fmt.Errorf("User with login \"%s\" already exist.", login)
	}
	users.users[login] = &User{login, password, nil}
	return nil
}
