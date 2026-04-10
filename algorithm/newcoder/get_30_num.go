// package newcoder
package main

import "fmt"

/*
* 孝庸提前批笔试题
* 一列数的规则如下：1、1、2、3、5、8、13、21、34..求第30位数是多少，用递归算法实现任何编程语言都可以，请写出具体的代码
 */

func Get30Num(x, y int, index int) int {
	if index == 30 {
		return x + y
	}

	index++
	return Get30Num(y, x+y, index)
}

func main() {
	res := Get30Num(1, 1, 3)
	fmt.Println(res)
}
