package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "jfwejoiwej"
	for i := range s {
		fmt.Println(i, s[i])
		fmt.Println(reflect.TypeOf(s[i]))
	}
}
