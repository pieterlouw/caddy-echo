package echoserver

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyfile"
)

const serverType = "tcpecho"

var directives = []string{
//"logs",
//"errors",
}

func init() {
	fmt.Println("echoserver init")

	flag.StringVar(&Port, serverType+".port", DefaultPort, "Default port")

	caddy.RegisterServerType(serverType, caddy.ServerType{
		Directives: func() []string { return directives },
		DefaultInput: func() caddy.Input {
			return caddy.CaddyfileInput{
				Filepath:       "Echofile",
				Contents:       []byte(fmt.Sprintf(".:%s\n", Port)),
				ServerTypeName: serverType,
			}
		},
		NewContext: newContext,
	})
}

func newContext() caddy.Context {
	return &tcpechoContext{}
}

type tcpechoContext struct {
	// keysToConfigs maps an address at the top of a
	// server block (a "key") to its Config. Not all
	// Configs will be represented here, only ones
	// that appeared in the Caddyfile.
	keysToConfigs map[string]*Config

	// configs is the master list of all site configs.
	configs []*Config
}

func (t *tcpechoContext) saveConfig(key string, cfg *Config) {
	t.configs = append(t.configs, cfg)
	t.keysToConfigs[key] = cfg
}

// InspectServerBlocks for tcpecho
func (t *tcpechoContext) InspectServerBlocks(sourceFile string, serverBlocks []caddyfile.ServerBlock) ([]caddyfile.ServerBlock, error) {

	log.Printf("[INFO] InspectServerBlocks %s", sourceFile)

	// For each address in each server block, make a new config
	for _, sb := range serverBlocks {

		for _, key := range sb.Keys {
			log.Printf("[INFO] range serverBlocks key %s", key)
			key = strings.ToLower(key)
			if _, dup := t.keysToConfigs[key]; dup {
				return serverBlocks, fmt.Errorf("duplicate: %s", key)
			}

			// Save the config to our master list, and key it for lookups
			cfg := &Config{}
			t.saveConfig(key, cfg)
		}
	}

	return serverBlocks, nil
}

// MakeServers uses the newly-created configs to create and return a list of server instances.
func (t *tcpechoContext) MakeServers() ([]caddy.Server, error) {
	// create a server
	var servers []caddy.Server

	s, err := NewServer(Port)
	if err != nil {
		return nil, err
	}
	servers = append(servers, s)

	return servers, nil
}

const (
	// DefaultPort is the default port.
	DefaultPort = "7777"
)

// These "soft defaults" are configurable by
// command line flags, etc.
var (
	// Port is the site port
	Port = DefaultPort
)
