package uree_package

type Request struct {
	Path     string
	Body     string
	Optional string
}

type Response struct {
	Body string
}

type UreeNavberPackage interface {
	Run(Request) Response
	SetUpOptional() string
	GetName() string
}

type UreeBodyPackage interface {
	Run(Request) Response
	SetUpOptional() string
	GetName() string
}

type UreeLeftPackage interface {
	Run(Request) Response
	SetUpOptional() string
	GetName() string
	GetIconPath() string
}
