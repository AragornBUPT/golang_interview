package main

import (
	"fmt"
	// "golang_interview/grammar/math_app"
	// "golang_interview/grammar/concurrency"
	// "golang_interview/application"
	// "golang_interview/grammar/string_my"
	// "golang_interview/algorithm/leetcode"
	// "golang_interview/application"
	// "golang_interview/grammar"
	// "golang_interview/application/random"
	// "golang_interview/application/lottery"
)

func main() {
	fmt.Println("main")
	fmt.Println()

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
	// 对时间戳取模，生成随机彩票数字
	// lottery.GenerateLottery()

	// 测试strings.Builder用法
	// stringmy.TestBuilder()

	// leetcode，随机生成题号
	// random.RandLeetcode()

	// 牛客网，随机生成题号
	// random.RandNewCoder()

	// leetcode题目，将杂乱无章的数字排序
	// mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	// nums := []int{991, 338, 38}
	// leetcode.SortJumbled(mapping, nums)
	// leetcode.TestSortMap(mapping)

	// 读取标准输入输出
	// grammar.Input()
}
