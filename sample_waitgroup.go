package main

import (
	"fmt"
	"github.com/yasarix/workergo"
	"sync"
	"time"
)

type Task struct {
	number  int
	message string
}

var wg sync.WaitGroup

func main() {
	d := workergo.NewDispatcherWG(2, 5, &wg)
	d.Run()
	defer d.Stop()

	wg.Add(1)
	go func() {
		fmt.Println("Sending tasks...")
		for i := 0; i < 20; i++ {
			fmt.Println("Sending", i)
			job := workergo.NewJob(workergo.TASK, NewTask(i, "Hello"), "DoWork")
			d.SubmitJob(job)
			fmt.Println("Job ID:", job.ID)
		}

		fmt.Println("Tasks sent.")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All tasks are done")
}

func NewTask(number int, message string) *Task {
	return &Task{
		number:  number,
		message: message,
	}
}

func (m *Task) DoWork() {
	fmt.Println(m.message, " -> ", m.number)
	// Sleep 1 second
	time.Sleep(time.Second * time.Duration(1))
}
