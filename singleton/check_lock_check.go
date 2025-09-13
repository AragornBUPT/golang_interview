package singleton

import (
	"sync"
)

type LockSingleton struct {
}

var lockInstance *LockSingleton
var mutex sync.Mutex

func GetLockInstance() *LockSingleton {
	if lockInstance != nil {
		return lockInstance
	}

	mutex.Lock()
	defer mutex.Unlock()

	if lockInstance == nil {
		lockInstance = &LockSingleton{}
	}

	return lockInstance
}
