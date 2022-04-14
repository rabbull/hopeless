package future

type Scheduler interface {
	Launch(func())
}

var DefaultScheduler Scheduler = NewNativeScheduler()

type NativeScheduler struct{}

func NewNativeScheduler() Scheduler {
	return &NativeScheduler{}
}

func (s *NativeScheduler) Launch(fn func()) {
	go fn()
}
