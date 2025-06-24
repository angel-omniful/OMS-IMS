package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	//"strings"
	"errors"
	"path/filepath"
	ht "net/http"
	"github.com/aws/aws-sdk-go-v2/aws"
	myaws "github.com/angel-omniful/oms/client"
	"github.com/angel-omniful/oms/myContext"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	//"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	//"github.com/omniful/go_commons/csv"
)

var s3Client *s3.Client

func UploadCSVFileToS3(c *gin.Context) {
	var body map[string]interface{}

	// Parse JSON body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Extract and validate fields
	localFilePathRaw, ok := body["local_file_path"]
	if !ok {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Missing 'local_file_path'"})
		return
	}
	keyRaw, ok := body["key"]
	if !ok {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Missing 'key'"})
		return
	}
	s3Client= myaws.GetS3Client()
	localFilePath, ok := localFilePathRaw.(string)
	if !ok {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "'local_file_path' must be a string"})
		return
	}
	key, ok := keyRaw.(string)
	if !ok {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "'key' must be a string"})
		return
	}

	ctx := myContext.GetContext()
	bucket := config.GetString(ctx, "aws.s3.bucket_name")

	// Check file extension
	if filepath.Ext(localFilePath) != ".csv" {
		c.JSON(int(http.StatusBadRequest), gin.H{"error": "Only .csv files are allowed"})
		return
	}

	// Open local file
	file, err := os.Open(localFilePath)
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to open file: " + err.Error()})
		return
	}
	defer file.Close()

	// Upload to S3
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to upload to S3: " + err.Error()})
		return
	}

	// Success
	c.JSON(int(http.StatusOK), gin.H{"message": "File uploaded successfully", "key": key})
}

//input will be file_name(key is always uploadds/filename)
func GenerateCsvUrl(c *gin.Context) {
	filename := c.Param("filename") // or c.Params.ByName("filename")
	filename = "uploads/" + filename

	ctx := myContext.GetContext()
	s3Client := myaws.GetS3Client()

	presignClient := s3.NewPresignClient(s3Client)

	presignedReq, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(config.GetString(ctx, "aws.s3.bucket_name")),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Println("Not able to generate URL:", err)
		c.JSON(int(http.StatusInternalServerError), gin.H{"error": "Failed to generate presigned URL"})
		return
	}

	log.Println("Download link:", presignedReq.URL)

	
	c.JSON(int(http.StatusOK), gin.H{
		"url": presignedReq.URL,
	})
}

func GenerateCsvUrlFunc(filename string) string {
	//filename := c.Param("filename") // or c.Params.ByName("filename")
	filename = "uploads/" + filename

	ctx := myContext.GetContext()
	s3Client := myaws.GetS3Client()

	presignClient := s3.NewPresignClient(s3Client)

	presignedReq, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(config.GetString(ctx, "aws.s3.bucket_name")),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Println("Not able to generate URL:", err)
		return ""
	}

	log.Println("Download link:", presignedReq.URL)

	return presignedReq.URL
}
//input will be url string
func ValidateCSVURL(url string) error {
resp, err := ht.Get(url)
	if err != nil {
		return errors.New("failed to reach URL: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != ht.StatusOK {
		return errors.New("non-200 status code: " + resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/csv" && contentType != "application/octet-stream" {
		return errors.New("invalid content type: " + contentType)
	}

	// Optional: Try reading a few rows
	// reader := csv.NewReader(resp.Body)
	// _, err = reader.Read() // Try reading first line
	// if err != nil && err != io.EOF {
	// 	return errors.New("unable to read CSV content: " + err.Error())
	// }

	return nil // success
}

//download logic
func DownloadFileFromPresignedURL(presignedURL string) ([]byte, error) {
	resp, err := ht.Get(presignedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to GET file from presigned URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != int(http.StatusOK) {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}

	return data, nil
}


