package main

import (
	"fmt"
	"sort"
)

func main() {

	a := [5]int{1, 2, 3, 4, 2}
	b := [5]int{3, 2, 2, 1, 4}

	//b := [5]int{1, 13, 3, 4, 5}
	//b := [7]int{1, 2, 3, 4, 5, 6, 7}
	if len(a) != len(b) {
		fmt.Println("arrays of different sizes...")
	} else {
		//a = sortingarray(a)
		a = sortingarrayusingsort(a)
		b = sortingarray(b)
		fmt.Println("array after sorting.array a is...", a)
		fmt.Println("array after sorting array b is..", b)

		flagvar2 := compareelements(a, b)
		if flagvar2 == false {
			fmt.Println("arrays are not equal...")
		} else {
			fmt.Println("arrays are equal..")
		}
	}

}

func sortingarray(a [5]int) [5]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			var temp int
			if a[i] > a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}
func sortingarrayusingsort(a [5]int) [5]int {
	return sort.Ints(a)
}
func compareelements(a [5]int, b [5]int) bool {

	for i := 0; i < len(a); i++ {

		if a[i] != b[i] {
			return false
		}

	}
	return true
}
