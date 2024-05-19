/*
This is a simple example of a mux server using the go 1.22 new http server
*/
package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /home/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome to the home page %v\n", request.URL.Path)
	})

	mux.HandleFunc("GET /task/{id}/{$}", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome to the task page %v\n", request.URL.Path)
	})

	slog.Info("Server is running on localhost:7100")
	http.ListenAndServe("localhost:7100", mux)
}
