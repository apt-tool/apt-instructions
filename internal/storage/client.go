package storage

type Client interface {
	Put(name, path string) error
	Get(name string) (string, error)
}
