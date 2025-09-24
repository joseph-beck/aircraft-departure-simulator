package preparer

// Preparer defines a generic interface for preparing data of type T into type V.
type Preparer[T, V any] interface {
	// Prepare transforms data of type T into type V.
	// It returns an error if the preparation fails.
	Prepare(T) (V, error)
}
