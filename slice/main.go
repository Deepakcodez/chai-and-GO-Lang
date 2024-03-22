package main

import "fmt"

func main() {

	fmt.Printf("Slices in Go lang\n")

	var fruit = []string{"mango", "peach"}

	fmt.Println("list of fruit is : ",fruit)
	
	fruit = append(fruit, "orange")
	
	fmt.Println("list of fruit is : ",fruit)
}
