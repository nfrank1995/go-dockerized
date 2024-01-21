package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"response": "Dockerizing Go Application",
		}

		err := json.NewEncoder(rw).Encode(response)
		if err != nil {
			log.Printf("Error encoding response: %s", err)
		}
	})

	log.Println("⚡️[bootup]: Server is running at port: 5000")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Printf("Error running the server: %s", err)
	}
}
