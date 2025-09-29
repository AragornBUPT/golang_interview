// 题目：设计一个类，我们只能生成该类的一个实例

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
