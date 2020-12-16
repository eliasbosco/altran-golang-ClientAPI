package main

import (
	"log"
	"net/http"

	"./ports"
)

func main() {
	http.HandleFunc("/ports", ports.GetPorts)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
