package singleton

import (
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
)

func GetSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
