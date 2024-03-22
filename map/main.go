package main

import "fmt"

func main() {

	fmt.Print("Maps in Go lang\n")

	languages := make(map[string]string) // string sting is data type of key and value it is like object in js

	languages["js"] = "javaScript"
	languages["py"] = "python"

	fmt.Print(languages)
	fmt.Print("\n")

	//printing specific key
	fmt.Print(languages["js"])
	fmt.Print("\n")

	//delete any key
	delete(languages, "js")
	fmt.Println("updated map: ", languages)

}
