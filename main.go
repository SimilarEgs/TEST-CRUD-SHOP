package main

import (
	"net/http"

	"github.com/SimilarEgs/main/tree/main/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
