package Singleton

import (
	"sync"
	"sync/atomic"
)

//simulate the sync.Once()
type singleton struct{}

var (
	instance    *singleton
	initialized uint32 //mutex operation causes may cause serious performance problems ,so add an atomic counter to reduce the use of mutex
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance

}

var so sync.Once

//based on sync.Once() singleton
func InstanceSyncOnce() *singleton {
	so.Do(func() {
		instance = &singleton{}
	})
	return instance
}
