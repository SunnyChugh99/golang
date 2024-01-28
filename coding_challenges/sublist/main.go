package main

import (
	"fmt"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	eqc := 0
		
	if len(l1)<len(l2){
		for i:=0;i<len(l1);i++{
			for j:=0; j<len(l2); j++{

				fmt.Println("check between",i,j)
				
				if i < len(l1){
					if j < len(l2){
					if l1[i] == l2[j]{
					eqc++
					fmt.Println(i)
					fmt.Println(j)
					i++
				}else{
					fmt.Println("in else")
					eqc = 0
					j = j - i
					i = 0
					continue
				}
				
			}
		}
		}
		break

		}
	}else {
		fmt.Println("main else")
		for i:=0;i<len(l2);i++{
			for j:=0; j<len(l1); j++{

				fmt.Println("check between",i,j)
				
				if i < len(l2){
					if j < len(l1){
					if l2[i] == l1[j]{
					eqc++
					fmt.Println(i)
					fmt.Println(j)
					i++
				}else{
					fmt.Println("in else")
					eqc = 0
					j = j - i
					i = 0
					continue
				}
				
			}
		}
		}
		break

		}

	}
	fmt.Println(eqc)

	if eqc == len(l1){
		if len(l1) == len(l2){
		return RelationEqual
		}
		if len(l2) > len(l1){
			return RelationSublist
		}
	}

	if eqc == len(l2){
		if len(l1) > len(l2){
			return RelationSuperlist
		}
	}
	return RelationUnequal
}

func main(){

	l1 := []int{0, 1, 2, 3, 4, 5}
	l2 := []int{2, 3, 1}
	fmt.Println("Performing sub list")
	fmt.Println(Sublist(l1,l2))	
	
}