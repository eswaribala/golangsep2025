package main

import (
	"log"
	"net/http"
)

func acceLink(link string, c chan string) {
	println("Accesing link", link)
	response, err := http.Get(link)
	if err != nil {
		println("Error:", err.Error())
		c <- "Error"
		return
	} else {
		c <- "Success"
		println("Status code:", response.StatusCode)
	}

}

func main() {

	link := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.reddit.com"}

	//create channel
	c := make(chan string)
	for _, l := range link {
		go acceLink(l, c)
	}

	for i := 0; i < len(link); i++ {
		//receive from channel
		result := <-c
		log.Println("Result:", result)
	}
	//no of messages exceeding the no of receivers
	// deadlock will occur
	// so we need to have equal no of receivers
	//or we can use buffered channel
	//result1 := <-c
	//log.Println("Result:", result1)

	log.Println("End of the program")

}
