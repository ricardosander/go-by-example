package main

import "fmt"

func zeroToValue(numberValue int) {
	numberValue = 0
}

func zeroToPointer(numberPointer *int) {
	*numberPointer = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroToValue(i)
	fmt.Println("zeroToValue:", i)

	zeroToPointer(&i)
	fmt.Println("zeroToPointer:", i)

	fmt.Println("pointer:", &i)
}
