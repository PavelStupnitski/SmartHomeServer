package app

import (
	"net/http"
	"sync"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Cert     string `yaml:"cert"`
}

type App struct {
	client      *http.Server
	config      Config
	channelData chan interface{}
	channelExit chan struct{}
	sync.RWMutex
}

type MockApp interface {
	SetClient(c http.Server)
	GetClient() http.Server
}

var instance *App = nil
var once sync.Once

func GetInstance() App {
	once.Do(func() {
		instance = new(App)
	})
	return App{}
}

func (m App) SetClient(n *http.Server) {
	m.Lock()
	defer m.Unlock()
	m.client = n
}

func (c *App) GetClient() *http.Server {
	c.RLock()
	defer c.RUnlock()
	return c.client
}
