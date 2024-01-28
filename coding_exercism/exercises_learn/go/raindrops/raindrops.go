package raindrops

import "strconv"
func Convert(number int) string {
	finalResponse := ""
	if number%3 == 0{
		finalResponse = finalResponse + "Pling"
	}
	if number%5 == 0{
		finalResponse = finalResponse + "Plang"
	}
	if number%7 == 0{
		finalResponse = finalResponse + "Plong"
	}
	if finalResponse == ""{
		finalResponse = finalResponse+ strconv.Itoa(number)
	} 
	return finalResponse
}
