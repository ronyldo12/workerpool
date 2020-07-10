package wp

import (
	"sync"
)

// Pool is a worker group that runs a number of tasks at a
// configured concurrency.
type Pool struct {
	Tasks []ITask

	concurrency int
	queueTasks  chan ITask
	wg          sync.WaitGroup
}

//AddTask add to add a new task in the pool
func (p *Pool) AddTask(t ITask) {
	p.Tasks = append(p.Tasks, t)
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
// concurrency is the quantity of task will be executed in the same time
func NewPool(concurrency int) *Pool {
	return &Pool{
		concurrency: concurrency,
		queueTasks:  make(chan ITask),
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.queueTasks <- task
	}

	// all workers return
	close(p.queueTasks)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *Pool) work() {
	for task := range p.queueTasks {
		task.DoWork(&p.wg)
	}
}
