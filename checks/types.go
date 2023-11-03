package checks

// Shit goes in, shit gets processed, shit gets either validated or rejected
type CheckFunc[T any] func([]T) error
