package main

import (
	"fmt"
)

func main() {
	var word1 = []string{"ab", "c"}
	var word2 = []string{"a", "bc"}
	var a string
	var b string

	for i := 0; i < len(word1); i++ {
		a = a + word1[i]
		fmt.Println(a)
	}
	for j := 0; j < len(word2); j++ {
		b = b + word2[j]
		fmt.Println(b)
	}

	if a == b {
		fmt.Println("True")

	} else {
		fmt.Println("False")
	}
}
