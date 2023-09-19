package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var list []string
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Print("Weeed")
	for item := range list {
		fmt.Println(item)
	}
}
