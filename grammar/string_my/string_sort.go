package stringmy

import (
	"fmt"
	"sort"
	"strings"
)

func SortString() {
	s := "joifjewofinoagoing"
	sList := strings.Split(s, "")
	sort.Strings(sList)
	fmt.Println(sList)

	s = strings.Join(sList, "")
	fmt.Println(s)
}
