package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()

	h2Handler := h2c.NewHandler(mux, &http2.Server{})
	server := &http.Server{Addr: ":3999", Handler: h2Handler}
	server.ListenAndServe()
}
