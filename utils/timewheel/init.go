package timewheel

import (
	"fmt"
	"time"

	"github.com/jhunters/timewheel"
)

var wheel *timewheel.TimeWheel[Runnable]

func GetTimingWheel() *timewheel.TimeWheel[Runnable] {
	return wheel
}

func Init() (result *timewheel.TimeWheel[Runnable], err error) {
	wheel, err = timewheel.New[Runnable](time.Second/2, 60)
	if err != nil {
		fmt.Printf("Timewheel initialize err: %s\n", err.Error())
		return nil, err
	}

	wheel.Start()
	return wheel, nil
}

func AddTask(delay time.Duration, taskObj Runnable) (err error) {
	newTask := timewheel.Task[Runnable]{
		Data: taskObj,
		TimeoutCallback: func(t timewheel.Task[Runnable]) {
			t.Data.Run()
		},
	}
	_, err = wheel.AddTask(delay, newTask)
	if err != nil {
		return err
	}

	return nil
}
