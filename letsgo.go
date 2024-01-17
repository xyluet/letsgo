package letsgo

// Must panics if err is not nil. Otherwise, it returns v.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
