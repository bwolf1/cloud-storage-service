package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type StorageClient struct {
	Client     *storage.Client
	ProjectID  string
	BucketName string
	Path       string
}

type Service interface {
	UploadFile(file multipart.File, object string) error
	GetFile(object string) ([]byte, error)
	DeleteFile(object string) error
	ListFiles() ([]string, error)
	MoveFile(object string, newPath string) error
}

// UploadFile uploads a bucket object
func (c *StorageClient) UploadFile(ctx context.Context, file multipart.File, object string) error {
	// Validate that the file is not nil
	if file == nil {
		return fmt.Errorf("file is nil")
	}

	// Validate that the object is not empty
	if object == "" {
		return fmt.Errorf("object is empty")
	}

	// Validate that the object is smaller than the configured max size
	configMax := os.Getenv("MAX_FILE_SIZE")
	max, err := strconv.Atoi(configMax)
	if err != nil {
		return fmt.Errorf("error converting maxSize to int: %v", err)
	}
	if len(object) > max {
		return fmt.Errorf("object is larger than the configured max size")
	}

	// Validate that the object is not a directory
	if object[len(object)-1:] == "/" {
		return fmt.Errorf("object is a directory")
	}

	maxTimeout, err := strconv.Atoi(os.Getenv("MAX_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("error converting timeout to int: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(maxTimeout))
	defer cancel()

	// Upload a bucket object
	wc := c.Client.Bucket(c.BucketName).Object(c.Path + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("error uploading file: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("error closing stream: %v", err)
	}

	return nil
}

// GetFile gets a bucket object
func (c *StorageClient) GetFile(ctx context.Context, object string) ([]byte, error) {
	maxTimeout, err := strconv.Atoi(os.Getenv("MAX_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error converting timeout to int: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(maxTimeout))
	defer cancel()

	rc, err := c.Client.Bucket(c.BucketName).Object(c.Path + object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ListFiles lists objects within a bucket
func (c *StorageClient) ListFiles(ctx context.Context) ([]string, error) {
	maxTimeout, err := strconv.Atoi(os.Getenv("MAX_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error converting timeout to int: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(maxTimeout))
	defer cancel()

	var files []string
	it := c.Client.Bucket(c.BucketName).Objects(ctx, nil)
	// TODO: Replace the default GCS bucket iterator pagination with custom pagination
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return files, err
		}
		files = append(files, attrs.Name)
	}

	return files, nil
}

// MoveFile moves a bucket object from one path to another
func (c *StorageClient) MoveFile(ctx context.Context, object string, newPath string) error {
	// Validate the path exists
	var files []string
	var err error
	// Pass 0 and "" to get all files
	if files, err = c.ListFiles(ctx); err != nil {
		return err
	}

	// Validate the path is not the same as the old path
	for _, file := range files {
		if file == newPath+object {
			return fmt.Errorf("the new path can not be the same as the old path")
		}
	}

	maxTimeout, err := strconv.Atoi(os.Getenv("MAX_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("error converting timeout to int: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(maxTimeout))
	defer cancel()

	src := c.Client.Bucket(c.BucketName).Object(c.Path + object)
	dst := c.Client.Bucket(c.BucketName).Object(newPath + object)
	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return err
	}

	if err := src.Delete(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteFile deletes a bucket object
func (c *StorageClient) DeleteFile(ctx context.Context, object string) error {
	maxTimeout, err := strconv.Atoi(os.Getenv("MAX_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("error converting timeout to int: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(maxTimeout))
	defer cancel()

	if err := c.Client.Bucket(c.BucketName).Object(c.Path + object).Delete(ctx); err != nil {
		return err
	}

	return nil
}
