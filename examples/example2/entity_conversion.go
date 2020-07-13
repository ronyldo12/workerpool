package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	wp "github.com/ronyldo12/workerpool"
)

//MyEntity my entity
type MyEntity struct {
	ID string
}

//MyTask task example
type MyTask struct {
	Entity MyEntity
	Err    error
	ID     string
}

//DoWork this func will be called by pool do exec the job task
//it's very very import do call wg.Done() when job task end
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

//GetEntity return id of task
func (t *MyTask) GetEntity() interface{} {
	return t.Entity
}

//GetEntityFromTask return MyEntity from task
func GetEntityFromTask(t wp.ITask) MyEntity {
	entity := t.GetEntity()
	return entity.(MyEntity)
}

func main() {

	//number of task will be executed in the same time
	concurrency := 3

	pool := wp.NewPool(concurrency)
	for i := 1; i <= 5; i++ {
		//entity
		entity := MyEntity{
			ID: "Task " + strconv.Itoa(i),
		}
		//create a task
		task := &MyTask{ID: "TASK" + strconv.Itoa(i), Entity: entity}
		//add task in the pool
		pool.AddTask(task)
	}
	pool.Exec()

	for _, task := range pool.Tasks {
		//so you can get the entity
		myEntity := GetEntityFromTask(task)
		fmt.Printf("Entity: %v \n", myEntity.ID)
		if task.GetError() != nil {
			fmt.Printf("%v", task.GetError())
		}
	}
}
