package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	arr := []int{1, 2, 3, 4}
	arr2 := []int{1, 2, 3, 4}
	c := append(arr, arr2...)
	fmt.Println(c)

}
