package main

import (
	"fmt"

	"github.com/panjf2000/ants/v2"
)

func main() {
	pool, err := ants.NewPool(10000)

	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < 10000; i++ {
		_ = pool.Submit(func() {
			fmt.Println(i)
		})
	}

	fmt.Println(ants.Running())

}
