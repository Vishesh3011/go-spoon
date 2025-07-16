package actor

import (
	"fmt"
	"github.com/vladopajic/go-actor/actor"
	"log"
	"time"
)

type poller struct {
	logger  *log.Logger
	mailbox actor.MailboxReceiver[string]
}

func (w *poller) DoWork(ctx actor.Context) actor.WorkerStatus {
	select {
	case <-ctx.Done():
		return actor.WorkerEnd
	case msg := <-w.mailbox.ReceiveC():
		w.logger.Println(fmt.Sprintf("Received mailbox: %s", msg))
		return actor.WorkerContinue
	case <-time.After(10 * time.Second):
		w.logger.Println("Polling timeout.")
		return actor.WorkerContinue
	}
}
