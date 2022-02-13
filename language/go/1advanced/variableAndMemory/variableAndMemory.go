package main

import "fmt"

func main() {
	// Ways to declare a slice
	// Initialize with values
	s1 := []int{1, 2, 3}

	// Initialize but with no vallues. Not preferred since memory is allocated but not used
	s2 := []int{}

	// Declare with var command: no memory allocation
	var s3 []int

	// Initialize with no value. Memory allocated, length and capacity of
	// slice is defined. If capacity is not set, its default is the same with the length.
	s4 := make([]int, 5, 7)

	// Compare each result
	fmt.Println(s1, s2, s3, s4)

	s4 = s4[2:4]
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	myslice1 := []int{1, 2, 3, 4, 5, 6}
	myslice2 := myslice1[2:4]
	fmt.Println("length of myslice2: ", len(myslice2))
	fmt.Println("capacity of myslice2: ", cap(myslice2))

	// Change the first element value of myslice2
	myslice2[0] = 777
	fmt.Println("third element of myslicd1: ", myslice1[2])

}
