package isa

type Aggregate interface {
	Execute(func(a Aggregate))
}

type CommandHandler interface {
}