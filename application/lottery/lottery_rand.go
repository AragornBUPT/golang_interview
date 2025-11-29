package lottery

import (
	"crypto/rand"
	"fmt"
	"golang_interview/algorithm/sort"
	"math/big"
	"strings"
)

func GenerateLottery() {
	num_list_prefix := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(35))
		num_list_prefix = append(num_list_prefix, int(num.Int64())+1)
	}
	sort.QuickSort(num_list_prefix, 0, len(num_list_prefix)-1)

	num_list_prefix_str := make([]string, 0, 5)
	for _, num := range num_list_prefix {
		num_str := fmt.Sprintf("%02d", num)
		num_list_prefix_str = append(num_list_prefix_str, num_str)
	}

	num_list_appendix := make([]int, 0, 2)
	for i := 0; i < 2; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(12))
		num_list_appendix = append(num_list_appendix, int(num.Int64())+1)
	}
	sort.QuickSort(num_list_appendix, 0, len(num_list_appendix)-1)

	num_list_appendix_str := make([]string, 0, 2)
	for _, num := range num_list_appendix {
		num_str := fmt.Sprintf("%02d", num)
		num_list_appendix_str = append(num_list_appendix_str, num_str)
	}

	fmt.Println(strings.Join(num_list_prefix_str, " "), "", strings.Join(num_list_appendix_str, " "))
}
