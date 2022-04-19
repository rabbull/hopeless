package main

import (
	"encoding/json"
	"log"
	"math"
	"time"

	future "github.com/rabbull/hopeless"
)

func now() int64 {
	return time.Now().UnixNano()
}

type Foo string

type Bar float64

type FooBar struct {
	Foo Foo `json:"foo"`
	Bar Bar `json:"bar"`
}

func LoadFoo() future.Future[Foo] {
	log.Printf("load foo: now=%v", now())
	return future.New(func() future.Result[Foo] {
		time.Sleep(time.Millisecond * 13)

		log.Printf("foo loaded: now=%v", now())
		return future.Ok(Foo("foo"))
	})
}

func LoadBar() future.Future[Bar] {
	log.Printf("load bar: now=%v", now())
	return future.New(func() future.Result[Bar] {
		time.Sleep(time.Millisecond * 7)

		log.Printf("bar loaded: now=%v", now())
		return future.Ok(Bar(math.Pi))
	})
}

func PackFooBar() future.Future[*FooBar] {
	log.Printf("pack foobar: now=%v", now())
	return future.Then(
		future.Bind(LoadFoo(), LoadBar()),
		func(res future.Result[*future.Tuple[Foo, Bar]]) future.Result[*FooBar] {
			if res.Err() != nil {
				return future.Err[*FooBar](res.Err())
			}

			return future.Ok(&FooBar{
				Foo: res.Val().A,
				Bar: res.Val().B,
			})
		},
	)
}

func main() {
	foobar, err := PackFooBar().Wait()
	if err != nil {
		log.Fatalf("failed to pack foobar: err=%v", err)
	}
	bi, _ := json.Marshal(foobar)
	log.Printf("foobar packed: foobar=%v", string(bi))
}
