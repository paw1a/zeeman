package app

import (
	"github.com/paw1a/interferometer/backend/internal/handler"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/api/interferometer", handler.GetImage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
