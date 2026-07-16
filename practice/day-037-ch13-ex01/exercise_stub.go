package main

import "net/http"

func createServeMux() *http.ServeMux {
	return http.NewServeMux()
}
