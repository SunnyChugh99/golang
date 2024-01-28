package main

import (
	"fmt"
	"regexp"
)

func main(){
	fmt.Println("REGEX ASSIGNMENT!!!!")
	// re:= regexp.MustCompile(`(a|b)+`)
	// fmt.Println(re)

	re := regexp.MustCompile(`\d+`)
    text := "The price is $42 and $99."

    match := re.FindString(text)
    fmt.Println("Single match:", match)

    matches := re.FindAllString(text, -1)
    fmt.Println("All matches:", matches)

	re2 := regexp.MustCompile(`(\w+)-(\d+)`)
    text2 := "apple-123 banana-456"

    submatches := re2.FindStringSubmatch(text2)
    fmt.Println("Submatches:", submatches)

    submatchIndexes := re2.FindAllString(text2, -1)
    fmt.Println("Submatch indexes:", submatchIndexes)




	re3 := regexp.MustCompile(`Go`)
    text3 := "Golang is awesome!"

    index := re3.FindStringIndex(text3)
    fmt.Println("Index:", index)



    re4 := regexp.MustCompile(`[a-z]+d*`)
    b := re4.MatchString("[a12]")
    fmt.Println(b)

    c := re4.MatchString("abcc12bas")
    fmt.Println(c)

    d := re4.MatchString("2324 44")
    fmt.Println(d)



    sb := re4.FindString("[a12]")
    fmt.Println(sb)

    sc := re4.FindString("abcc12bas")
    fmt.Println(sc)

    sd := re4.FindString("2324 44")
    fmt.Println(sd)

}