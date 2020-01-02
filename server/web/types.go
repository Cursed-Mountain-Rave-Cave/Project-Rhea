package web

import (
	"encoding/json"
	"log"
)

//Request structure
type Request struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//String return Request string representation
func (r Request) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//Response structure
type Response struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//String return Response string representation
func (r Response) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//Info response structure
type Info struct {
	Info string `json:"info"`
}

//String return Info string representation
func (r Info) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//NewInfo returns new info response
func NewInfo(info string) Response {
	return Response{Type: "error", Data: Error{info}.String()}
}

//Error response structure
type Error struct {
	Info string `json:"info"`
}

//String return Error string representation
func (r Error) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//NewError returns new error response
func NewError(err string) Response {
	return Response{Type: "error", Data: Error{err}.String()}
}

//Login request stucture
type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//String return Info string representation
func (r Login) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//Register request stucture
type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//String return Register string representation
func (r Register) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//SendAll request structure
type SendAll struct {
	Message string `json:"message"`
}

//String return SendAll string representation
func (r SendAll) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}

//ReceiveAll response structure
type ReceiveAll struct {
	Login   string `json:"login"`
	Message string `json:"message"`
}

//String return ReceiveAll string representation
func (r ReceiveAll) String() string {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	return string(JSON)
}
