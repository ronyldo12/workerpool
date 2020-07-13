package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	wp "github.com/ronyldo12/workerpool"
)

//MyAditionalDataToTask adicional data to task
type MyAditionalDataToTask struct {
	ID          string
	Description string
}

//MyTask task example
type MyTask struct {
	MyAditionalDataToTask MyAditionalDataToTask
	Err                   error
}

//DoWork this func will be called by pool do exec the job task
func (t *MyTask) DoWork() {
	fmt.Printf("Start execution: %s \n", t.MyAditionalDataToTask.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.MyAditionalDataToTask.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTask) GetError() error {
	return t.Err
}

func main() {

	//number of task will be executed in the same time
	concurrency := 3

	pool := wp.NewPool(concurrency)
	for i := 1; i <= 20; i++ {
		//create a task
		task := &MyTask{
			MyAditionalDataToTask: MyAditionalDataToTask{ID: strconv.Itoa(i)},
		}
		//add task in the pool
		pool.AddTask(task)
	}
	pool.Exec()

	for _, task := range pool.Tasks {

		switch t := task.(type) {
		case MyTask:
			if task.GetError() != nil {
				fmt.Printf("%s -> %v", task.GetError())
			}
		}

	}
}
