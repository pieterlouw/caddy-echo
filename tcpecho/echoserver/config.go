package echoserver

// Config configuration for a single echo server.
type Config struct {
	// The port to listen on.
	Port string

	// Server is the server that handles this config
	Server *Server
}
