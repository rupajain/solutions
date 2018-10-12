package main

import (
	"fmt"
	"sort"
)

func main() {
	var len1 int
	var len2 int
	fmt.Println("enter lentgh of slice1.......")
	fmt.Scan(&len1)
	fmt.Println("enter lentgh of slice2.......")
	fmt.Scan(&len2)
	s := make([]int, len1)
	fmt.Println("enter elements of slice 1....:")
	for i := 0; i < len(s); i++ {
		fmt.Scan(&s[i])
	}
	// s[0] = 1
	// s[1] = 2
	// s[2] = 3
	// s[3] = 4
	// s[4] = 5

	fmt.Println("slice s is.............:", s)

	//s1 := make([]int, 7)
	s1 := make([]int, len2)

	fmt.Println("enter elements of slice 2....:")
	for i := 0; i < len(s1); i++ {
		fmt.Scan(&s1[i])
	}
	// s1[0] = 5
	// s1[1] = 4
	// s1[2] = 1
	// s1[3] = 3
	// s1[4] = 2
	// s1[5] = 6
	// s1[6] = 7
	fmt.Println("slice s2 is.............:", s1)

	if len(s) != len(s1) {
		fmt.Println("slices are not equal...")
	} else {
		//s = sortingarray(s)
		s1 = sortingarray(s1)

		s = sortsliceusingmethodsort(s)
		fmt.Println("slices after sorting........", s)
		fmt.Println("slices after sorting ..........", s1)

		flagvar2 := compareelements(s, s1)
		if flagvar2 == false {
			fmt.Println("slices are not equal...")
		} else {
			fmt.Println("slices are equal..")
		}
	}

}

func sortsliceusingmethodsort(a []int) []int {
	sort.Sort(sort.IntSlice(a))
	return a
}
func sortingarray(a []int) []int {
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
func compareelements(a []int, b []int) bool {

	for i := 0; i < len(a); i++ {

		if a[i] != b[i] {
			return false
		}

	}
	return true
}
