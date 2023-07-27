package opensearch

import (
	opensearch2 "github.com/opensearch-project/opensearch-go/v2"
	"sync"
)

type Config struct {
	Hosts    []string `yaml:"hosts"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	CaCert   string   `yaml:"ca_cert"`
}

type OpenSearch struct {
	client      *opensearch2.Client
	config      Config
	channelData chan interface{}
	channelExit chan struct{}
	sync.RWMutex
}

type MockOpenSearch interface {
	SetClient(c opensearch2.Client)
	GetClient() opensearch2.Client
}

var instance *OpenSearch = nil
var once sync.Once

func GetInstance() OpenSearch {
	once.Do(func() {
		instance = new(OpenSearch)
	})
	return OpenSearch{}
}

func (c OpenSearch) SetClient(n *opensearch2.Client) {
	c.Lock()
	defer c.Unlock()
	c.client = n
}

func (c *OpenSearch) GetClient() *opensearch2.Client {
	c.RLock()
	defer c.RUnlock()
	return c.client
}
