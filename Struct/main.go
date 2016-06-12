package main

import "fmt"

func main() {

	rectObj := Rectangle{5, 6} //creating a struct or an object
	rectObj.x = 10             //sets first var x to 10
	area1 := rectObj.area()

	fmt.Println(area1)

	anotherRect := Rectangle{3, 2}
	fmt.Println(anotherRect.area())

}

//Rectangle x,y
type Rectangle struct { //create a struct with 2 variables representings sides of rect
	x, y int
}

func (rect Rectangle) area() int { //METHOD that calculates area
	//Methods have recievers, in this case the reciever is rect Rectangle
	//if you declare like func (rect *Rectangle) using a pointer instead
	// then any modification of rect.x or rect.y inside the function will change the Rectangle
	//object we created as well
	return rect.x * rect.y
}
