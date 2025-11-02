package main

import (
	"fmt"
	// "golang_interview/grammar/math_app"
	// "golang_interview/grammar/concurrency"
	// "golang_interview/application"
	"golang_interview/grammar/string_my"
)

func main() {
	fmt.Println("main")

	// 测试大数
	// mod := 10 ^ 9 + 7
	// fmt.Println(mod)
	// mod1 := 1000000007
	// fmt.Println(mod1)

	/* // leetcode lru链表
	var lruCache leetcode.LRUCache
	lruCache = leetcode.Constructor(2)

	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)
	lruCache.Get(2)
	lruCache.Put(4, 4)
	lruCache.Get(1)
	lruCache.Get(3)
	lruCache.Get(4) */

	// 使用grammar包中的IsTransformable函数
	// result := grammar.IsTransformable("321", "123")
	// fmt.Printf("Is \"321\" transformable to \"123\": %v\n", result)

	// 测试byte和rune类型
	// grammar.GetType()

	// 牛客网题目，slice作为参数传递后是否会被修改
	// grammar.TransferList()

	// 牛客网题目，for中使用闭包，传参问题
	// concurrency.TestWaitGroup()

	// 牛客网题目，统计好三元组的数量
	// arr := []int{3, 0, 1, 1, 9, 7}
	// a := 7
	// b := 2
	// c := 3
	// math_app.CountGoodTriplets(arr, a, b, c)

	// 将彩票结果格式化
	// application.FormatNumbers()

	// 测试strings.Builder用法
	stringmy.TestBuilder()
}
