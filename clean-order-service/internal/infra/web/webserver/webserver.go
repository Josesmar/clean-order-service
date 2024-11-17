package webserver

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux" // Importando o pacote mux
)

type WebServer struct {
	Router        *mux.Router // Usando *mux.Router ao invés de *chi.Mux
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := mux.NewRouter() // Criando o roteador do mux

	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddRoute(method string, path string, handler http.HandlerFunc) {
	switch method {
	case "GET":
		s.Router.HandleFunc(path, handler).Methods("GET") // Definindo a rota GET
	case "POST":
		s.Router.HandleFunc(path, handler).Methods("POST") // Definindo a rota POST
	case "PUT":
		s.Router.HandleFunc(path, handler).Methods("PUT") // Definindo a rota PUT
	case "DELETE":
		s.Router.HandleFunc(path, handler).Methods("DELETE") // Definindo a rota DELETE
	default:
		log.Printf("Unsupported method: %s\n", method)
	}
	fmt.Printf("Added Route: %s %s\n", method, path)
}

func (s *WebServer) Start() {
	fmt.Printf("Starting web server on port %s\n", s.WebServerPort)
	if err := http.ListenAndServe(s.WebServerPort, s.Router); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
