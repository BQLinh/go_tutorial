package main

import "fmt"

type Person struct {
	name string
	age  int
}

func increase(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	*x += 1
}

func changeInfo(p *Person, name string) {
	p.age += 1
	p.name = name
}

func main() {
	p := Person{name: "linh", age: 12}

	changeInfo(&p, "tom")

	fmt.Println(p.name, p.age)

}
