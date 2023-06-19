package validation

type Validator interface {
	Validate() error
}

type RequestValidator interface {
	Validator
	ToResponse() error
}
