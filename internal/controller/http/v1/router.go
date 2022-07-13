package v1

import (
	"net/http"
	"reegle/config"
	"reegle/internal/router"
	"reegle/internal/usecase/search/delivery"
	"reegle/pkg/dict"
	"reegle/pkg/parser"
)

func NewRouter(db *dict.WordDB, config *config.Config) *http.ServeMux {
	handler := router.NewRouter()
	handler.Static("static")

	handler.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(
			parser.ParseTemplate("templates/home.template.html", nil),
		)
	})
	delivery.NewSearchDelivery(handler, db)

	return handler.Mux()
}
