package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const helloInThai = "สวัสดี"

	fmt.Println(helloInThai)
	fmt.Println("Len:", len(helloInThai))

	for i := 0; i < len(helloInThai); i++ {
		fmt.Printf("%x ", helloInThai[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(helloInThai))

	for index, runeValue := range helloInThai {
		fmt.Printf("%#U starts at %d \n", runeValue, index)
	}

	fmt.Println("\nUsing DecodeRuneInString:")
	for currentIndex, lastWidth := 0, 0; currentIndex < len(helloInThai); currentIndex += lastWidth {
		runeValue, runeWidth := utf8.DecodeRuneInString(helloInThai[currentIndex:])
		fmt.Printf("%#U starts at %d \n", runeValue, currentIndex)
		lastWidth = runeWidth
		examineRune(runeValue)
	}
}

func examineRune(value rune) {
	if value == 't' {
		fmt.Println("Found tee")
		return
	}
	if value == 'ส' {
		fmt.Println("Found so sua")
	}
}
