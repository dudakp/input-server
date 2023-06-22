package validation

type Validator[T any] interface {
	Validate(T) error
}
