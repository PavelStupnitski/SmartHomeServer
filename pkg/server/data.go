package server

import (
	"net/http"
	"sync"
)

type Config struct {
	Hosts    []string `yaml:"hosts"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	CaCert   string   `yaml:"ca_cert"`
}

type Server struct {
	client      *http.Client
	config      Config
	channelData chan interface{}
	channelExit chan struct{}
	sync.RWMutex
}

type MockServer interface {
	SetClient(c http.Client)
	GetClient() http.Client
}

var instance *Server = nil
var once sync.Once

func GetInstance() Server {
	once.Do(func() {
		instance = new(Server)
	})
	return Server{}
}

func (c Server) SetClient(n *http.Server) {
	c.Lock()
	defer c.Unlock()
	c.client = n
}

func (c *Server) GetClient() *http.Server {
	c.RLock()
	defer c.RUnlock()
	return c.client
}
