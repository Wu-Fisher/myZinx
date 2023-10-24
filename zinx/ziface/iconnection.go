package ziface
import "net"

type IConnection interface {
	// Start connection
	Start()
	// Stop connection
	Stop()
	// Get connection ID
	GetConnID() uint32
	// Get TCP connection
	GetTCPConnection() *net.TCPConn
	// Get remote client IP address
	RemoteAddr() net.Addr
}


type HandFunc func(*net.TCPConn, []byte, int) error