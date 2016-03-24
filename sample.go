package main

import (
	"fmt"
	"github.com/yasarix/workergo"
)

type Task struct {
	number  int
	message string
}

func main() {
	d := workergo.NewDispatcher(3, 10)
	d.Run()
	defer d.Stop()

	for i := 0; i < 20; i++ {
		fmt.Println("Sending", i)
		job := workergo.NewJob(workergo.TASK, NewTask(i, "Hello"), "DoWork")
		d.SubmitJob(job)
		fmt.Println("Job ID:", job.ID)
	}

	// Infinite loop to wait until all go routines finished
	for {
	}
}

func NewTask(number int, message string) *Task {
	return &Task{
		number:  number,
		message: message,
	}
}

func (m *Task) DoWork() {
	fmt.Println(m.message, " -> ", m.number)
}
