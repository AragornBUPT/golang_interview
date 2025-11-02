package stringmy

import (
	"fmt"
	"reflect"
)

func GetType() {
	s := "afjewifo"
	for i, c := range s {
		fmt.Print("type of s[i] is ", reflect.TypeOf(s[i]))
		fmt.Print("; type of c is ", reflect.TypeOf(c))
		fmt.Printf("; type of c is %T\n", c)
	}
}
