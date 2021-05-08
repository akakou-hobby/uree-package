package uree_package

type Request struct {
	path     string
	body     string
	optional string
}

type Response struct {
	body string
}

type UreePackage interface {
	Run(Request) Response
}
