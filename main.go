package main

import (
    "log"
    "net/http"

	"Tugas-kedua-api/database"
	"Tugas-kedua-api/handlers"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	// Read
	router.HandleFunc("/orders/{orderId}", handlers.GetOrder).Methods("GET")
	// Read-all
	router.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	// Update
	router.HandleFunc("/orders/{orderId}", handlers.UpdateOrder).Methods("PUT")
	// Delete
	router.HandleFunc("/orders/{orderId}", handlers.DeleteOrder).Methods("DELETE")
	// Initialize db connection
	database.InitDB()

    log.Fatal(http.ListenAndServe(":8081", router))
}