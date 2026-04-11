package concurrency

import (
	"fmt"
	"time"
)

// 程序一
// 高版本的Go：输出无序；且如果主协程结束太快，会没有输出
// 低版本的Go：可能最后输出3个c
func PrintClosure() {
	values := []string{"a", "b", "c"}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(3 * time.Second) // 明确等待3秒，而非3纳秒
}

// 程序二
// 正确按序输出
func PrintClosure2() {
	values2 := []string{"a", "b", "c"}
	for _, val := range values2 {
		func() {
			fmt.Println(val)
		}()
	}
}

// 程序三
// 输出无序；且如果主协程结束太快，会没有输出
func PrintClosure3() {
	values4 := []string{"a", "b", "c"}
	for _, val := range values4 {
		v := val
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second) // 明确等待3秒，而非3纳秒
}

// 程序四
// 输出无序；且如果主协程结束太快，会没有输出
func PrintClosure4() {
	values5 := []string{"a", "b", "c"}
	for _, val := range values5 {
		go func(v string) {
			fmt.Println(v)
		}(val)
	}

	time.Sleep(3 * time.Second) // 明确等待3秒，而非3纳秒
}
