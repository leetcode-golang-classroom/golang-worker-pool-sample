package work

type task struct {
	execute      func() error
	errorHandler func(error)
}

func NewTask(execute func() error, errorHandler func(error)) *task {
	return &task{
		execute:      execute,
		errorHandler: errorHandler,
	}
}

func (t *task) Execute() error {
	return t.execute()
}

func (t *task) OnError(err error) {
	t.errorHandler(err)
}
