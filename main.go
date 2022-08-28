package main

import (
	"fmt"
	"time"
)

func main() {
	go coolCount("heekng")
	go coolCount("test")
	time.Sleep(time.Second * 5)
}

func coolCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is cool", i)
		time.Sleep(time.Second)
	}
}
