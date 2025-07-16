package actor

import (
	"fmt"
	"github.com/vladopajic/go-actor/actor"
	"log"
	"time"
)

const (
	msgProduceTime = time.Second * 5
)

type processor struct {
	logger  *log.Logger
	mailbox actor.MailboxSender[string]
}

func (w *processor) DoWork(ctx actor.Context) actor.WorkerStatus {
	select {
	case <-ctx.Done():
		return actor.WorkerEnd
	case <-time.After(msgProduceTime):
		if err := w.mailbox.Send(ctx, time.Now().String()); err != nil {
			w.logger.Println(fmt.Sprintf("Error sending mail: %s", err))
		} else {
			w.logger.Println(fmt.Sprintf("Sent msg to poller"))
		}
		return actor.WorkerContinue
	}
}
