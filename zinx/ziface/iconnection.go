package ziface
import "net"

type IConnection interface {
	// Start connection
	Start()
	// Stop connection
	Stop()
	// Get connection ID
	GetConnID() uint32
}


type HandFunc func(*net.TCPConn, []byte, int) error