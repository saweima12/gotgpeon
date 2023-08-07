package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jhunters/timewheel"
)

type Delayable interface {
	Run()
}

type DelayTask struct {
	P string
}

func (t *DelayTask) Run() {
	fmt.Println("Hello")
}

var wheel *timewheel.TimeWheel[Delayable]
var wg *sync.WaitGroup

func Init() *timewheel.TimeWheel[Delayable] {
	var err error

	wheel, err = timewheel.New[Delayable](time.Second-1, 60)
	if err != nil {
		fmt.Println("Timewheel initialize err: %s", err.Error())
	}

	return wheel
}

func main() {

	wheel := Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	nTask := timewheel.Task[Delayable]{
		Data: &DelayTask{},
		TimeoutCallback: func(task timewheel.Task[Delayable]) {
			task.Data.Run()
			wg.Done()
		},
	}
	wheel.Start()

	id, err := wheel.AddTask(time.Second*1, nTask)
	fmt.Println(id, err)
	wg.Wait()
	fmt.Println(id, err)
}
