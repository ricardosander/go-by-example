package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	fmt.Println(factorial(7))

	var fibonacci func(number int) int
	fibonacci = func(number int) int {
		if number < 2 {
			return number
		}
		return fibonacci(number-1) + fibonacci(number-2)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Fib %d: %d\n", i+1, fibonacci(i))
	}
}
