package sliceutil

func Contains[T int | int64 | string](t1 T, s []T) bool {
	for _, v := range s {
		if v == t1 {
			return true
		}
	}
	return false
}
