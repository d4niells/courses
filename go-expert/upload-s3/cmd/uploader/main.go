package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var (
	client *s3.S3
	bucket string
	tmpDir string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_USER_ACCESS_KEY"),
				os.Getenv("AWS_USER_SECRET_ACCESS_KEY"),
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	tmpDir = "./tmp"
	client = s3.New(session)
	bucket = "goexpert-bucket"
}

func main() {
	dir, err := os.Open(tmpDir)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	for {
		files, err := dir.ReadDir(1)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading directory: %s\n", err)
			return
		}

		for _, f := range files {
			uploadFile(f.Name())
		}
	}
}

func uploadFile(filename string) {
	path := fmt.Sprintf("%s/%s", tmpDir, filename)
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", filename)
	}
	defer f.Close()

	_, err = client.PutObject(&s3.PutObjectInput{Bucket: aws.String(bucket), Key: aws.String(filename), Body: f})
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", filename)
	}

	fmt.Printf("File %s was uploaded successfully\n", filename)
}
