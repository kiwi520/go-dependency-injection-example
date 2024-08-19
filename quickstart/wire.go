//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

func InitializeEvent(msg string) (Event, error) {
	//wire.Build(NewEvent, NewGreeter, NewMessage)
	wire.Build(EventSet)
	return Event{}, nil
}

func InitializeDistinguishingTypes() DistinguishingTypes {
	return DistinguishingTypes{MyIntO: ProvideIntO(), MyIntT: ProvideIntT()}
}
