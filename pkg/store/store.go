package store

// Store is a wrapper for communication with storage.
type Store interface {
	Manufacturer() ManufacturerRepository
}
