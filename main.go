package main

import "fmt"

func main() {
	//a := 2
	//b := a
	//a = 10
	//fmt.Println(a, b)
	// 주소값을 확인
	//fmt.Println(&a, &b)

	a := 2
	b := &a
	a = 5
	// *: 살펴본다, 훑어본다., 메모리에 할당된 값을 확인한다.
	*b = 20
	fmt.Println(*b)
	fmt.Println(a)
}
