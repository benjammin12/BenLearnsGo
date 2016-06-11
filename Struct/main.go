package main

import "fmt"

func main() {

	rect := Rectangle{5, 6} //creating a struct or an object
	rect.x = 10             //sets first var x to 10
	area1 := rect.area()

	fmt.Println(area1)

	anotherRect := Rectangle{3, 2}
	fmt.Println(anotherRect.area())

}

//Rectangle x,y
type Rectangle struct {
	x, y int
}

func (rect *Rectangle) area() int {
	return rect.x * rect.y
}
