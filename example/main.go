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
	return future.New(func() (Foo, error) {
		time.Sleep(time.Millisecond * 13)

		log.Printf("foo loaded: now=%v", now())
		return "foo", nil
	})
}

func LoadBar() future.Future[Bar] {
	log.Printf("load bar: now=%v", now())
	return future.New(func() (Bar, error) {
		time.Sleep(time.Millisecond * 7)

		log.Printf("bar loaded: now=%v", now())
		return math.Pi, nil
	})
}

func PackFooBar() future.Future[*FooBar] {
	log.Printf("pack foobar: now=%v", now())
	return future.Then(future.Bind(LoadFoo(), LoadBar()), func(tuple *future.Tuple[Foo, Bar], err error) (*FooBar, error) {
		if err != nil {
			return nil, err
		}
		return &FooBar{
			Foo: tuple.A,
			Bar: tuple.B,
		}, nil
	})
}

func main() {
	foobar, err := PackFooBar().Wait()
	if err != nil {
		log.Fatalf("failed to pack foobar: err=%v", err)
	}
	bi, _ := json.Marshal(foobar)
	log.Printf("foobar packed: foobar=%v", string(bi))
}
