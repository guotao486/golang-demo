package task

type TaskProvider interface {
	Add(task *TaskStore) error
	Get() *TaskStore
	Run(task *TaskStore)
	Finish(task *TaskStore)
	Fail(task *TaskStore)
}

func ProviderRegister(name string, provider TaskProvider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := Provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	Provides[name] = provider
}
