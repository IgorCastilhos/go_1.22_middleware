/*
This is a simple example of a mux server using the go 1.22 new http server
*/
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /path/", Logger(func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, "path capturado\n")
	}))

	mux.HandleFunc("/task/{id}/{$}", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome to the task page %v\n", request.URL.Path)
	})

	slog.Info("Server is running on localhost:7100")
	http.ListenAndServe("localhost:7100", mux)
}

func Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(writer, request)
		elapsedTime := time.Since(startTime)

		slog.Info("http request", slog.String("method", request.Method), slog.String("path", request.URL.Path), slog.String("elapsed_time", elapsedTime.String()))
	}
}
