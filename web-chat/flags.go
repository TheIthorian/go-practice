package main

import (
	"flag"
	"fmt"
	"strings"
)

type Options struct {
	port     *int
	ip       *string
	protocol *string
	isHost   *bool
	username *string
}

var allowedProtocols = []string{"tcp"}

func getOptions() (*Options, error) {
	port := flag.Int("port", 80, "Port for the websocket")
	ip := flag.String("ip", "127.0.0.1", "Ip address")
	protocol := flag.String("protocol", "tcp", "One of [tcp]")
	isHost := flag.Bool("isHost", false, "Are you the host?")
	username := flag.String("username", "(unknown)", "You display name")

	flag.Parse()

	if !isValidProtocol(protocol) {
		return nil, fmt.Errorf("%v is not a valid protocol. Must be one of [%v]", *protocol, strings.Join(allowedProtocols, ", "))
	}

	return &Options{
		port,
		ip,
		protocol,
		isHost,
		username,
	}, nil
}

func isValidProtocol(protocol *string) bool {
	for i := 0; i < len(allowedProtocols); i++ {
		if allowedProtocols[i] == *protocol {
			return true
		}
	}

	return false
}
