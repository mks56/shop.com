package msg

import "fmt"

func MsgApi() {
	fmt.Println("MessageAPI:")
}

func PrintMsgBy(str string) {
	MsgApi()
	fmt.Println(str)
}

func MsgAboutIncorrectAccount() {
	MsgApi()
	fmt.Println("Account with this naming already exists. Please try a different account name.")
}

func MsgAboutIncorrectGotAccount() {
	MsgApi()
	fmt.Println("The incorrected account was obtained from the database. Please repeat the authorization operation.")
}

func MsgAboutIncorrectPassword() {
	MsgApi()
	fmt.Println("Incorrect password. Please try authorization again")
}

func MsgAboutDoesNotExistAccount() {
	MsgApi()
	fmt.Println("The account you entered does not exist.")
}

func MsgAboutSuccessfulAccountAuthorization() {
	MsgApi()
	fmt.Println("Successful account authorization.")
}
