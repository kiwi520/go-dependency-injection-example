package main

import (
	"fmt"
	"github.com/google/wire"
)

type Message struct {
	msg string
}

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{msg: msg}
}

// NewGreeter Greeter构造函数
func NewGreeter(msg Message) Greeter {
	return Greeter{Message: msg}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	greet := e.Greeter.Greet()
	fmt.Println(greet.msg)
}

func (g Greeter) Greet() Message {
	return g.Message
}

// EventSet Event通常是一起使用的一个集合，使用wire.NewSet进行组合
var EventSet = wire.NewSet(NewEvent, NewMessage, NewGreeter)

type IntO int
type IntT int

func ProvideIntO() IntO {
	return 1
}

func ProvideIntT() IntT {
	return 2
}

type DistinguishingTypes struct {
	MyIntO IntO
	MyIntT IntT
}

// var StructSet = wire.NewSet(ProvideIntO, ProvideIntT, wire.Struct(new(DistinguishingTypes), "MyIntO", "MyIntT"))
// 通过wire.Struct来指定那些字段要被注入到结构体中，如果是全部字段，也可以简写成：
var StructSet = wire.NewSet(ProvideIntO, ProvideIntT, wire.Struct(new(DistinguishingTypes), "*"))

func main() {
	event, err := InitializeEvent("hello world!")
	if err != nil {
		fmt.Println(err)
	}
	event.Start()

	types := InitializeDistinguishingTypes()
	fmt.Println(types.MyIntO)
}
