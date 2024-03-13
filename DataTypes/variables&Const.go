package main

import "fmt"

func main() {

	fmt.Println("Data types in Go Lang")

	// integer
	 var sallery int = 3000;
	 fmt.Print(sallery)                        // println use for print only  variable
	 fmt.Printf("My sallery is : %d \n  ",sallery)   // printf used for print line
	 fmt.Printf("My sallery is : %T \n",sallery)   //  printf used for print line


	 //string
	 var name string = "Deepak";
	 fmt.Print(name);
	 fmt.Printf("My name is : %s \n", name)

     //bool
	 var isLoggin bool = false;
	 fmt.Print(isLoggin);
	 fmt.Printf("Is am login : %t \n", isLoggin)

	 //float
	 var salary float32 = 1000.08980;
     fmt.Print(salary);
     fmt.Printf("My salary is : %f \n", salary)


	 //const
	 const myConst string = "secret const";
	//  also
	 const (
		friend string = "preet";
		friend2 string = "Abhi";
	 )
        
     
     fmt.Printf("constant is : %s \n", myConst)
     fmt.Printf("friend is : %s \n", friend)
	 fmt.Printf("friend2 is : %s \n", friend2)


}