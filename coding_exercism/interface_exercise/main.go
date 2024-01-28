package main

import "fmt"
func main(){


	type Greeter interface{
		LanguageName() string
		Greet(string) string
	}

	type Italian struct{
		language string
	}

	// => "I can speak German: Hallo Dietrich!"
	//SayHello("Dietrich", germanGreeter)

	func SayHello(visitorName string, greeter Greeter){
		fmt.Println("pass")
		fmt.Println(visitorName)
		fmt.Println(greeter)
	}

	func (i Italian) LanguageName() string{
		return i.language
	}




}