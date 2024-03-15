package main

import "fmt"

type messageToSend struct {
	message  string
	sender   User
	receiver User
}

type User struct {
	name   string
	number int
}

func canSendMessage(msgToSend messageToSend) bool {

	if msgToSend.message == "" {
		return false
	}
	if msgToSend.sender == (User{}) {
		return false
	}
	if msgToSend.receiver == (User{}) {
		return false
	}
	 
	return true
}

func main() {

	fmt.Print("Struct in Go lang")
	user1 := User{name: "Deepak", number: 6749377}
	user2 := User{name: "Muskan", number: 787897979}

	msg := messageToSend{
		message: "Hello",
		sender: user1,
		receiver: user2,
	}

	fmt.Printf("message is :",canSendMessage(msg))

}
