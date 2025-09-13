package bootstrap

import (
	"log"
	"net/http"
)

func StartServer() {
	err := http.ListenAndServe(":8080", nil)
	log.Println("Server is running on port 8080")
	if err != nil {
		log.Println("Failed to start server:", err)
	}
}
