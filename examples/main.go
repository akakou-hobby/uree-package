package main

import (
	"fmt"

	uree_package "github.com/akakou-hobby/uree-package"
)

type MyPackage struct{}

func (myPackage MyPackage) Run(req uree_package.Request) uree_package.Response {
	fmt.Println(req.Path)
	fmt.Println(req.Body)
	fmt.Println(req.Optional)

	resp := uree_package.Response{
		Body: "body",
	}

	return resp
}

func main() {
	req := uree_package.Request{
		Path:     "path",
		Body:     "body",
		Optional: "optional",
	}

	myPackage := MyPackage{}

	myPackage.Run(req)
}
