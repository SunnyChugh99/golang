package main

import (
	"fmt"
	"regexp"
)

func main(){

	fmt.Println("reg ex exercise")
	
	re := regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	
	b := re.MatchString("[ERR] A good error here")
	fmt.Println(b)

// 	SplitLogLine("section 1<*>section 2<~~~>section 3")
// // => []string{"section 1", "section 2", "section 3"},
	text := "section 1<*>section 2<~~~>section 3"
	re2 := regexp.MustCompile(`(?:<\*\>)|(?:<~~~>)`)

	// re2 := regexp.MustCompile(`<[^~~~]*>`)
	out := re2.Split(text, -1)
	fmt.Println(out)

	re3 := regexp.MustCompile(`"[^"]*password[^"]*"`)
	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,  // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`, // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one 
	}

	for _,line := range lines{
		fmt.Println(line)
		b := re3.MatchString(line)    
		fmt.Println(b)
	}


	re4 := regexp.MustCompile(`end-of-line\d+`)
	// s = re4.ReplaceAllString("12abc34(ef)", "X") // => "12X(X)"
	text2 := "[INF] end-of-line23033 Network Failure end-of-line27"
	result := re4.ReplaceAllString(text2, "")

	// separator := " "


	// result := strings.Join(s, separator)
	fmt.Println(result)
}