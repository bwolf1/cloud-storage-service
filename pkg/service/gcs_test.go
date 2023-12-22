package service

import (
	"context"
	"mime/multipart"
	"reflect"
	"testing"

	"cloud.google.com/go/storage"
)

func TestStorageClient_UploadFile(t *testing.T) {
	type fields struct {
		Client     *storage.Client
		ProjectID  string
		BucketName string
		UploadPath string
	}
	type args struct {
		ctx    context.Context
		file   multipart.File
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &StorageClient{
				Client:     tt.fields.Client,
				ProjectID:  tt.fields.ProjectID,
				BucketName: tt.fields.BucketName,
				UploadPath: tt.fields.UploadPath,
			}
			if err := c.UploadFile(tt.args.ctx, tt.args.file, tt.args.object); (err != nil) != tt.wantErr {
				t.Errorf("StorageClient.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageClient_GetFile(t *testing.T) {
	type fields struct {
		Client     *storage.Client
		ProjectID  string
		BucketName string
		UploadPath string
	}
	type args struct {
		ctx    context.Context
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &StorageClient{
				Client:     tt.fields.Client,
				ProjectID:  tt.fields.ProjectID,
				BucketName: tt.fields.BucketName,
				UploadPath: tt.fields.UploadPath,
			}
			got, err := c.GetFile(tt.args.ctx, tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageClient.GetFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageClient.GetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageClient_ListFiles(t *testing.T) {
	type fields struct {
		Client     *storage.Client
		ProjectID  string
		BucketName string
		UploadPath string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &StorageClient{
				Client:     tt.fields.Client,
				ProjectID:  tt.fields.ProjectID,
				BucketName: tt.fields.BucketName,
				UploadPath: tt.fields.UploadPath,
			}
			got, err := c.ListFiles(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageClient.ListFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StorageClient.ListFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageClient_MoveFile(t *testing.T) {
	type fields struct {
		Client     *storage.Client
		ProjectID  string
		BucketName string
		UploadPath string
	}
	type args struct {
		ctx     context.Context
		object  string
		newPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &StorageClient{
				Client:     tt.fields.Client,
				ProjectID:  tt.fields.ProjectID,
				BucketName: tt.fields.BucketName,
				UploadPath: tt.fields.UploadPath,
			}
			if err := c.MoveFile(tt.args.ctx, tt.args.object, tt.args.newPath); (err != nil) != tt.wantErr {
				t.Errorf("StorageClient.MoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageClient_DeleteFile(t *testing.T) {
	type fields struct {
		Client     *storage.Client
		ProjectID  string
		BucketName string
		UploadPath string
	}
	type args struct {
		ctx    context.Context
		object string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &StorageClient{
				Client:     tt.fields.Client,
				ProjectID:  tt.fields.ProjectID,
				BucketName: tt.fields.BucketName,
				UploadPath: tt.fields.UploadPath,
			}
			if err := c.DeleteFile(tt.args.ctx, tt.args.object); (err != nil) != tt.wantErr {
				t.Errorf("StorageClient.DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
