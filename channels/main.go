package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.amazon.com",
		"https://go.dev",
	}

	// Create channel
	c := make(chan string)

	for _, link := range links {
		// Create child routines
		go checkLink(link, c)
	}

	//for {
	//	//fmt.Println(<- c)
	//	go checkLink(<-c, c)
	//}

	//for l := range c {
	//	go checkLink(l, c)
	//}

	// Sleep for 5 seconds before making another call
	// Using function literal
	for l := range c {
		go func(link string) {
			// Sleep for 5 second
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		//c <- "Might be down I think!"
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	//c <- "Yup, its up !"
	c <- link
}
