package main

import "fmt"

func main() {
	fmt.Println(returnTwo(12))
	fmt.Println(isEvenOrOdd(14))

}

//functions in go can return two Values
func returnTwo(val int) (int, int) { //simply put the return values in parenthesis and seperate by a comma
	return val + 1, val + 2 //seperate the return by a comma, only write return once
}

func isEvenOrOdd(val int) (int, bool) {
	return val / 2, val%2 == 0
}
