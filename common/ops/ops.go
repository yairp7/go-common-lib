package ops

func T[R any](cond bool, a, b R) R {
	if cond {
		return a
	}
	return b
}
