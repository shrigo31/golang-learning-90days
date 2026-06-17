package main

import "fmt"

func main() {

	rectangle := Rectangle{
		Length: 10,
		Width:  5,
	}

	circle := Circle{
		Radius: 5,
	}

	displayArea(rectangle)
	displayArea(circle)

	dog1 := Dog{}

	cat1 := Cat{}

	makeSound(dog1)
	makeSound(cat1)

	//-------------

	addtiondata1 := AdditionData{NumOne: 1, NumTwo: 2}
	subtractiondata1 := SubtractionData{NumOne: 4, NumTwo: 2}
	Addition(addtiondata1)
	Addition(subtractiondata1)

	rectangle1 := NewRectangle{
		Length: 10,
		Width:  5,
	}

	circle1 := NewCircle{
		Radius: 5,
	}

	square := Square{
		Side: 4,
	}

	displayArea(rectangle1)
	displayArea(circle1)
	displayArea(square)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

//-------------

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

//-------------------Use interface

func displayArea(s Shape) {
	fmt.Println(s.Area())
}

type Animal interface {
	Sound()
}

//-------------------------

type Dog struct {
}

func (d Dog) Sound() {
	fmt.Println("Bark")
}

//----------------------------

type Cat struct {
}

func (c Cat) Sound() {
	fmt.Println("Meow")
}

//---------------
func makeSound(a Animal) {
	a.Sound()
}

//problem 5
type Calculator interface {
	Calculate()
}

type AdditionData struct {
	NumOne int
	NumTwo int
}

func (a AdditionData) Calculate() {
	result := a.NumOne + a.NumTwo
	fmt.Println(result)
}

func Addition(c Calculator) {
	c.Calculate()
}

//---------------

type SubtractionData struct {
	NumOne int
	NumTwo int
}

func (s SubtractionData) Calculate() {
	result := s.NumOne - s.NumTwo
	fmt.Println(result)
}

//------------------ mini assignment

type NewShape interface {
	Area() float64
}

type NewRectangle struct {
	Length float64
	Width  float64
}

func (r NewRectangle) Area() float64 {
	return r.Length * r.Width
}

type NewCircle struct {
	Radius float64
}

func (c NewCircle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
