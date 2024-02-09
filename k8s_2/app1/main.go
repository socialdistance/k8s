package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	app2URL, ok := os.LookupEnv("APP2_URL")
	if !ok {
		slog.Error("missing APP2_URL")
		os.Exit(1)
	}

	getTimeURL := fmt.Sprintf("%s/time", app2URL)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Recieved requirest to /hello endpoint")

		app2Resp, err := http.DefaultClient.Get(getTimeURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "App2 (%s) is not available: %v", getTimeURL, err)
			return
		}
		defer app2Resp.Body.Close()

		respTime, _ := io.ReadAll(app2Resp.Body)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world at %s", string(respTime))
	})

	slog.Info("Starting server on port 8890")

	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		slog.Error("Application finished with error", "error", err)
	}
}
