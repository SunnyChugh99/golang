package main

import "fmt"


func main(){
	fmt.Println("pointers assignment")
	var a int
	a = 5
	fmt.Println(a)

	//accessing address of the variable
	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)


	// accessing value stored in pointer
	var b int
	b = *p
	fmt.Println(*(&a))
	fmt.Println(b)

	// pointers to struct

	type Person struct{
		Name string
		Age int
	}

	p2 := Person{Name: "Sunny", Age: 24}
	fmt.Println(p2)

	fmt.Println(p2.Name)
	var ptr *Person

	ptr = &p2
	fmt.Println(ptr.Name)


}
