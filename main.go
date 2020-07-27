package main

import (
	"net/http"

	"github.com/bryan01-Martin/goWeb/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
