package actor

import (
	"fmt"
	"github.com/vladopajic/go-actor/actor"
	"log"
	"time"
)

type Worker struct {
	Logger *log.Logger
}

func (w *Worker) Run() {
	mailbox := actor.NewMailbox[string]()

	pro := &processor{
		logger:  w.Logger,
		mailbox: mailbox,
	}

	poll := &poller{
		logger:  w.Logger,
		mailbox: mailbox,
	}

	act := actor.Combine(mailbox, actor.New(pro), actor.New(poll)).Build()
	act.Start()
	defer act.Stop()

	w.Logger.Println(fmt.Sprintf("Worker started at %v", time.Now()))

	select {}
}
