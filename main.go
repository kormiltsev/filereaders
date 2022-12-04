package main

import "github.com/kormiltsev/filereaders/readers"

var Sources = map[string]string{
	"iwatchsleep": &readers.MiWatchSleep,
}

func main() {
	for _, structure := range Sources {
		structure.Read()
	}
}
