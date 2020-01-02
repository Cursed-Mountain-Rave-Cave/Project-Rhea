package storage

import (
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
	Users   map[string]User
	rwmutex sync.RWMutex
}

var users Users

func Login(web.Login) error {

	return nil
}

func Register(web.Register) error {

	return nil
}
