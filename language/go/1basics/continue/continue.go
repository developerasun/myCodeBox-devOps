package main

import "fmt"

func main() {

	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

}
