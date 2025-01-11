package test

import (
	"fmt"
	"testing"
)

type Response map[string]interface{}

func TestNewTestSuite(t *testing.T) {
	response := Response{}
	response["status"]=123
	fmt.Println(response)
}
