package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		response := map[string]string{"status": "OK"}
		bytes, err := json.Marshal(&response)
		if err != nil {
			log.Fatal(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
