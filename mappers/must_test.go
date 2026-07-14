package mappers_test

// Must unwraps values that would normally return T, error but here returns T, strings2.Stats
func Must[T any, E any](v T, err E) T {
	return v
}
