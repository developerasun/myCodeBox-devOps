package main

import (
	"fmt"
	"strings"
)

func panicTest() {
	// recover from panic
	defer func() {
		r := recover()
		fmt.Println(r)
	}()

	var a = [4]int{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	panicTest()
	fmt.Println("Hello") // still executable after panic
	result := strings.Contains("Hello World", "k") // still executable after panic
	fmt.Println(result) // still executable after panic
}
