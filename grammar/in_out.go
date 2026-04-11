package grammar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() {
	// 每次读取一个字符串（以空格作为分隔符）
	// var line string
	// fmt.Scan(&line)
	// fmt.Println(line)

	// var word1, word2 string
	// fmt.Scan(&word1, &word2)
	// fmt.Println(word1, word2)
	// fmt.Println(word2)

	a := ""
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		}

		fmt.Println(a)
	}
}

func InputList() {
	// var s string = ""
	// fmt.Scanln(&s)
	// sList := strings.Split(s, " ")

	// fmt.Println(sList)

	// var iList []int
	// for _, i := range sList {
	// 	i_t, _ := strconv.Atoi(i)
	// 	iList = append(iList, i_t)
	// }

	// fmt.Println(iList)

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
		fmt.Println(input)
	}

	inputList := strings.Fields(input)
	fmt.Println(inputList)

	var iList []int
	for _, i_s := range inputList {
		i_t, _ := strconv.Atoi(i_s)
		iList = append(iList, i_t)
	}
	fmt.Println(iList)
}
