package main

import (
	"flag"
	"net"
	"os"
)

func init() {
	flag.StringVar(&Address, "port", DefaultAddress, "Default address i.e 127.0.0.1:7777")
	flag.StringVar(&Message, "message", DefaultMessage, "Default message to send i.e Caddy is awesome!")

}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", Address)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(Message))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", Message)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	conn.Close()
}

const (
	// DefaultAddress is the default port.
	DefaultAddress = "127.0.0.1:7777"

	// DefaultMessage is the default message
	DefaultMessage = "Caddy is awesome!"
)

// These "soft defaults" are configurable by
// command line flags, etc.
var (
	// Address is the destination address
	Address = DefaultAddress

	// Message is the message to echo
	Message = DefaultMessage
)
