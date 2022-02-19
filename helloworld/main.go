package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	version := os.Getenv("VERSION")
	if version == "" {
		version = "v1"
	}
	log.Println(version)
	fmt.Fprintf(w, "Hello world %s\n", version)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	//设置json格式
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/api/hello", handler)
	http.HandleFunc("/health", handler2)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
