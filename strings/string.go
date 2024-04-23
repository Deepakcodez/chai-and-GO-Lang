package main

import (
	"fmt"
	"strings"
)

func main() {
	//separate words from string by any separator
	data := "mango,banana,cherry"
	dataInParts := strings.Split(data, ",")
	fmt.Print(dataInParts,"\n")

	//count of any charcter/words in sentence/word
	str := "one two two three"
	countOfTwo := strings.Count(str,"two")
	fmt.Print(countOfTwo,"\n")

	//trim white spaces
	str = "   hello duniya       "
	res := strings.TrimSpace(str)
	fmt.Print(res,"\n")


	//join string
	str1 :=  "word1";
	str2 := "word2";
	result := strings.Join([]string{str1,str2}," ")  // array of  string , source, separator
	fmt.Print(result)
}