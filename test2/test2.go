package main

import (
	"math"
	"fmt"
)

type Shaper interface{
	Area() float32
}

type Square struct {
	a float32
}

type Cycle struct{
	r float32
}

type Rectangle struct{
	a, b float32

}

func main(){

	s := Square{5}
	c := Cycle{2.5}
	r := Rectangle{3,4}

	shapes := []Shaper{Shaper(s), c, r}

	for n, _ :=range shapes{
		fmt.Print("the detail of shape is ", shapes[n])
		fmt.Println("the area of this shape is ", shapes[n].Area())
	}


}

func (this Square) Area() float32{
	return this.a * this.a
}

func (this Cycle) Area() float32{
	return math.Pi * this.r * this.r
}

func (this Rectangle) Area() float32{
	return this.a * this.b
}