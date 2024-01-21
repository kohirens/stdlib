package web

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/kohirens/stdlib/path"
	"os"
)

func loadEvent(s string) events.LambdaFunctionURLRequest {
	if !path.Exist(s) {
		panic(fmt.Sprintf("file %s not found", s))
	}

	content, e1 := os.ReadFile(s)
	if e1 != nil {
		panic(fmt.Sprintf("could not read file %s", s))
	}

	var req events.LambdaFunctionURLRequest

	if e := json.Unmarshal(content, &req); e != nil {
		panic(fmt.Sprintf("could not decode JSON file %v: %v", s, e.Error()))
	}

	return req
}
