package echoserver

import (
	"flag"
	"fmt"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyfile"
)

const serverType = "echo"

//tcpecho don't have directives
var directives = []string{}

func init() {
	flag.StringVar(&TCPPort, serverType+".tcpport", DefaultTCPPort, "Default TCP port")

	caddy.RegisterServerType(serverType, caddy.ServerType{
		Directives: func() []string { return directives },
		DefaultInput: func() caddy.Input {
			return caddy.CaddyfileInput{
				Contents:       []byte(fmt.Sprintf(":%s\n", TCPPort)),
				ServerTypeName: serverType,
			}
		},
		NewContext: newContext,
	})
}

func newContext() caddy.Context {
	return &tcpechoContext{}
}

type tcpechoContext struct{}

// InspectServerBlocks for tcpecho is a no-op
func (t *tcpechoContext) InspectServerBlocks(sourceFile string, serverBlocks []caddyfile.ServerBlock) ([]caddyfile.ServerBlock, error) {
	return serverBlocks, nil
}

// MakeServers uses the newly-created configs to create and return a list of server instances.
func (t *tcpechoContext) MakeServers() ([]caddy.Server, error) {
	// create a server
	var servers []caddy.Server

	s, err := NewServer(TCPPort)
	if err != nil {
		return nil, err
	}
	servers = append(servers, s)

	return servers, nil
}

const (
	// DefaultTCPPort is the default TCP port.
	DefaultTCPPort = "7777"
)

// These "soft defaults" are configurable by
// command line flags, etc.
var (
	// TCPPort is the TCP port to listen on
	TCPPort = DefaultTCPPort
)
