package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of boolss

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	fmt.Println("test1111")
    fmt.Println(cb)
    character := cb[file]
    returnValue := 0
    for _,val := range character{
        if val == true{
            returnValue++
        }
    }
    return returnValue
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    returnValue := 0
    for k, _ := range cb{
        extBool := cb[k]
    	if extBool[rank] == true{
            returnValue++
        }
    }
	return returnValue
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
