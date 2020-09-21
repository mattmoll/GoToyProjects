package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello, world!")

	var x int = 5
	var y int = 7
	fmt.Println(x +y)

	// Infers type
	a := 12
	fmt.Println(a)

	if x < 6{
		fmt.Println("I know ifs in go lol")
	}

	var array[5]int
	fmt.Println(array)

	anotherArray := [5]int{1,4,8,12,65}
	fmt.Println(anotherArray)

	// Slices are an abstraction on top of arrays
	slice := []int{1,2,6}

	// Append is immutable, returns a new one.
	slice = append(slice, 13)

	myDictionary := make(map[string]bool)
	myDictionary["map sintax is weird"] = true
	myDictionary["but I'm having fun"] = true
	fmt.Println(myDictionary)
}