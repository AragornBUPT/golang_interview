// 题目：设计一个类，我们只能生成该类的一个实例

package singleton

import "sync"

type LazySingleton struct {}

var lazyInstance *LazySingleton
var once sync.Once

func GetLazyInstance() *LazySingleton {
	once.Do(func ()  {
		lazyInstance = &LazySingleton{}
	})

	return lazyInstance
}
