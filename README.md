# cloud-storage-service

Service for exposing CRUD operations on a [Google Cloud Storage](https://cloud.google.com/storage) (GCS) bucket, with a [Gin](https://github.com/gin-gonic/gin) REST API.

## Goals

- Originally a "take-home" assignment for a job interview
- Have a template to use for something else

## Actions

1. Upload a new or existing (update) file to a GCS bucket folder
2. List all files in a GCS bucket folder
3. Get/Download the contents of a specific file in a GCS bucket folder
4. Move a file from one GCS bucket folder to another
5. Delete a file from a GCS bucket folder


## Perquisites

An active GCS account with at least one available bucket, and at least one
bucket folder (two or more bucket folders for moving files between folders).

## Running

### Configure the service

See the default `.env` file placeholders

### Build the service

```shell
go build
```

### Run the service

```shell
./cloud-storage-service
```

### Run the tests

```shell
go test -v ./...
```

### Example requests to a local development server

#### Upload a file

```shell
curl --location 'http://localhost:8080/upload' --form 'file_input=@"/path/to/your/testFile.txt"'
```

```json
{
    "message": "success"
}
```

#### Download a file

```shell
curl --location 'http://localhost:8080/download/testFile.txt'
```

```text
these are the contents of testFile
```

#### List files

```shell
curl --location 'http://localhost:8080/list'
```

```json
{
    "files": [
        "test-files/",
        "test-files/testFile-1024x576-2.jpg",
        "test-files/testFile-1024x576.jpg",
        "test-files/testFile.txt",
        "test-files/testFile1.txt",
        "test-files/testFile1.txt.zip",
        "test-files/testFile2.txt",
        "test-files/testFile3.txt",
        "test-files/testFile4.txt",
        "test-files2/"
    ]
}
```

#### Move a file

```shell
curl --location --request PUT 'http://localhost:8080/move/testFile.txt?folder=test-files2'
```

```json
{
    "message": "success"
}
```

#### Delete a file

```shell
curl --location --request DELETE 'http://localhost:8080/delete/testFile.txt'
```

```json
{
    "message": "success"
}
```
