package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerInfo struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []HandlerInfo
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []HandlerInfo{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	validMethods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}
	isValidMethod := false

	for _, m := range validMethods {
		if method == m {
			s.Handlers = append(s.Handlers, HandlerInfo{
				Method:  method,
				Path:    path,
				Handler: handler,
			})
			isValidMethod = true
			break
		}
	}
	if !isValidMethod {
		panic(fmt.Sprintf("unsupported method on handler: %s", method))
	}

}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)

	for _, handlerInfo := range s.Handlers {
		switch handlerInfo.Method {
		case http.MethodGet:
			s.Router.Get(handlerInfo.Path, handlerInfo.Handler)
		case http.MethodPost:
			s.Router.Post(handlerInfo.Path, handlerInfo.Handler)
		case http.MethodPut:
			s.Router.Put(handlerInfo.Path, handlerInfo.Handler)
		case http.MethodDelete:
			s.Router.Delete(handlerInfo.Path, handlerInfo.Handler)
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
