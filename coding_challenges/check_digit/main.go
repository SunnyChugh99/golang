package main

import (
	"fmt"
	"strconv"
)

func checkResult(num1, num2, result int) string {
	strNum1 := strconv.Itoa(num1)
	strNum2 := strconv.Itoa(num2)
	strResult := strconv.Itoa(result)

	// Ensure that the result has no more digits than the maximum of num1 and num2
	if len(strResult) > len(strNum1) && len(strResult) > len(strNum2) {
		return "Length mismatch"
	}

	for i := 0; i < len(strResult); i++ {
		digitNum1 := 0
		digitNum2 := 0

		if i < len(strNum1) {
			digitNum1, _ = strconv.Atoi(string(strNum1[len(strNum1)-1-i]))
		}

		if i < len(strNum2) {
			digitNum2, _ = strconv.Atoi(string(strNum2[len(strNum2)-1-i]))
		}

		digitResult, _ := strconv.Atoi(string(strResult[len(strResult)-1-i]))

		if (digitNum1+digitNum2)%10 != digitResult {
			return strconv.Itoa(i + 1) // Index (position) of the first incorrect digit
		}
	}

	return "ok"
}

func main() {
	// Example usage
	num1 := 1234
	num2 := 567
	calculatedResult := 1796

	result := checkResult(num1, num2, calculatedResult)

	if result == "ok" {
		fmt.Println("The calculated result is correct.")
	} else {
		fmt.Println("The calculated result is incorrect. Error at index:", result)
	}
}
