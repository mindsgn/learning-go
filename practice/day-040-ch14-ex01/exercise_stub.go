package main

import "net/http"

func Timeout(ms int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler { return h }
}
