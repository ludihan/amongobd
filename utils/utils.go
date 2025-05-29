package utils

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](val T, err error) T {
	Check(err)
	return val
}
