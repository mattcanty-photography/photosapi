package main

import (
	"log"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/mattcanty-photography/photoapi/internal/photoapi"
)

func main() {
	imageSource := &photoapi.S3ImageSource{
		Region:     "eu-west-2",
		BucketName: os.Getenv("PHOTO_BUCKET_NAME"),
		Prefix:     "",
	}

	config := photoapi.Configuration{
		RoutePatternPrefix: "/{apiStage}",
		JsonLogging:        false,
		ImageSource:        imageSource,
	}

	log.Fatal(gateway.ListenAndServe(":3000", photoapi.SetupRouting(config)))
}
