package main

import (
	"fmt"
)

func main()  {
	
	mycar := struct{
		colour string
	    price int   } {
			colour : "red",
			price : 20000,
		}
	

	fmt.Printf("car colour is %s ",mycar.colour)

}