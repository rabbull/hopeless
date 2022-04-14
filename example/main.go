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
	return future.New(func() (*FooBar, error) {
		tuple, err := future.Bind(LoadFoo(), LoadBar()).Wait()
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
	PackFooBar().Then(func(foobar *FooBar) {
		bi, _ := json.MarshalIndent(foobar, "", "\t")
		log.Printf("foobar packed: foobar=%v, now=%v", string(bi), now())
	})

	log.Print("we have no hope")
	time.Sleep(time.Second)
}
