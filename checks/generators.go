package checks

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

/*
	Even I cannot understand all of this.
	If you do, please explain it to me.
	This abstracted-to-hell abomination
	is the result of me trying to play
	God with the generic methods. May
	God have mercy upon my soul.
*/

func ContainsInRange[T comparable](element T, min_length, max_length int) CheckFunc[T] {
	return func(t []T) error {
		count := Count(element, t)
		if count < min_length {
			return fmt.Errorf("must contain at least %d of %v", min_length, element)
		}
		if count < max_length {
			return fmt.Errorf("must contain at most %d of %v", max_length, element)
		}
		return nil
	}
}

func ContainsExact[T comparable](element T, lenght int) CheckFunc[T] {
	return ContainsInRange(element, lenght, lenght)
}

func ContainsAtLeast[T comparable](element T, length int) CheckFunc[T] {
	return ContainsInRange(element, length, math.MaxInt)
}

func ContainsAtMost[T comparable](element T, length int) CheckFunc[T] {
	return ContainsInRange(element, 0, length)
}

func ContainsFromInRange[T comparable, S []T](s S, min_count, max_count int) CheckFunc[T] {
	return func(t []T) error {
		count := 0
		for _, element := range t {
			if slices.Contains(s, element) {
				count++
			}
		}

		if count < min_count {
			return fmt.Errorf("property must contain at least %d of certain set", min_count)
		}

		if count > max_count {
			return fmt.Errorf("property must contain at most %d of certain set", max_count)
		}

		return nil
	}
}

func ContainsFromExact[T comparable, S []T](s S, count int) CheckFunc[T] {
	return ContainsFromInRange(s, count, count)
}

func ContainsFromAtLeast[T comparable, S []T](s S, min_count int) CheckFunc[T] {
	return ContainsFromInRange(s, min_count, math.MaxInt)
}

func ContainsFromAtMost[T comparable, S []T](s S, max_count int) CheckFunc[T] {
	return ContainsFromInRange(s, 0, max_count)
}

func LengthInRange[T any](min_length, max_length int) CheckFunc[T] {
	return func(t []T) error {
		if len(t) < min_length {
			return fmt.Errorf("must be of length of at least %d", min_length)
		}
		if len(t) > max_length {
			return fmt.Errorf("must be of length of at most %d", max_length)
		}
		return nil
	}
}

func LengthExact[T any](length int) CheckFunc[T] {
	return LengthInRange[T](length, length)
}

func LengthAtLeast[T any](length int) CheckFunc[T] {
	return LengthInRange[T](length, math.MaxInt)
}

func LengthAtMost[T any](length int) CheckFunc[T] {
	return LengthInRange[T](0, length)
}

func ContainsOnlyFrom[T comparable, S []T](s S) CheckFunc[T] {
	return func(t []T) error {
		for _, element := range t {
			if !slices.Contains(s, element) {
				return errors.New("unallowed element detected")
			}
		}
		return nil
	}
}
