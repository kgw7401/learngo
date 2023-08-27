package main

import (
	"fmt"

	mydict "github.com/kgw7401/learngo/mydict"
)

// // Account
// func main() {
// 	account := accounts.NewAccount("geonwoo")
// 	account.Deposit(10)
// 	fmt.Println(account.Balance())
// 	err := account.WithDraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.Balance(), account.Owner())
// 	fmt.Println(account)
// }

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	// // Search
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// // Add
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search("hello")
	// fmt.Println("found", word, "definition:", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// // Update
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	// Delete
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
