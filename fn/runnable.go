package fn

type Runnable interface {
	Run()
	CheckedRun() error
}

type RunnableFn func()

func (fn RunnableFn) Run() {
	fn()
}

func (fn RunnableFn) CheckedRun() error {
	fn.Run()
	return nil
}

type RunnableCheckedFn func() error

func (fn RunnableCheckedFn) Run() {
	_ = fn.CheckedRun()
}

func (fn RunnableCheckedFn) CheckedRun() error {
	return fn()
}

func RunnableOf(fn func()) Runnable {
	return RunnableFn(fn)
}

func RunnableCast(fn func() error) Runnable {
	return RunnableCheckedFn(fn)
}
