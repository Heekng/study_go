package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	// defer: funtion이 끝난 후 작동
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("heekng")
	fmt.Println(totalLength, up)
}
