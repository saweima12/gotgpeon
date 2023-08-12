package main

import (
	"fmt"
)

func main() {
	var list []string

	for item := range list {
		fmt.Println(item)
	}
}
