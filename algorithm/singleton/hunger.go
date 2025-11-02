// 题目：设计一个类，我们只能生成该类的一个实例

package singleton

type HungerSingleton struct {}

var hungerInstance = &HungerSingleton{}

func GetHungerInstance() *HungerSingleton {
	return hungerInstance
}
