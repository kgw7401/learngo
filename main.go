package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

// // URL Checker
// type result struct {
// 	url    string
// 	status string
// }

// func main() {
// 	results := make(map[string]string)
// 	c := make(chan result)
// 	urls := []string{"https://www.naver.com/", "https://www.google.com/", "https://www.amazon.com/", "https://www.facebook.com/", "https://www.instagram.com/"}
// 	for _, url := range urls {
// 		go hitURL(url, c)
// 	}
// 	for i := 0; i < len(urls); i++ {
// 		result := <-c
// 		results[result.url] = result.status
// 	}

// 	for url, status := range results {
// 		fmt.Println(url, status)
// 	}
// }

// func hitURL(url string, c chan<- result) { // chac<- : send only channel
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	status := "OK"
// 	if err != nil || resp.StatusCode >= 400 {
// 		status = "FAILED"
// 	}
// 	c <- result{url: url, status: status}
// }

// Job Scrapper
type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=data%20engineer"

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // defer : function이 끝나고 무조건 실행

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, mainC chan<- []extractedJob) {
	jobs := []extractedJob{}
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // defer : function이 끝나고 무조건 실행

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // defer : function이 끝나고 무조건 실행

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages, _ = fmt.Println(s.Find("a").Length())
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
