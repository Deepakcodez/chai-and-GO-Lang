package main

import "fmt"

func main() {

	fmt.Print("pointers in GO")

	myNumber := 45

	var ptr = &myNumber;
	
	fmt.Println("Adress storing in ptr",ptr)
	fmt.Println("Actual value store in ptr",*ptr)

	//changing value of myNumber by pointer

	*ptr = *ptr + 5;
	//actual value =  actual value + 5 ;

    fmt.Println("new value of myNumber is :", myNumber);

}
