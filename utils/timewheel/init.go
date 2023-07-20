package timewheel

import (
	"fmt"
	"time"

	"github.com/jhunters/timewheel"
)

var wheel *timewheel.TimeWheel[string]

func Init() *timewheel.TimeWheel[string] {
	var err error

	wheel, err = timewheel.New[func()](time.Second, 60)
	if err != nil {
		fmt.Println("Timewheel initialize err: %s", err.Error())
	}

	return wheel
}

func AddTask() {
	wheel.AddTask(5, timewheel.Task[string]{
    Data: "1234",
    TimeoutCallback: time.AfterFunc
  }})
}
