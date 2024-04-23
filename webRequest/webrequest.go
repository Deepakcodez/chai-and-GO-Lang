package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Print("web request in GO")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if(err != nil){
		fmt.Println("something error in fetching data")
		return
	}
	defer resp.Body.Close() 
	// defer means this line run at the end of the main function for sure
	// we need to close the response for optimization

	// fmt.Println(resp)    this we give unreadable formate of response 

	data, err := ioutil.ReadAll(resp.Body)
    //data in the form of binary we have to convert into string
	if(err != nil){
		fmt.Println("Error in reading fetched data")
	}

	fmt.Print(string(data))   

}