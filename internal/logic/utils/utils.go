package utils

func CreatePointer[T any](value T) *T {
	pointer := new(T)
	*pointer = value
	return pointer
}
