package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)


type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool

	ExitBuffChan chan bool
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		Router: router,
		isClosed: false,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader(){
	fmt.Println("Reader Goroutine is running...")

	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()
	
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			// 小修
			if err == io.EOF {
				fmt.Println("remote closed!")
				return
			}
			fmt.Println("recv buf err ", err)
			continue
		}

		// if err:=c.handleAPI(c.Conn, buf, cnt);err!=nil{
		// 	fmt.Println("ConnID ", c.ConnID, " handle is error")
		// 	c.ExitBuffChan <- true
		// 	return
		// }

		req := Request{
			conn: c,
			data: buf,
		}
		go func(request ziface.IRequest){
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}


func (c *Connection)Start(){
	go c.StartReader()
	for {
		// wait for the connection to close
		<-c.ExitBuffChan
		// close the connection
		// c.Stop()		// exit the loop
		break
	}
}


func (c *Connection) Stop(){
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.ExitBuffChan <- true
	c.Conn.Close()

	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn{
	return c.Conn
}

func (c *Connection) GetConnID() uint32{
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}



