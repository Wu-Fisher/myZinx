package znet

import (
	"fmt"
	"net"
	"reflect"
	"time"
	"zinx/utils"
	"zinx/ziface"
)

// IServer implementation
type Server struct {
	// Server name
	Name string
	// Server bind IP version
	IPVersion string
	// Server bind IP
	IP string
	// Server port
	Port int
	//
	Router ziface.IRouter
}


// func CallBackToClient(conn *net.TCPConn,data []byte,cnt int) error{
// 	fmt.Println("[Conn Handle] CallBackToClient")
// 	if _,err := conn.Write(data[:cnt]);err != nil{
// 		fmt.Println("write back buf err ",err)
// 		return err
// 	}

// 	return nil
// }

func (s *Server) Start() {

	fmt.Printf("[Start] Server Listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	fmt.Print("[Zinx] Version ", utils.GlobalObject.Version, " MaxConn ", utils.GlobalObject.MaxConn, " MaxPacketSize ", utils.GlobalObject.MaxPacketSize, "\n")
	go func(){
		addr,err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}

		listener,err:=	net.ListenTCP(s.IPVersion, addr)

		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}

		fmt.Printf("start Zinx server %s success, now listenning...\n", s.Name)
		
		var cid uint32 = 0
		//cid = 0

		for {
			conn, err  := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {

	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()

	for{
		time.Sleep(10*time.Second)
	}

}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	// fmt.Println("Add Router success!")
	fmt.Println("Add Router success! ", reflect.TypeOf(router).Name())
}

func NewServer(name string) ziface.IServer{

	utils.GlobalObject.Reload()

	s:= &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TcpPort,
		Router: nil,
	}

	return s
}