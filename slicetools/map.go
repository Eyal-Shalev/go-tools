package slicetools

func Map[S ~[]From, From, To any](s S, f func(From) To) []To {
	out := make([]To, len(s))
	for i, val := range s {
		out[i] = f(val)
	}
	return out
}
