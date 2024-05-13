package router

import (
	"errors"
	"log"
	"strconv"

	ilog "attendance-record/infrastructure/log"
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

func NewWebServerFactory(instance int, logInstance int, port string) Server {
	l, err := ilog.NewLoggerFactory(logInstance)
	if err != nil {
		log.Fatalln(err)
	}

	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		l.Fatalln(err)
	}

	switch instance {
	case InstanceEcho:
		return newEchoServer(l, Port(p))
	default:
		l.Fatalln(errInvalidWebServerInstance)
	}
	return nil
}
