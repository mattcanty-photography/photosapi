package main

import (
	"log"
	"net/http"

	"github.com/mattcanty-photography/photoapi/internal/photoapi"
)

func main() {
	imageSource := &photoapi.FileSystemImageSource{
		Directory: "~/Photos",
	}

	config := photoapi.Configuration{
		RoutePatternPrefix: "",
		JsonLogging:        false,
		ImageSource:        imageSource,
	}

	log.Println("Starting Photo API")
	log.Println(http.ListenAndServe(":3001", photoapi.SetupRouting(config)))
}
