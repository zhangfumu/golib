package server

import (
	conf "github.com/nicholaskh/jsconf"
	"time"
)

type Server struct {
	*conf.Conf

	Name       string
	configFile string
	StartedAt  time.Time
	pid        int
	hostname   string
}

func NewServer(name string) (this *Server) {
	this = new(Server)
	this.Name = name

	return
}
