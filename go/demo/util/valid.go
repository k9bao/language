package util

import (
	"errors"
	"regexp"
)

func ValidName(name string) error {
	if matched, _ := regexp.Match("^[0-9a-zA-Z_]+$", []byte(name)); !matched {
		Log.Errorf("name is illegal: %s", name)
		return errors.New("name is illegal")
	}
	return nil
}
