package jsonx

type Optional[T any] *T

func OptionalValue[T any](v T) Optional[T] {
	return &v
}
