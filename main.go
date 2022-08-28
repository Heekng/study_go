package main

import "fmt"

func canIDrink(age int) bool {
	//if age < 18 {
	//	return false
	//} else {
	//	return true
	//}

	//if age < 18 {
	//	return false
	//}
	//return true

	// variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
