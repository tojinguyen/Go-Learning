package main

import "fmt"

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(msg string) {
	fmt.Println("[Console]", msg)
}
