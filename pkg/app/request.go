package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gookit/goutil/dump"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		dump.P(err.Key, err.Message)
	}

	return
}
