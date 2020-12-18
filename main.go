package main

import (
	"./ports"
	"./types"
	"log"
	"net/http"
)

func main() {
	// Set up all configurations variables
	config := types.SetupConfig()
	log.Printf("main.SetupConfig: %#v", config)

	log.Printf("main.http.HandleFunc - endpoint(/ports)")
	http.HandleFunc("/ports", ports.GetPorts)

	log.Printf("main.http.HandleFunc - endpoint(/ports-db)")
	http.HandleFunc("/ports-db", ports.GetPortsDb)

	log.Fatal(http.ListenAndServe(config.APIAddress, nil))
}
