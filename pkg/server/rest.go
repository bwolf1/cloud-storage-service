package server

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/bwolf1/cloud-storage-service/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var storageClient *service.StorageClient

func New() {
	// Load the config from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	storageClient = &service.StorageClient{
		Client:     client,
		BucketName: os.Getenv("BUCKET_NAME"),
		ProjectID:  os.Getenv("PROJECT_ID"),
		Path:       os.Getenv("OBJECT_PATH"),
	}

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		f, err := c.FormFile(os.Getenv("FILE_INPUT"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"file key configuration error": err.Error(),
			})
			return
		}

		blobFile, err := f.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error opening file": err.Error(),
			})
			return
		}

		err = storageClient.UploadFile(c, blobFile, f.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error uploading file": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	router.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		data, err := storageClient.GetFile(c, filename)
		message := "error retrieving file"
		if err != nil {
			if err == storage.ErrObjectNotExist {
				c.JSON(http.StatusNotFound, gin.H{
					message: err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				message: err.Error(),
			})
			return
		}

		c.Data(http.StatusOK, "application/octet-stream", data)
	})

	router.GET("/list", func(c *gin.Context) {
		files, err := storageClient.ListFiles(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error listing files": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"files": files,
		})
	})

	router.PUT("/move/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		newPath := c.Query("folder")
		err := storageClient.MoveFile(c, filename, newPath)
		message := "error moving file"
		if err != nil {
			if err == storage.ErrObjectNotExist {
				c.JSON(http.StatusNotFound, gin.H{
					message: err.Error(),
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				message: err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	router.DELETE("/delete/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		err := storageClient.DeleteFile(c, filename)
		message := "error deleting file"
		if err != nil {
			if err == storage.ErrObjectNotExist {
				c.JSON(http.StatusNotFound, gin.H{
					message: err.Error(),
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				message: err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run()
}
