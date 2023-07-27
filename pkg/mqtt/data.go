package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

type Config struct {
	Broker   string `yaml:"broker"`
	ClientID string `yaml:"client_id"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Handler  string `yaml:"handler"`
}

type MQTT struct {
	client      *mqtt.Client
	config      Config
	channelData chan interface{}
	channelExit chan struct{}
	sync.RWMutex
}

type MockMQTT interface {
	SetClient(c mqtt.Client)
	GetClient() mqtt.Client
}

var instance *MQTT = nil
var once sync.Once

func GetInstance() MQTT {
	once.Do(func() {
		instance = new(MQTT)
	})
	return MQTT{}
}

func (m MQTT) SetClient(n *mqtt.Client) {
	m.Lock()
	defer m.Unlock()
	m.client = n
}

func (c *MQTT) GetClient() *mqtt.Client {
	c.RLock()
	defer c.RUnlock()
	return c.client
}
