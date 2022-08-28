package main

import "fmt"

func main() {
	//array
	//names := [5]string{"h", "e", "e"}
	//names[3] = "k"
	//names[4] = "n"
	//fmt.Println(names)

	//slice
	names := []string{"h", "e", "e"}
	names = append(names, "k")
	names = append(names, "n")
	names = append(names, "g")
	fmt.Println(names)
}
