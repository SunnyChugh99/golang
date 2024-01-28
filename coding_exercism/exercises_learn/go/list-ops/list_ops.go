package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
    result := initial
    for _, value := range s {
        result = fn(result, value)
    }
    return result}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
    result := initial
    for i := len(s) - 1; i >= 0; i-- {
        result = fn(s[i], result)
    }
    return result

}

func (s IntList) Filter(fn func(int) bool) IntList {
	var finalList []int
	for _, value := range (s){
		if fn(value) {
			finalList = append(finalList, value)
		}
	}
	return finalList
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
    var newList IntList // Use the IntList type for the return value
    for _, value := range s {
        newValue := fn(value)
        newList = append(newList, newValue)
    }
    return newList
}

func (s IntList) Reverse() IntList {
	length := len(s)

    // Iterate until reaching the middle of the array
    for i := 0; i < length/2; i++ {
        // Swap elements at the beginning and end of the array
        s[i], s[length-i-1] = s[length-i-1], s[i]
    }
	return s
}

func (s IntList) Append(lst IntList) IntList {
	for _,second_value :=  range (lst) {
		s = append(s, second_value)
	}
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _,list_second :=  range (lists) {
		for _, second_value := range(list_second){
			s = append(s, second_value)
		}

	}
	return s
}
