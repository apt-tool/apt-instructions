package storage

import (
	"context"
	"fmt"
	"time"

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

	ctx := context.Background()
	if flag, err := conn.BucketExists(ctx, cfg.Bucket); err != nil {
		return nil, fmt.Errorf("failed to get bucket status=%w", err)
	} else if !flag {
		if er := conn.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{}); er != nil {
			return nil, fmt.Errorf("failed to create bucket=%w", er)
		}
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
	ctx := context.Background()

	if _, err := c.connection.FPutObject(ctx, c.config.Bucket, name, path, minio.PutObjectOptions{
		ContentType: "application/text",
	}); err != nil {
		return fmt.Errorf("[minio.Client.Put] failed to put object error=%w", err)
	}

	return nil
}

func (c client) Get(name string) (string, error) {
	ctx := context.Background()

	url, err := c.connection.PresignedGetObject(ctx, c.config.Bucket, name, 2*time.Hour, nil)
	if err != nil {
		return "", fmt.Errorf("[minio.Client.Get] failed to get url error=%w", err)
	}

	return url.String(), nil
}
