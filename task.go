package wp

import "sync"

//ITask task interface. Every task need to have these methods
type ITask interface {
	DoWork(wg *sync.WaitGroup)
	GetError() error
	GetID() string
}

//Task task example
type Task struct {
	Err error
	ID  string
}

//DoWork this func will be called by pool do exec the job task
//it's very very import do call wg.Done() when job task end
func (t *Task) DoWork(wg *sync.WaitGroup) {
	wg.Done()
}

//GetError if some erro happen during the task execution you can return here
func (t *Task) GetError() error {
	return t.Err
}

//GetID return id of task
func (t *Task) GetID() string {
	return t.ID
}
