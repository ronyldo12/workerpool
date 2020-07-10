package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	wp "github.com/ronyldo12/workpool"
)

//MyTaskTypeOne task example
type MyTaskTypeOne struct {
	Err error
	ID  string
}

//DoWork this func will be called by pool do exec the job task
//it's very very import do call wg.Done() when job task end
func (t *MyTaskTypeOne) DoWork(wg *sync.WaitGroup) {
	fmt.Printf("Start execution: %s \n", t.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.ID)
	wg.Done()
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTaskTypeOne) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *MyTaskTypeOne) GetID() string {
	return t.ID
}

//MyTaskTypeTwo task example
type MyTaskTypeTwo struct {
	Err error
	ID  string
}

//DoWork this func will be called by pool do exec the job task
//it's very very import do call wg.Done() when job task end
func (t *MyTaskTypeTwo) DoWork(wg *sync.WaitGroup) {
	fmt.Printf("Start execution task type two: %s \n", t.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem task type two")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution task type two: %s \n", t.ID)
	wg.Done()
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTaskTypeTwo) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *MyTaskTypeTwo) GetID() string {
	return t.ID
}

func main() {

	//number of task will be executed in the same time
	concurrency := 20

	pool := wp.NewPool(concurrency)
	for i := 1; i <= 10; i++ {
		//create a task
		task := &MyTaskTypeOne{ID: "TASK_TYPE_ONE_" + strconv.Itoa(i)}
		//add task in the pool
		pool.AddTask(task)
	}
	for i := 1; i <= 5; i++ {
		//create a task
		task := &MyTaskTypeTwo{ID: "TASK_TYPE_TWO_" + strconv.Itoa(i)}
		//add task in the pool
		pool.AddTask(task)
	}
	pool.Run()

	var numErrors int
	for _, task := range pool.Tasks {
		if task.GetError() != nil {
			fmt.Printf("%v", task.GetError())
			numErrors++
		}
	}
}