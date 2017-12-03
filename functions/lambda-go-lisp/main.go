package main

import (
	"encoding/json"
	apex "github.com/apex/go-apex"
	"github.com/otakumesi/lispon/lisp"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		runResult := lisp.Interpreter("./hello.lpn")

		return runResult[len(runResult)-1], nil
	})
}
