package main

import (
	"fmt"
	"unicode/utf8"
)


func main(){
	fmt.Println("Runes learn!!")
	myRune := 'r'
	fmt.Println(myRune)

	// type of rune object
	fmt.Printf("%T is type of rune object", myRune)
	
	// typeStore := fmt.Sprintf("%T is type of rune object", myRune)
	// fmt.Println("\n",typeStore)

	// read character stored in rune object
	fmt.Printf("\n%c is unicode character of r", myRune)

	// unicode code point of rune object
	fmt.Printf("\n%U is unicode code point of r", myRune)
	

	myString := "!hello"
	for index,char := range myString{
		fmt.Printf("\nindex %d\t Character %c \t Unicode code point %U", index, char, char)

	}

	stringLength := len(myString)

	numberOfRunes := utf8.RuneCountInString(myString)

	fmt.Printf("\nmyString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)
}