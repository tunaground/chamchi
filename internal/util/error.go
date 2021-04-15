package util

import (
	"strings"
)

func StackError(errors ...error) string {
	var es []string
	for _, e := range errors {
		es = append(es, e.Error())
	}
	return strings.Join(es, ": ")
}
