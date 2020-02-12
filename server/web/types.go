package web

import (
	"encoding/json"
	"fmt"
	"log"
)

/*==========================================================================
	Basic structures
===========================================================================*/

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

/*==========================================================================
	Response structures
===========================================================================*/

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
	return Response{Type: "info", Data: Error{info}.String()}
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

/*==========================================================================
	Request structures
===========================================================================*/

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

//UnwrapLogin convert JSON string to Login structure
func UnwrapLogin(JSON string) (Login, error) {
	var login Login
	err := json.Unmarshal([]byte(JSON), &login)
	if err != nil {
		return Login{}, err
	}
	return login, nil
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

//UnwrapRegister convert JSON string to Register structure
func UnwrapRegister(JSON string) (Register, error) {
	var register Register
	err := json.Unmarshal([]byte(JSON), &register)
	if err != nil {
		return Register{}, err
	}
	return register, nil
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

//UnwrapSendAll convert JSON string to SendAll structure
func UnwrapSendAll(JSON string) (SendAll, error) {
	var sendAll SendAll
	err := json.Unmarshal([]byte(JSON), &sendAll)
	fmt.Println(sendAll.Message)
	if err != nil {
		return SendAll{}, err
	}
	return sendAll, nil
}
