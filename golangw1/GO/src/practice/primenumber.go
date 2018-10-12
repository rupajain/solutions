//prime number is a whole number greater than 1 and divisible by 1 and itself.
package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("enter prime number to check...")
	var number int
	fmt.Scanln(&number)

	//b := checkPrime(number)
	//b := checkprimebysquareroot(number)
	//b := checkprimenumber1(number)
	//b := checkprimenumber2(number)
	//b := checkPrimeNew(number)
	b := checkPrimeNew1(number)
	if b == true {
		fmt.Println("number  is prime...", number, b)
	} else {
		fmt.Println("number is not prime...", number, b)
	}
}

func checkPrime(n int) bool {
	fmt.Println("number to check whether it is prime or not..", n)
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {

		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}

	}
	return true
}

func checkprimebysquareroot(n int) bool {

	b := true

	for i := 2; math.Sqrt(float64(n)) >= float64(i); i++ {
		if n%i == 0 {
			b = false

		}

	}
	return b
}

func checkprimenumber1(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}

func checkprimenumber2(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
func checkPrimeNew(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	b := true

	for i := 3; math.Sqrt(float64(n)) >= float64(i); i = i + 2 {
		if n%i == 0 {
			b = false

		}

	}
	return b
}

func checkPrimeNew1(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	b := true

	for i := 3; n >= i*i; i = i + 2 {
		if n%i == 0 {
			b = false

		}

	}
	return b
}
