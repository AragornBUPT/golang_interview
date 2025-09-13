package singleton

type HungerSingleton struct {}

var hungerInstance = &HungerSingleton{}

func GetHungerInstance() *HungerSingleton {
	return hungerInstance
}
