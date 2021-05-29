package photoapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
)

func SetupRouting(config Configuration) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON: config.JsonLogging,
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(9, "text/html"))

	r.HandleFunc(fmt.Sprintf("%s/favicon.ico", config.RoutePatternPrefix), render.NoContent)
	r.HandleFunc(fmt.Sprintf("%s/photo/*", config.RoutePatternPrefix), config.thumbnailHandler)

	r.NotFound(notFoundPageHandler)

	return r
}
