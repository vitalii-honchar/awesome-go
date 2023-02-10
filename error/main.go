package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

func main() {
	err := errors.New("test error")
	e := wrapError(err, "test message: '%s'", "test argument")
	fmt.Println(e)
}
