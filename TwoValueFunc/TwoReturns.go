package main

import "fmt"

func main() {
	fmt.Println(returnTwo(4))
}

//functions in go can return two Values
func returnTwo(val int) (int, int) { //simply put the return values in parenthesis and seperate by a comma
	return val + 1, val + 2 //seperate the return by a comma, only write return once
}
