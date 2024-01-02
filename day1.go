package main

import (
	"fmt"
)

func main() {
	var a [3]int = [3]int{1, 2, 3}
	//print the indices and values
	for i, v := range a {
		fmt.Printf("The value %d is at %d\n", v, i)
	}
	fmt.Printf("The value in the array are: ")
	for _, v := range a {
		fmt.Printf("%d , ", v)
	}
	// usage of ...
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	// Another example of the usage of the ...
	r := [...]int{1: -1, 2: -2}
	for _, k := range r {
		fmt.Println(k)
	}
	r[2] = 25
	for _, k := range r {
		fmt.Println(k)
	}
}
