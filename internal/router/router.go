package router

import (
	"github.com/gorilla/mux"
	"github.com/ruspatrick/tasks/internal/controllers"
	"net/http"
)

type route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func New(c *controllers.Controller) *mux.Router {
	routes := []route{
		{Name: "create", Method: http.MethodPost, Path: "/", Handler: c.CreateTask},
		{Name: "get", Method: http.MethodGet, Path: "/{id:[0-9]+}", Handler: c.GetTask},
	}

	r := mux.NewRouter()
	for _, route := range routes {
		r.Name(route.Name).Methods(route.Method).Path(route.Path).HandlerFunc(route.Handler)
	}

	return r
}
