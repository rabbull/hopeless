package hopeless

import (
	"sync"
	"time"
)

type Future[T any] interface {
	Wait() Result[T]
	WaitWithTimeout(t time.Duration) Result[T]
}

type futureImpl[T any] struct {
	result Result[T]

	scheduler Scheduler
	wg        sync.WaitGroup
}

func (f *futureImpl[T]) recover() {
	if err := recover(); err != nil {
		if err, ok := err.(error); ok {
			f.result = Err[T](err)
		} else {
			f.result = Err[T](ErrUnknownPanic)
		}
	}
}

func (f *futureImpl[T]) Wait() Result[T] {
	f.wg.Wait()
	return f.result
}

func (f *futureImpl[T]) WaitWithTimeout(t time.Duration) Result[T] {

	res := make(chan Result[T])
	defer close(res)

	go func() {
		defer f.recover()
		res <- f.Wait()
	}()

	select {
	case x := <-res:
		return x
	case <-time.After(t):
		return Err[T](ErrTimeout)
	}
}

func New[T any](job func() Result[T]) Future[T] {
	return NewWithScheduler(DefaultScheduler, job)
}

func NewWithScheduler[T any](scheduler Scheduler, job func() Result[T]) Future[T] {
	future := futureImpl[T]{
		scheduler: scheduler,
		wg:        sync.WaitGroup{},
	}

	future.wg.Add(1)
	scheduler.Launch(func() {

		// release lock
		defer future.wg.Done()

		// panic recovery
		defer future.recover()

		future.result = job()
	})

	return &future
}

func Then[T, S any](f Future[T], handler func(Result[T]) Result[S]) Future[S] {
	return New(func() Result[S] {
		return handler(f.Wait())
	})
}

func Join[T, S any](t Future[T], s Future[S]) Future[Tuple[T, S]] {
	return New(func() Result[Tuple[T, S]] {
		return JoinResult(t.Wait(), s.Wait())
	})
}
