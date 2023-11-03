package checks

// How much of t is found in s.
// NOTE: hey, Golang, add this fucking function
// to golang.org/x/exp/slices. I'm sick of
// writing it again and again.
func Count[T comparable, S []T](t T, s S) int {
	count := 0

	for _, element := range s {
		if element == t {
			count++
		}
	}

	return count
}
