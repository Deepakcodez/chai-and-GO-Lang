package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("example.txt")

	if err!=nil{
		fmt.Print("error in creating file","\n")
		return
	}
	defer file.Close();

	content := "this is demo text"

	byte,err := io.WriteString(file,content+"\n")
	if err!=nil{
		fmt.Print("error in writing file")
		return
	}
	fmt.Print(byte)  //this will give bytes present in file
}