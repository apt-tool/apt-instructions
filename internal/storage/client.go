package storage

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client interface {
	Put(name, path string) error
	Get(name string) (string, error)
}

func New(cfg Config) (Client, error) {
	conn, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Access, cfg.Secret, ""),
		Secure: cfg.SSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Minio error=%w", err)
	}

	return client{
		config:     cfg,
		connection: conn,
	}, nil
}

type client struct {
	config     Config
	connection *minio.Client
}

func (c client) Put(name, path string) error {

}

func (c client) Get(name string) (string, error) {

}
