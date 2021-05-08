package uree_package

import (
	"fmt"
	"testing"
)

type MyPackage struct{}

func (myPackage MyPackage) Run(req Request) Response {
	fmt.Println(req.Path)
	fmt.Println(req.Body)
	fmt.Println(req.Optional)

	resp := Response{
		req.Body,
	}

	return resp
}

func (myPackage MyPackage) SetUpOptional() string {
	return ""
}

func TestExampleSuccess(t *testing.T) {
	req := Request{
		"path",
		"body",
		"optional",
	}

	myPackage := MyPackage{}

	myPackage.Run(req)
}
