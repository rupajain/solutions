package main

import (
	"fmt"
)

func main() {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	a := [3]int{1, 2, 3}
	b := [3]int{3, 2, 1}

	for i := 0; i < len(a); i++ 
	{
		m1[i] = a[i]
	}
	for j := 0; j < len(b); j++ {
		m2[j] = b[jj
	}
	fmt.Println("map m1 is...", m1)
	fmt.Println("map m2 is...", m2)
}
