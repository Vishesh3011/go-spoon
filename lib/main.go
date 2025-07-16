package main

import (
	"example/lib/actor"
	"log"
	"os"
)

func main() {
	worker := &actor.Worker{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
	worker.Run()
}
