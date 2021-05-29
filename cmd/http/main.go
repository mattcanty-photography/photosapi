package main

import (
	"log"
	"net/http"

	"github.com/mattcanty-photography/photoapi/internal/photoapi"
)

func main() {
	// imageSource := &photoapi.FileSystemImageSource{
	// 	Directory: "~/Photos",
	// }

	imageSource := &photoapi.S3ImageSource{
		Region:     "eu-west-2",
		BucketName: "photo-mattcanty-com-dev-photos-4681677",
		Prefix:     "",
	}

	config := photoapi.Configuration{
		RoutePatternPrefix: "",
		JsonLogging:        false,
		ImageSource:        imageSource,
	}

	log.Println("Starting Photo API")
	log.Println(http.ListenAndServe(":3001", photoapi.SetupRouting(config)))
}
