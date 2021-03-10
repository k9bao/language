package driver

import (
	"context"
	"fmt"

	"github.com/knowledgebao/gotest/tags/src/imply/body"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/syncmap"
)

var (
	log     = logrus.StandardLogger()
	drivers = syncmap.Map{}
)

type Driver interface {
	Create(ctx context.Context, c body.Config) (body.Body, error)
	Init() error
	Release() error
}

func Register(name string, driver Driver) error {
	_, existed := drivers.LoadOrStore(name, driver)
	if existed {
		return fmt.Errorf("caputre driver %s existed!", name)
	}
	return nil
}

func Init() error {
	drivers.Range(func(key, value interface{}) bool {
		err := value.(Driver).Init()
		if err != nil {
			println("init driver fail", key)
		} else {
			println("init driver OK", key)
		}
		return true
	})
	return nil
}

func Release() error {
	drivers.Range(func(key, value interface{}) bool {
		err := value.(Driver).Release()
		if err != nil {
			log.Error("release driver", key, err.Error())
		} else {
			log.Info("release driver", key, "OK")
		}
		return true
	})
	return nil
}

func List() []string {
	var res []string
	drivers.Range(func(key, value interface{}) bool {
		res = append(res, key.(string))
		return true
	})
	return res
}

func Create(ctx context.Context, config body.Config) (body.Body, error) {
	driver, existed := drivers.Load(config.Type)
	if !existed {
		log.Errorf("driver %s Not found", config.Type)
		return nil, fmt.Errorf("driver %s Not found", config.Type)
	}
	return driver.(Driver).Create(ctx, config)
}
