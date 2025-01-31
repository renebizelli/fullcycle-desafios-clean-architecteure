package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	h := make(map[string]map[string]http.HandlerFunc)
	h["GET"] = make(map[string]http.HandlerFunc)
	h["POST"] = make(map[string]http.HandlerFunc)

	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      h,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddGetHandler(path string, handler http.HandlerFunc) {
	s.addHandler("GET", path, handler)
}

func (s *WebServer) AddPostHandler(path string, handler http.HandlerFunc) {
	s.addHandler("POST", path, handler)
}

func (s *WebServer) addHandler(method string, path string, handler http.HandlerFunc) {
	fmt.Println("Adding handler", method, path)
	s.Handlers[method][path] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {

	s.Router.Use(middleware.Logger)

	for method, _map := range s.Handlers {
		for path, handler := range _map {
			s.Router.Method(method, path, handler)
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
