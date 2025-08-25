package main

func main() {
	var logger Logger = &ConsoleLogger{}
	logger.Log("Hello, World!")
}
