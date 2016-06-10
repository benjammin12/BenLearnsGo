package main

import "fmt"

func main() {
	x := 5

	z := &x //declared z that points to x in the memory adresss

	// & means takes adress

	// * means follow pointer to the value

	fmt.Println(z) //so this will only print x's memory adress

	fmt.Println(*z) //this will print the value of x

	*z = 12 //changes the value of x to 9 because z is pointing to x's memory adress
	//*z follows z to its memory adress and changes it
	fmt.Println(x)
}
