package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func getIP(r *http.Request) string {
	// Coba ambil IP dari header X-Forwarded-For (jika pakai reverse proxy atau load balancer)
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		return xForwardedFor
	}

	// Kalau tidak ada, ambil dari RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Your IP address is: %s\n", ip)
}

func main() {
	http.HandleFunc("/", handler)

	port := "8080"
	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
