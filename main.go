package main

import (
	"fmt"
	"net/http"
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

// func main() {
// 	dictionary := mydict.Dictionary{"first": "First word"}
// 	// Search
// 	definition, err := dictionary.Search("second")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(definition)
// 	}

// 	// Add
// 	word := "hello"
// 	definition := "Greeting"
// 	err := dictionary.Add(word, definition)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	hello, _ := dictionary.Search("hello")
// 	fmt.Println("found", word, "definition:", hello)
// 	err2 := dictionary.Add(word, definition)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// 	// Update
// 	baseWord := "hello"
// 	dictionary.Add(baseWord, "First")
// 	err := dictionary.Update(baseWord, "Second")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	word, _ := dictionary.Search(baseWord)
// 	fmt.Println(word)

// 	// Delete
// 	baseWord := "hello"
// 	dictionary.Add(baseWord, "First")
// 	dictionary.Search(baseWord)
// 	dictionary.Delete(baseWord)
// 	word, err := dictionary.Search(baseWord)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(word)
// }

// // Go Routines & Channels
// func main() {
// 	c := make(chan string)
// 	people := [2]string{"geonwoo", "minwoo"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	fmt.Println("Waiting for messages...")
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("Receive this message:", <-c)
// 	}
// 	// fmt.Println(<-c) // Error! Becuase there is no more channel
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + " is sexy"
// }

// URL Checker
type result struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{"https://www.naver.com/", "https://www.google.com/", "https://www.amazon.com/", "https://www.facebook.com/", "https://www.instagram.com/"}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- result) { // chac<- : send only channel
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
