package main

import "fmt"

func main() {
	a := 2
	fmt.Println(a)
	fmt.Println(&a)
	a += 1
	fmt.Println(&a)
	a = 4
	fmt.Println((&a))
	b := a
	fmt.Println(&b)
	a += 1
	fmt.Println(&b)
	fmt.Println(a)
	fmt.Println(b)
}
