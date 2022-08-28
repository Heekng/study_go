package main

import (
	"fmt"
	"github.com/heekng/study_go/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"
	//fmt.Println(dictionary)
	//fmt.Println(dictionary["first"])
	/*
		definition, err := dictionary.Search("first")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/
	/*
		word := "hello"
		definition = "Greeting"
		err = dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		hello, _ := dictionary.Search(word)
		fmt.Println("found", word, "definition:", hello)
		err2 := dictionary.Add(word, definition)
		if err2 != nil {
			fmt.Println(err2)
		}
	*/
	/*
		baseWord := "hello"
		dictionary.Add(baseWord, "First")
		err := dictionary.Update(baseWord, "Second")
		if err != nil {
			fmt.Println(err)
		}
		word, _ := dictionary.Search(baseWord)
		fmt.Println(word)
	*/

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
