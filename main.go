package main

import (
	"fmt"
	"net/http"
)

func main() {
	cchan := make(chan int, 100)

	for i := 0; i < 100; i++ {
		cchan <- i
	}

	go processUpdate(cchan)
	http.ListenAndServe(":8000", nil)
}

func processUpdate(ch chan int) {
	for update := range ch {
		fmt.Println(update)
	}
}
