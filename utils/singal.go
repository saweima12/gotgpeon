package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func HandleSingal() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-c:
				break
			}
		}
	}()

	return c
}
