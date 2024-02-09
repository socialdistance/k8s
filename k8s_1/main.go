package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Recieved requirest to /hello endpoint")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello world")
	})

	slog.Info("Starting server on port 8890")

	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		slog.Error("Application finished with error", "error", err)
	}
}
