package checks

// I wanted to explain this too, but if you
// cannot get the gist of it from the name
// alone, then I have no hope for you.
// Jokes aside, runs a bunch of CheckFuncs
// on a single input and throws the first
// error it sees (if any) at you.
func RunChecklist[T any, S []T](element S, checks []CheckFunc[T]) (int, error) {
	for i, check := range checks {
		err := check(element)
		if err != nil {
			return i, err
		}
	}

	return -1, nil
}
