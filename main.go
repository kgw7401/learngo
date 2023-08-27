package main

import (
	"fmt"

	"github.com/kgw7401/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("geonwoo")
	account.Deposit(10)
	fmt.Println(account.Balance())
	// err := account.WithDraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
