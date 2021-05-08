package uree_package

import (
	"fmt"
	"testing"
)

type MyPackage struct{}

func (myPackage MyPackage) Run(req Request) Response {
	fmt.Println(req.path)
	fmt.Println(req.body)
	fmt.Println(req.optional)

	resp := Response{
		req.body,
	}

	return resp
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
