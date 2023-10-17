package znet

import(
	"fmt"
	"net"
	"time"
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
}


func CallBackToClient(conn *net.TCPConn,data []byte,cnt int) error{
	fmt.Println("[Conn Handle] CallBackToClient")
	if _,err := conn.Write(data[:cnt]);err != nil{
		fmt.Println("write back buf err ",err)
		return err
	}

	return nil
}

func (s *Server) Start() {

	fmt.Printf("[Start] Server Listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)
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
			dealConn := NewConnection(conn, cid, CallBackToClient)
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

func NewServer(name string) ziface.IServer{
	s:= &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 7777,
	}

	return s
}