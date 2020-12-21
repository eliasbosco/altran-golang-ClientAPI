// Altran Golang ClientAPI
// Following the requirements asked, this service is
// responsible to fetch all the file ports (ports.json)
// objects by the request of the endpoint "/ports",
// which is provided with the mandatories query
// parameters: skip and limit. Every time when each
// request is made on this endpoint, the resulted records
// are compiled and sent to service PortDomainService via
// gRPC protocol with protobuf to envelope the data, which
// insert or update every record on SQLite database attached.
//
// There's another endpoint (/ports-db) which is responsible
// to fetch all the database records and provide it as response
// json. There are some query parameters: skip and limit, in
// case to slice part of the result and port_id, to return
// only the port related to it's ID.
package main

import (
	"./ports"
	"./types"
	"log"
	"net/http"
)

// Start setting up the configuration and put all the
// endpoint up
func main() {
	// Set up all configurations variables
	config := types.SetupConfig()
	log.Printf("main.SetupConfig: %#v", config)

	log.Printf("main.http.HandleFunc(ports.GetPorts) - endpoint(/ports)")
	http.HandleFunc("/ports", ports.GetPorts)

	log.Printf("main.http.HandleFunc(ports.GetPortsDb) - endpoint(/ports-db)")
	http.HandleFunc("/ports-db", ports.GetPortsDb)

	log.Fatal(http.ListenAndServe(config.APIAddress, nil))
}
