package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/liuzl/gocc"
)

var ChPtn = regexp.MustCompile("[\u3400-\u4DBF\u4E00-\u9FFF\uF900-\uFAFF]")

func main() {
	s2t, err := gocc.New("s2t")
	if err != nil {
		fmt.Print(err)
	}

	chSlice := ChPtn.FindAllString("台湾是个好地方，就是交通很乱，人们都是怪物。", -1)
	chStr := strings.Join(chSlice, "")

	newStr, err := s2t.Convert(chStr)
	fmt.Println(newStr)
}
