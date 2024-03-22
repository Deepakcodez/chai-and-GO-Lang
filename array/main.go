package main

import "fmt"

func main() {

	
	var names [3]string                    // way 2

	names[0] = "Deepak"
	names[1] = "Musu"
	names[2] = "Abhi"


	fmt.Println("lenght of name list is: ", len(names))
	fmt.Println(" Name list is: ", names)
}
