package main

import (
	"fmt"
	"github.com/heekng/study_go/accounts"
)

func main() {
	//account := accounts.Account{Owner: "heekng"}
	//fmt.Println(account)
	//account.Owner = "pepe"
	//fmt.Println(account)

	account := accounts.NewAccount("heekng")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	//err := account.Withdraw(20)
	//if err != nil {
	//log.Fatalln(err)
	//fmt.Println(err)
	//}
	fmt.Println(account.Balance(), account.Owner())
}
