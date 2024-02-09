package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to /time endpoint")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, time.Now().String())
	})

	slog.Info("Starting server 2 on port 8891")

	err := http.ListenAndServe(":8891", nil)
	if err != nil {
		slog.Error("Application 2 finished with error", "error", err)
	}
}
