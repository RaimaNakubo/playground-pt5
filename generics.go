package main

import (
	"fmt"
)

func main() {

	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) //compairing a slice of ints to int with function Index()

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) //compairing a slice of strings to string with the same function Index()
}

// function Index returns index of x value inside s slice, or -1 if x value have not been found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s { //v and x both have type T, type T have costraint: "comparable",
		if v == x { //so operator == can be used inside this function
			return i
		}
	}
	return -1
}
