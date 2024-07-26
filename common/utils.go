package common

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T Number](num T) T {
	if num < 0 {
		return -num
	}
	return num
}
