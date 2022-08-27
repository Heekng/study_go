package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	//fmt.Println(multiply(2, 2))

	// go 는 한번도 사용하지 않는 변수, 상수에 대해 에러를 나타낸다.
	//totalLength, upperName := lenAndUpper("heekng")
	//fmt.Println(totalLength, upperName)

	//아래와 같이 리턴값의 개수를 조절할 수 있다.
	//totalLength, _ := lenAndUpper("heekng")
	//fmt.Println(totalLength)

	// 다중
	repeatMe("h", "e", "e", "k", "n", "g")
}
