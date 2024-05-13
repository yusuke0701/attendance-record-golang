package router

import (
	"errors"
	"log"
	"strconv"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid web server instance")
)

const (
	InstanceEcho int = iota
)

func NewWebServerFactory(instance int, port string) (Server, error) {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	switch instance {
	case InstanceEcho:
		return newEchoServer(Port(p)), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
