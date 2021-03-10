//+build body1

package driver

import (
	"context"
	"fmt"

	"github.com/knowledgebao/gotest/tags/src/imply/body"
	"github.com/knowledgebao/gotest/tags/src/imply/body/body1"
)

var _ Driver = &driver_body1{}

type driver_body1 struct {
}

func (d *driver_body1) Create(ctx context.Context, c body.Config) (body.Body, error) {
	return body1.CreateBody(ctx, c)
}

func (d *driver_body1) Release() error {
	return nil
}

func (d *driver_body1) Init() error {
	return nil
}

func init() {
	Register("body1", &driver_body1{})
	fmt.Println("body1 drivate init")
}
