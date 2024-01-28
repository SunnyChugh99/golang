package main

import "fmt"

func main() {
	SumMultiples(20, 3, 5) // Pass divisors as separate arguments
}

func SumMultiples(limit int, divisors ...int) int { // Use variadic parameter syntax
	fmt.Println(divisors)
	var multList []int
	for _, d := range divisors {
		fmt.Println(d)
		for i := 1; i <= limit; i++ {
			if d*i < limit {
				fmt.Println("inside")
				in := d * i

				multList = append(multList, in)
			}
		}
	}
	fmt.Println("Printing multiples list")
	fmt.Println(multList)

	set := make(map[int]bool)
	for _, ele := range multList {
		set[ele] = true
	}
	fmt.Println(set)
	finalVal := 0
	for key := range set {
		fmt.Println(key)
		finalVal += key
	}
	fmt.Println(finalVal)
	return finalVal
}
