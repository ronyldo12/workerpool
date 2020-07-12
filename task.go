package wp

//ITask task interface. Every task need to have these methods
type ITask interface {
	GetError() error
	GetEntity() interface{}
	DoWork()
}

//GerenricEntity empty entity
type GerenricEntity struct {
	UD string
}

//Task task example
type Task struct {
	GerenricEntity GerenricEntity
	Err            error
}

//DoWork this func will be called by pool do exec the job task
//it's very very import do call wg.Done() when job task end
func (t *Task) DoWork() {
}

//GetError if some erro happen during the task execution you can return here
func (t *Task) GetError() error {
	return t.Err
}

//GetEntity return entity
func (t *Task) GetEntity() interface{} {
	return t.GerenricEntity
}
