package main

import (
	"fmt"
)
func Reverse(input string) string {	
    result := []rune(input)

	stringLength := len(result)

	for i:=0; i<stringLength/2; i++{
		result[i], result[stringLength-i-1] = result[stringLength-i-1],result[i]
	}

	return string(result)
}


func main() {
	input := "abc"
	reversed := Reverse(input)
	fmt.Println(reversed)
}

