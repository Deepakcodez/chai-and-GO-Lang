package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Print("JSON in go\n")
	person1 := Person{Name:"Deepak", Age: 20, IsAdult: true}
	fmt.Println("person1 data is:", person1)



	//convert obj into json
	bytedata,err := json.Marshal(person1)

	if err != nil{
		fmt.Println("error in marshling",err)
	}
	fmt.Println("peron json : ",string(bytedata))


	// unmarshling (decoding)
    
	var persondata Person

	err = json.Unmarshal(bytedata, &persondata)  //output will store in person data

	fmt.Print("unmarshal person data:", persondata)
}
