# Workpool - Allow to create a pool of differents types of tasks and execute it concurrently

This workpool is simple to use and allow you creat differents type of taks and execute all of then in the same. To do that you need to creat tasks sctucts the is compatible with ITask interface.

Also, you can set the concurrency. It mean that you can configure how many tasks are executed in the same time.

## How it work?

The pool will create workers and when all of workers are busy another tasks will wait. When some worker became idle, it will receive the task and execute until all tasks be executed.

## Installation
To install this package, you need to setup your Go workspace.  The simplest way to install the library is to run:
```
$ go get github.com/ronyldo12/workerpool
```

## Example - Simple workpool
```go
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	wp "github.com/ronyldo12/workerpool"
)

//MyTask task example
type MyTask struct {
	Entity interface{}
	Err    error
	ID     string
}

//DoWork this func will be called by pool do exec the job task
func (t *MyTask) DoWork() {
	fmt.Printf("Start execution: %s \n", t.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTask) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *MyTask) GetID() string {
	return t.ID
}

//GetEntity task entity
func (t *MyTask) GetEntity() interface{} {
	return t.Entity
}

func main() {

	//number of task will be executed in the same time
	concurrency := 3

	pool := wp.NewPool(concurrency)
	for i := 1; i <= 20; i++ {
		//create a task
		task := &MyTask{ID: "TASK" + strconv.Itoa(i)}
		//add task in the pool
		pool.AddTask(task)
	}
	pool.Exec()

	for _, task := range pool.Tasks {
		if task.GetError() != nil {
			fmt.Printf("%v", task.GetError())
		}
	}
}

```

## Example - Entity conversion
You can use a func like this to covert task Entity in yout type of Entity
```go
//GetEntityFromTask return MyEntity from task
func GetEntityFromTask(t wp.ITask) MyEntity {
	entity := t.GetEntity()
	return entity.(MyEntity)
}
```

## Example Two - Different type of tasks running in the same pool
```go

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	wp "github.com/ronyldo12/workerpool"
)

//MyTaskTypeOne task example
type MyTaskTypeOne struct {
	Entity interface{}
	Err    error
	ID     string
}

//DoWork this func will be called by pool do exec the job task
func (t *MyTaskTypeOne) DoWork() {
	fmt.Printf("Start execution: %s \n", t.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution: %s \n", t.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTaskTypeOne) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *MyTaskTypeOne) GetID() string {
	return t.ID
}

//GetEntity task entity
func (t *MyTaskTypeOne) GetEntity() interface{} {
	return t.Entity
}

//MyTaskTypeTwo task example
type MyTaskTypeTwo struct {
	Entity interface{}
	Err    error
	ID     string
}

//DoWork this func will be called by pool do exec the job task
//it's very very import do call wg.Done() when job task end
func (t *MyTaskTypeTwo) DoWork() {
	fmt.Printf("Start execution task type two: %s \n", t.ID)
	secondDuration := rand.Intn(5)
	if secondDuration > 4 {
		t.Err = fmt.Errorf("Some problem task type two")
	}

	time.Sleep(time.Second * time.Duration(secondDuration))
	fmt.Printf("=> End execution task type two: %s \n", t.ID)
}

//GetError if some erro happen during the task execution you can return here
func (t *MyTaskTypeTwo) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *MyTaskTypeTwo) GetID() string {
	return t.ID
}

//GetEntity return id of task
func (t *MyTaskTypeTwo) GetEntity() interface{} {
	return t.Entity
}

func main() {

	//number of task will be executed in the same time
	concurrency := 4

	pool := wp.NewPool(concurrency)

	//add tasks typed MyTaskTypeOne
	for i := 1; i <= 50; i++ {
		//create a task
		task := &MyTaskTypeOne{ID: "TASK_TYPE_ONE_" + strconv.Itoa(i)}
		//add task in the pool
		pool.AddTask(task)
	}

	//add tasks typed MyTaskTypeTwo
	for i := 1; i <= 5; i++ {
		//create a task
		task := &MyTaskTypeTwo{ID: "TASK_TYPE_TWO_" + strconv.Itoa(i)}
		//add task in the pool
		pool.AddTask(task)
	}

	//exec MyTaskTypeOne and MyTaskTypeTwo together
	pool.Exec()

	for _, task := range pool.Tasks {
		if task.GetError() != nil {
			fmt.Printf("%v", task.GetError())
		}
	}
}


```