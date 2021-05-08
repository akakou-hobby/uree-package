package uree_package

type Request struct {
	Path     string
	Body     string
	Optional string
}

type Response struct {
	Body string
}

type UreePackage interface {
	Run(Request) Response
}
