package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("test1")

	fmt.Println(time.Now())
	fmt.Println("test1")



	const shortForm = "2006-Jan-02"
	t, _ := time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t1, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t1)

	fmt.Println("check")
	t2 := time.Date(1995,time.September,22,13,0,0,0,time.UTC)
    formattedTime := t2.Format("Mon, 01/02/2006, 15:04") // string
    fmt.Println(formattedTime)


	fmt.Println("solve")
	date := "Tue, 09/22/1995, 13:00"
    layout := "Mon, 01/02/2006, 15:04"

    t, err := time.Parse(layout,date)
	
	t3 := time.Date(2019, time.July, 25, 13, 45, 0,0,time.UTC)
	formattedTime3 := t3.Format("7/25/2019 13:45:00") // string
	fmt.Println(formattedTime3)
//	13:45:00
}