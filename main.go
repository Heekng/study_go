package main

import "fmt"

// struct: like object
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	//heekng := person{"heekng", 28, favFood}
	heekng := person{
		name:    "heekng",
		age:     28,
		favFood: favFood,
	}
	fmt.Println(heekng)
	fmt.Println(heekng.name)
}
