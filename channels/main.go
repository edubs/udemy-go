package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// 'watch' the channel c, whenever a value comes out of it, assign it to the value l
	// once l has a value, the body of the for loop is executed immediately
	for l := range c {
		// pass l to the function literal, creates a 'copy' which the function has access to, freeing up the memory
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send a message into our channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send a message into our channel
	c <- link
}
