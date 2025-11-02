package grammar

import (
	"container/list"
	"fmt"
)

// IsTransformable 判断字符串s是否可以通过相邻数字交换转换为字符串t
func IsTransformable(s string, t string) bool {
	pos := make(map[int]*list.List)
	for i := range 10 {
		pos[i] = list.New()
	}
	for i, d := range s {
		digit := int(d - '0') // 转换为实际数字值
		pos[digit].PushBack(i)
	}

	for _, d := range t {
		digit := int(d - '0') // 转换为实际数字值
		digit_queue := pos[digit]

		// s与t的某个字符数量不一致
		if digit_queue.Len() == 0 {
			return false
		}
		digit_ele := digit_queue.Front()
		digit_index := digit_ele.Value.(int)
		// 检查比当前数字小的所有数字
		for j := 0; j < digit; j++ {
			if pos[j].Len() > 0 {
				pos_index := pos[j].Front().Value.(int)
				if pos_index < digit_index {
					return false
				}
			}
		}

		digit_queue.Remove(digit_ele)
	}

	return true
}

// 牛客网题目
func TransferList() {
	s := make([]int, 0, 3)
	s = append(s, 1)
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	s = append(s, 2)
	s[0] = 2
	fmt.Println("in f", s)
}
