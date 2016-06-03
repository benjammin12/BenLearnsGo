package main

import "fmt"

func main() {
	/*
	  Main difference between a slice and and array is that a slice is declared with out a
	  value n inside the brackets [n].  This is because slices are aloud to change sizes
	*/

	sly := make([]int, 4) //makes a slice
	fmt.Println(sly)
	sly = append(sly, 6)
	fmt.Println(sly)

	xs := []int{3, 4, 5, 6} //creates a slice with 4 data members
	fmt.Println(xs)
	copy(sly, xs) //copys the second parameter into the first parameter
	fmt.Println(sly)

	ds := make([]int, 4)     //creates a blank slice of 4
	sly = append(sly, ds...) //appends the entire ds to the end of sly
	fmt.Println(sly)

}
