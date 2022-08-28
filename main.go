package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [4]string{"hee", "kng", "test1", "test2"}
	for _, person := range people {
		go isCool(person, c)
	}
	fmt.Println("Waiting for messages")
	//resultOne := <-c
	//resultTwo := <-c
	//resultThree := <-c
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for", i)
		fmt.Println(<-c)
	}
	//fmt.Println("Received this message: " + resultOne)
	//fmt.Println("Received this message: " + resultTwo)
	//fmt.Println("Received this message: " + resultThree)
}

func coolCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is cool", i)
		time.Sleep(time.Second)
	}
}

func isCool(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is cool"
}
