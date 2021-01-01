package main

import (
	"net/http"

	"github.com/yamad07/go-modular-monolith/pkg/config"
)

func main() {
	config.Init()
	r := NewRouter()
	http.ListenAndServe(":9000", r)
}
