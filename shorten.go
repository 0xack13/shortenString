package main

import (
	"fmt"
	"strings"
)

func shortenString(message string) func() string {
	return func() string {
		messageSlice := strings.Split(message, " ")
		fmt.Println(messageSlice)
		wordLength := len(messageSlice)
		if wordLength < 1 {
			return "Nothing left"
		} else {
			messageSlice = messageSlice[:(wordLength-1)]
			message = strings.Join(messageSlice, " ")
			return message
		}
	}
}

func main() {
	myString := shortenString("Welcome to concurrency in go! ...")
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

