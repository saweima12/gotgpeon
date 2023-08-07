package timewheel

import (
	"fmt"
	"time"

	"github.com/jhunters/timewheel"
)

var wheel *timewheel.TimeWheel[Delayable]

func GetTimingWheel() *timewheel.TimeWheel[Delayable] {
	return wheel
}

func Init() (result *timewheel.TimeWheel[Delayable], err error) {
	wheel, err = timewheel.New[Delayable](time.Second/2, 60)
	if err != nil {
		fmt.Printf("Timewheel initialize err: %s\n", err.Error())
		return nil, err
	}

	wheel.Start()
	return wheel, nil
}

func AddTask(delay time.Duration, taskObj Delayable) (err error) {
	newTask := timewheel.Task[Delayable]{
		Data: taskObj,
		TimeoutCallback: func(t timewheel.Task[Delayable]) {
			t.Data.Run()
		},
	}
	_, err = wheel.AddTask(delay, newTask)
	if err != nil {
		return err
	}

	return nil
}
