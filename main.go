package main

import (
	"log"
	"magnet_api/api"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	EnvironmentVar()

	// Init the mux router
	r := mux.NewRouter()

	// Route Handlers
	r.HandleFunc("/v1/magnet/records", api.GetRecords).Methods("GET")
	r.HandleFunc("/v1/magnet/records/{id}", api.GetRecord).Methods("GET")
	r.HandleFunc("/v1/magnet/records", api.CreateRecord).Methods("POST")
	r.HandleFunc("/v1/magnet/records/{id}", api.DeleteRecord).Methods("DELTE")
	r.HandleFunc("/v1/magnet/records/{id}", api.UpdateRecord).Methods("PUT")

	// Start Server
	log.Fatal(http.ListenAndServe(":8080", r))

}
