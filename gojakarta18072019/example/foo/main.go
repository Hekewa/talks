package main

import (
	"fmt"
)

func main() {
	case1 := 1 + 0xF
	case2 := 1 - 1
	i := 1
	case3 := i + 2
	fmt.Printf("case 1: %s\n", case1)
	fmt.Printf("case 2: %s\n", case2)
	fmt.Printf("case 3: %s\n", case3)
}
