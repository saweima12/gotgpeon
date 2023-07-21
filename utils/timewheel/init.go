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

func Init() (wheel *timewheel.TimeWheel[Delayable], err error) {

	wheel, err = timewheel.New[Delayable](time.Second, 60)
	if err != nil {
		fmt.Println("Timewheel initialize err: %s", err.Error())
		return nil, err
	}

	return wheel, nil
}

func AddTask(delay time.Duration, taskObj Delayable) (taskId uint64, err error) {
	newTask := timewheel.Task[Delayable]{
		Data: taskObj,
		TimeoutCallback: func(t timewheel.Task[Delayable]) {
			t.Data.Run()
		},
	}
	nid, err := wheel.AddTask(delay, newTask)
	if err != nil {
		return 0, err
	}

	taskId = uint64(nid)
	return taskId, nil
}
