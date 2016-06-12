package main

import "fmt"

func main() {

	rect1 := Rectangle{4, 5} //create a Rectangle object called rect1
	// or var Rectangle rect1
	// or rec1 := new(Rectangle)
	square1 := Square{8}
	fmt.Println("Rectange area is", getArea(rect1))
	fmt.Println("Square area is", getArea(square1))

	fmt.Println("Rectangle perimeter is", getPerimeter(rect1))
	fmt.Println("Square perimeter is", getPerimeter(square1))

}

//AreaOfShape returns area
type AreaOfShape interface { //Area interface that gets Area
	area() int
}

//PerimeterOfShape returns perimeter
type PerimeterOfShape interface {
	perimeter() int
}

//Rectangle width and height
type Rectangle struct { //Rectange structure, similiar to creating objects
	width  int //width
	height int //height
}

//Square one side
type Square struct { //square struct
	side int
}

func (s Square) area() int { //area of square
	return s.side * s.side
}

func (rect Rectangle) area() int { //create an area method that points to Rectangle struct
	return rect.width * rect.height
}

func (s Square) perimeter() int {

	return s.side * 4
}
func (rect Rectangle) perimeter() int {
	return 2 * (rect.width * rect.height)
}

func getArea(shape AreaOfShape) int { //use to get area of whatever shape you want
	return shape.area()
}

func getPerimeter(sh PerimeterOfShape) int {
	return sh.perimeter()
}
