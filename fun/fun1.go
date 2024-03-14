package main

import "fmt"

func addString(arg1 string, arg2 string) string {

	return arg1 + arg2
}

func main() {

	fmt.Println(addString("Hello", " Go"))
}