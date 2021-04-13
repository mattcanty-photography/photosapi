package main

import (
	"log"
	"net/http"

	"github.com/mattcanty-photography/photoapi/internal/photoapi"
)

func main() {
	imageSource := &photoapi.FileSystemImageSource{
		Directory: "/Users/mattcanty/Documents/20-29 Photography/21 BPCG/21.01 Weekly Photos/2021-04-02",
	}

	// imagesource := &S3ImageSource{
	// 	Region:     "eu-west-2",
	// 	BucketName: "photo-mattcanty-com-dev-photos-4681677",
	// 	Prefix:     "portfolios/bpcg/albums/2021-02-12/fulls",
	// }

	config := photoapi.Configuration{
		RoutePatternPrefix: "",
		JsonLogging:        false,
		ImageSource:        imageSource,
	}

	log.Println("Starting Photo API")
	log.Println(http.ListenAndServe(":3001", photoapi.SetupRouting(config)))
}
