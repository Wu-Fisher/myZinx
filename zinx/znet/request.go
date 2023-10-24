package znet

import(
	"zinx/ziface"
)

type Request struct {
	// not tcp connection
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
