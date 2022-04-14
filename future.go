package future

import "sync"

type Future[T any] interface {
	Then(func(res T)) Future[T]
	Catch(func(err error)) Future[T]
	Wait()
}

type futureImpl[T any] struct {
	payload T
	err     error

	wg sync.WaitGroup
}

func (f *futureImpl[T]) Then(do func(res T)) Future[T] {
	f.wg.Wait()
	if f.err == nil {
		do(f.payload)
	}
	return f
}

func (f *futureImpl[T]) Catch(do func(err error)) Future[T] {
	f.wg.Wait()
	if f.err != nil {
		do(f.err)
	}
	return f
}

func (f *futureImpl[T]) Wait() {
	f.wg.Wait()
}

func New[T any](job func() (T, error)) Future[T] {
	return NewWithScheduler(DefaultScheduler, job)
}

func NewWithScheduler[T any](scheduler Scheduler, job func() (T, error)) Future[T] {
	future := futureImpl[T]{
		err: nil,
		wg:  sync.WaitGroup{},
	}

	future.wg.Add(1)
	scheduler.Launch(func() {
		defer future.wg.Done()

		val, err := job()
		if err != nil {
			future.err = err
		} else {
			future.payload = val
		}
	})

	return &future
}
