package future

type Scheduler interface {
	Dispatch(func())
}

type NativeScheduler struct{}

func NewNativeScheduler() Scheduler {
	return &NativeScheduler{}
}

func (s *NativeScheduler) Dispatch(fn func()) {
	go fn()
}
