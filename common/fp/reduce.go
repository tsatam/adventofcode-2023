package fp

func Reduce[T, U any](in []T, identity U, combine func(curr U, next T) U) U {
	result := identity
	for _, next := range in {
		result = combine(result, next)
	}
	return result
}
