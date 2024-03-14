package main

import "fmt"

func main() {
	TotalMsg := 0
	newmsg := 5
	TotalMsg = incrementMsg(TotalMsg, newmsg)
	fmt.Printf("Total %d new messages received", TotalMsg)
}

func incrementMsg(value1, value2 int) int {
	return value1 + value2
}
