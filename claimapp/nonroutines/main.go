package main

import (
	"log"
	"net/http"
)

func acceLink(link string) {
	println("Accesing link", link)
	response, err := http.Get(link)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("Status code:", response.StatusCode)
}

func main() {

	link := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.reddit.com"}
	for _, l := range link {
		acceLink(l)
	}

	log.Println("End of the program")

}
