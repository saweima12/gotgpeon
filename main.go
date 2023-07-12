package main

import (
	"fmt"

	"github.com/liuzl/gocc"
)

func main() {
	s2t, err := gocc.New("s2t")
	if err != nil {
		fmt.Print(err)
	}

	newStr, err := s2t.Convert("台湾")
	fmt.Println(newStr)
}
