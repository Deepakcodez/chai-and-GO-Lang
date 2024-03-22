package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Printf("Slices methods in Go lang\n")

	//another way to initialize

	scores := make([]int, 4)

	scores[0] = 55
	scores[1] = 63
	scores[2] = 454
	scores[3] = 444
	fmt.Print(scores)
	fmt.Print("\n")

	fmt.Println("cut the slice : ", append(scores[1:3]))
	fmt.Print("\n")

    //sorting
    
	fmt.Println("checking sorted or no :",sort.IntsAreSorted(scores))
	
	sort.Ints(scores)
	fmt.Println("sort the slice",scores)

}
