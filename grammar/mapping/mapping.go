package main

import (
	"fmt"
	"reflect"
)

func main() {
	s2t := make(map[byte]byte)
	fmt.Println(s2t['s'])
	fmt.Println(reflect.TypeOf('s'))
}
