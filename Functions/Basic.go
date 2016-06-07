package main

import "fmt"
import "sort"

var sliceOfNums = []int{3, 4, 6, 7, 20, 21, 7} //create a slice of ints

func main() {

	xs := filter(sliceOfNums, func(n int) bool { //call the filter functions and give it a slice of ints, it will call a function that will append the conditions if they are true
		return n%2 == 0 //if the value 'n' which iterates through each spot in your slice is true, then it will be appended to the slice, otherwise it won't be
	})
	fmt.Println(xs)

	fmt.Println(reverse(sliceOfNums))
}

func doubleMe(val int) int { // basic functions, takes first parenthesis is the parameters
	// the return value always comes after the parameters

	return val * 2
}

func factorial(val int) int { //recursive function
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

//example of a callback function
/*
create a func called visit that takes a slice of
ints as a parameter.  it will return a func called 'call'
that takes a parameter of an int, and bool and returns a slice of ints
*/
func filter(num []int, call func(int) bool) []int {
	xs := []int{} //declare a spot in the memory to hold the array

	for _, n := range num {
		if call(n) { //if this condition is true
			xs = append(xs, n) //you will append 'n' to the slice 'xs'
		}
	}

	return xs //return the slice xs because your function needs to return a slice as declared in line 14
}

func reverse(val []int) []int {
	xs := make([]int, len(sliceOfNums))
	xs = append(sliceOfNums)

	sort.Sort(sort.Reverse(sort.IntSlice(xs)))

	return xs
}
