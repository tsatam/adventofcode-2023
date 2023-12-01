package fp

func Map[T, U any](in []T, mapper func(t T) U) []U {
	result := make([]U, len(in))
	for i, v := range in {
		result[i] = mapper(v)
	}
	return result
}
