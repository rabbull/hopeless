package hopeless

import (
	"sync"
)

type Future[T any] interface {
	Wait() (T, error)
}

type futureImpl[T any] struct {
	payload T
	err     error

	scheduler Scheduler
	wg        sync.WaitGroup
}

func (f *futureImpl[T]) Wait() (T, error) {
	f.wg.Wait()
	return f.payload, f.err
}

func New[T any](job func() (T, error)) Future[T] {
	return NewWithScheduler(DefaultScheduler, job)
}

func NewWithScheduler[T any](scheduler Scheduler, job func() (T, error)) Future[T] {
	future := futureImpl[T]{
		err:       nil,
		scheduler: scheduler,
		wg:        sync.WaitGroup{},
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

func Then[T, S any](f Future[T], handler func(T, error) (S, error)) Future[S] {
	return New(func() (S, error) {
		return handler(f.Wait())
	})
}

type Tuple[T, S any] struct {
	A T
	B S
}

func Bind[T, S any](t Future[T], s Future[S]) Future[*Tuple[T, S]] {
	return New(func() (*Tuple[T, S], error) {

		t, err := t.Wait()
		if err != nil {
			return nil, err
		}

		s, err := s.Wait()
		if err != nil {
			return nil, err
		}

		return &Tuple[T, S]{
			A: t,
			B: s,
		}, nil
	})
}
