package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	Router        *mux.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := mux.NewRouter()

	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddRoute(method string, path string, handler http.HandlerFunc) {
	switch method {
	case "GET":
		s.Router.HandleFunc(path, handler).Methods("GET")
	case "POST":
		s.Router.HandleFunc(path, handler).Methods("POST")
	case "PUT":
		s.Router.HandleFunc(path, handler).Methods("PUT")
	case "DELETE":
		s.Router.HandleFunc(path, handler).Methods("DELETE")
	default:
		log.Printf("Unsupported method: %s\n", method)
	}
	fmt.Printf("Added Route: %s %s\n", method, path)
}

func (s *WebServer) Start() {
	if err := http.ListenAndServe(s.WebServerPort, s.Router); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
