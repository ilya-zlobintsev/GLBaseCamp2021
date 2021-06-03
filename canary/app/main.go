package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var VERSION string

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index_handler)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Printf("Failed initializing web! %v", err)
	}
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App version: %v\n", VERSION)

	interfaces, _ := net.Interfaces()

	for _, iface := range interfaces {
		ips, _ := iface.Addrs()

		for _, ip := range ips {
			if !ip.(*net.IPNet).IP.IsLoopback() && !strings.Contains(ip.String(), ":") { // Show ipv4 only
				fmt.Fprintf(w, "Server IP Address: %v\n", ip)
			}
		}
	}
}
