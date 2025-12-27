package util

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
