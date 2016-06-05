package main

import "fmt"

func main() {

	fmt.Println(addUpAll(4, 3, 2, 3, 21, 22)) //takes mult values seperate with commas
	fmt.Println(doubleMe(13))
	fmt.Println(factorial(5))
}

func doubleMe(val int) int { // basic functions, takes first parenthesis is the parameters
	// the return value always comes after the parameters

	return val * 2
}

func factorial(val int) int {
	if val == 0 {
		return 1
	}

	return val * factorial(val-1)

}

func addUpAll(args ...int) int { // use (args ...int) if you don't know how many values you will take in
	sum := 0

	for _, val := range args { //use for loop to add numbers to sum
		sum += val
	}

	return sum
}
