package main

import "fmt"

func main() {

	fruits := [3]string{"peach", "mango"}  // way 1
	var names [3]string                    // way 2

	names[0] = "Deepak"
	names[1] = "Musu"
	names[2] = "Abhi"

	fmt.Println("fruit list is: ", fruits)
	fmt.Println("lenght of fruit list is: ", len(fruits))
	fmt.Println(" Name list is: ", names)
}
