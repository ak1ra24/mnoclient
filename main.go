package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v6"
)

func main() {
	endpoint := os.Getenv("MINIO_SERVER")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	useSSL := true

	var (
		bucket       string
		file         string
		downloadfile string
	)
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.StringVar(&file, "file", "", "file name")
	flag.StringVar(&downloadfile, "download", "", "download file path")
	flag.Parse()

	if bucket != "" && file != "" && downloadfile != "" {
		// Initialize minio client object.
		minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
		if err != nil {
			log.Fatalln(err)
		}

		if err := minioClient.FGetObject(bucket, file, downloadfile, minio.GetObjectOptions{}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Download: ", downloadfile)
		}
	} else {
		fmt.Println("Set -bucket, -file, -download")
		os.Exit(1)
	}
}
