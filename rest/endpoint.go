package rest

import (
	"net/http"

	"github.com/pkg/errors"
)

// RegisteredEndpoints is a global list that the http handler will loop into and register the handler accordingly
var RegisteredEndpoints []Endpoint

type Endpoints []Endpoint

func (eps Endpoints) MustRegister() {
	for _, ep := range eps {
		RegisteredEndpoints = append(RegisteredEndpoints, ep)
	}
}

// Endpoint is an abstract object that stores endpoint metadata
type Endpoint struct {
	Method  string       `json:"method"`
	Url     string       `json:"url"`
	Name    string       `json:"name"`
	Handler http.Handler `json:"-"`
}

func (ep Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	if ep.Handler == nil {
		return errors.Errorf("doesn't have handler %s", ep.Name)
	}
	ep.Handler.ServeHTTP(w, r)
	return nil
}
