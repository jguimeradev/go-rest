package main

import (
	"log"
	"net/http"

	"github.com/jguimeradev/go-rest/internals/handler"
	"github.com/jguimeradev/go-rest/internals/store"
)

func main() {
	s := store.New()
	h := handler.New(s)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":6969", mux))
}
