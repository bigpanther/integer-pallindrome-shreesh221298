package main

import "fmt"

func IsPalindrome(input int) bool {
	// Do your implementation here
	i := 1
	newNumber := 0
	for input%i != input {
		newNumber = (newNumber * 10) + ((input / i) % 10)
		i = i * 10
	}
	return (newNumber == input)
}

func main() {
	fmt.Println(123) // false
	fmt.Println(121) // true
	fmt.Println(1)   // true
	fmt.Println(11)  // true
	fmt.Println(12)  // false
}
