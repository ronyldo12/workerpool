package wp

import (
	"sync"
)

//Task task interface. Every task need to have these methods
type Task interface {
	GetError() error
	DoWork()
}

// Pool a worker group
type Pool struct {
	Tasks       []Task
	concurrency int
	queueTasks  chan Task
	wg          sync.WaitGroup
}

//AddTask add to add a new task in the pool
func (p *Pool) AddTask(t Task) {
	p.Tasks = append(p.Tasks, t)
}

// NewPool create a new pool
func NewPool(concurrency int) *Pool {
	return &Pool{
		concurrency: concurrency,
		queueTasks:  make(chan Task),
	}
}

// Exec create works and execute all tasks on the workers and wait all of tasks finish
func (p *Pool) Exec() {
	for i := 0; i < p.concurrency; i++ {
		//creating workers to receive and execute the tasks
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		//tasks are added in the channel.
		//workers are listening this channel and when a worker is idle
		//it will receive the task to execute
		p.queueTasks <- task
	}

	// close the channel when all task was executed
	close(p.queueTasks)

	p.wg.Wait()
}

// The worker execute tasks from the channel
func (p *Pool) work() {
	for task := range p.queueTasks {
		//when the work is idle it receive a new task from the channel
		task.DoWork()
		p.wg.Done()
	}
}
