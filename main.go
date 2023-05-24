package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	websites := []string{
		"https://sr-rs.facebook.com/",
		"https://mondo.rs/",
		"https://www.narodnopozoriste.rs/",
		"https://atelje212.rs/",
		"https://www.blic.rs/",
		"https://www.youtube.com/",
		"https://novi.kupujemprodajem.com/",
		"https://www.polovniautomobili.com/",
		"https://www.lilly.rs/",
		"https://www.dm.rs/",
	}

	ch := make(chan string)

	wg.Add(10)

	for _, website := range websites {
		go checkSite(website, ch)
	}

	wg.Wait()

}

func checkSite(link string, ch chan string) {

	_, err := http.Get(link)
	wg.Done()
	
	if err!=nil {
		fmt.Println(link, "The site is unavailable")
		ch <- link
		return
	}

	fmt.Println(link, "The site is available!")
	ch <- link
	
}
