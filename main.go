package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.golang.org",
		"http://www.stackoverflow.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checklink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		go func(link string) { // function literal
			time.Sleep(3 * time.Second)
			checklink(link, c)
		}(l) //parenthesis to call the function literal

	}

}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "might be down I think!"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "yep it is up"
	c <- link
}
