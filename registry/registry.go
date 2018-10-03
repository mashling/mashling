package registry

import (
	"sync"
)

// Service encapsulates everything necessary to execute a step against a target.
type Service interface {
	Execute() (err error)
	UpdateRequest(values map[string]interface{}) (err error)
}

type Factory func(name string, settings map[string]interface{}) (Service, error)

var (
	mutex    sync.RWMutex
	registry = make(map[string]Factory, 32)
)

func Register(t string, factory Factory) {
	mutex.Lock()
	defer mutex.Unlock()
	registry[t] = factory
}

func Lookup(t string) Factory {
	mutex.RLock()
	defer mutex.RUnlock()
	return registry[t]
}
