package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func SortJumbled(mapping []int, nums []int) []int {
	sortMapList := make([]SortMap, 0, len(nums))
	for _, num := range nums {
		s := SortMap{
			mapping: mapping,
			num:     num,
		}
		sortMapList = append(sortMapList, s)
	}
	fmt.Println(sortMapList)

	sort.Sort(ByMap(sortMapList))
	fmt.Println(sortMapList)

	res := make([]int, 0, len(nums))
	for _, sortMap := range sortMapList {
		res = append(res, sortMap.num)
	}
	return res
}

func TestSortMap(mapping []int) {
	sortMap := SortMap{
		mapping: mapping,
		num:     991,
	}

	sortMap.Mapping()
}

type SortMap struct {
	mapping []int
	num     int
}

func (s SortMap) Mapping() int {
	num := s.num
	mapping := s.mapping

	res := 0
	for i := 0; num != 0; i++ {
		key := num % 10
		val := mapping[key]
		res = res + int(math.Pow10(i))*val

		num = (num - key) / 10
	}

	// fmt.Println(res)
	return res
}

type ByMap []SortMap

func (b ByMap) Len() int {
	return len(b)
}

func (b ByMap) Less(i, j int) bool {
	return b[i].Mapping() < b[j].Mapping()
}

func (b ByMap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
