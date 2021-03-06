package body1

import (
	"context"
	"time"

	"github.com/knowledgebao/gotest/tags/src/imply/body"
)

var _ body.Body = &body1Imp{}

type body1Imp struct {
	result chan *body.Result
}

func CreateBody(ctx context.Context, c body.Config) (body.Body, error) {
	b := body1Imp{}
	b.result = make(chan *body.Result, 100)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(b.result)
				return
			default:
				b.result <- &body.Result{Desc: c.Type + c.IP}
				time.Sleep(time.Second)
			}
		}
	}()
	return &b, nil
}

func (b *body1Imp) Output() chan *body.Result {
	return b.result
}

func (b *body1Imp) Close() error {
	return nil
}
