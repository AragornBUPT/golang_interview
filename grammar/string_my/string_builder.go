package stringmy

import (
	"fmt"
	"strings"
)

func TestBuilder() {
	s := "aaa"

	builder := strings.Builder{}
	for _, ch := range s {
		builder.WriteRune(ch)
		builder.WriteRune(ch)
	}

	fmt.Println(builder.String())
}
