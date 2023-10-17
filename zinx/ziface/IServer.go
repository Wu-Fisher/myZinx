package ziface

// Define a server interface

type IServer interface {
	// Start server
	Start()
	// Stop server
	Stop()
	// Run server
	Serve()
}

