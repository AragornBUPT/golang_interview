package grammar

import "fmt"

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
