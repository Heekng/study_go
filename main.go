package main

import "fmt"

func main() {
	//map[key]value
	//heekng := map[string]string{"name": "heekng", "age": "28"}
	//fmt.Println(heekng)

	//map[key]value
	heekng := map[string]string{"name": "heekng", "age": "28"}
	fmt.Println(heekng)
	for key, value := range heekng {
		fmt.Println(key, value)
	}
}
