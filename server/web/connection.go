package web

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

	"../utils"
)

//Connection type
type Connection struct {
	connection net.Conn
	login      string
	reader     *bufio.Reader
	writer     *bufio.Writer
	rMutex     sync.Mutex
	wMutex     sync.Mutex
}

//NewConnection handle net.Conn to parallel safe form
func NewConnection(c net.Conn) *Connection {
	return &Connection{connection: c, reader: bufio.NewReader(c), writer: bufio.NewWriter(c)}
}

//SetLogin sets connection's user login
func (c *Connection) SetLogin(login string) {
	c.login = login
}

//GetLogin returns connection's user login
func (c *Connection) GetLogin() string {
	return c.login
}

//SendResponse send response to this Connection
func (c *Connection) SendResponse(r Response) error {
	c.wMutex.Lock()
	defer c.wMutex.Unlock()
	log.Println("Send", r.String(), "to", c.RemoteAddr().String())
	data := []byte(r.String())
	len := len(data)
	n := 0
	for n < len {
		p, err := c.connection.Write(data)
		if err != nil && p == 0 {
			return err
		}
		n += p
	}
	return nil
}

//ReceiveRequest returns Request from this Connection
func (c *Connection) ReceiveRequest() (Request, error) {
	c.rMutex.Lock()
	defer c.rMutex.Unlock()
	JSON, err := utils.ReadJSON(c.reader)
	if err != nil {
		return Request{}, err
	}

	var request Request
	err = json.Unmarshal(JSON, &request)
	if err != nil {
		return Request{}, err
	}
	log.Println("Get", request.String(), "from", c.RemoteAddr().String())
	return request, nil
}

//Read implements the Conn Read method.
func (c *Connection) Read(b []byte) (int, error) {
	return c.connection.Read(b)
}

//Write implements the Conn Write method.
func (c *Connection) Write(b []byte) (int, error) {
	return c.connection.Write(b)
}

//RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.
func (c *Connection) RemoteAddr() net.Addr {
	return c.connection.RemoteAddr()
}

//LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.
func (c *Connection) LocalAddr() net.Addr {
	return c.connection.LocalAddr()
}

//Close closes the connection.
func (c *Connection) Close() error {
	return c.connection.Close()
}

//SetDeadline implements the Conn SetDeadline method.
func (c *Connection) SetDeadline(t time.Time) error {
	return c.connection.SetDeadline(t)
}

//SetReadDeadline implements the Conn SetReadDeadline method.
func (c *Connection) SetReadDeadline(t time.Time) error {
	return c.connection.SetReadDeadline(t)
}

//SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *Connection) SetWriteDeadline(t time.Time) error {
	return c.connection.SetWriteDeadline(t)
}
