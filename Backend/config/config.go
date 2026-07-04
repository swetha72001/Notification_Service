package config

import (
	"sync"
)

var mutex sync.RWMutex

type NSConfig struct {
	Port              string
	MongoURL          string
	MongoDatabaseName string
}

func (a *NSConfig) GetConfig() *NSConfig {
	mutex.RLock()

	defer mutex.RUnlock()
	return a
}

func (a *NSConfig) SetConfig(new NSConfig) {
	mutex.Lock()

	defer mutex.Unlock()
	*a = new

}
