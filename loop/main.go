package main

import "fmt"

func main() {

	fmt.Print("Loop in Go lang")

	names := []string{"Deepak", "musu", "abhi", "nimit", "tony"}

	for i := 0; i <= 4; i++ {
		fmt.Print(names[i], "\n")
	}

	for i := range names {
		fmt.Print(names[i], "\n")
	}

	for index, data := range names {
		fmt.Println("index", index, "data", data)
	}

	myValue := 0
	//like while loop
	for myValue < 10 {
		print(myValue)
		myValue++
	}

}
