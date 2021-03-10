//+build body2

package driver

import (
	"context"
	"fmt"

	"github.com/knowledgebao/gotest/tags/src/imply/body"
	"github.com/knowledgebao/gotest/tags/src/imply/body/body2"
)

var _ Driver = &driver_body2{}

type driver_body2 struct {
}

func (d *driver_body2) Create(ctx context.Context, c body.Config) (body.Body, error) {
	return body2.CreateBody(ctx, c)
}

func (d *driver_body2) Release() error {
	return nil
}

func (d *driver_body2) Init() error {
	return nil
}

func init() {
	Register("body2", &driver_body2{})
	fmt.Println("body2 drivate init")
}
