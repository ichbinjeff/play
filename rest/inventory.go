package rest

import (
	"fmt"
	"net/http"
)

func init() {
	Endpoints{
		getInventory,
	}.MustRegister()
}

var getInventory = Endpoint{
	Name:   "getInventory",
	Url:    "/inventory",
	Method: http.MethodGet,
	Handler: GetInventoryHandler{
		middlewares: []Middleware{
			func() {
				fmt.Println("first middleware")
			},
			func() {
				fmt.Println("second middleware")
			},
		},
	},
}

type Middleware func()

type GetInventoryHandler struct {
	middlewares []Middleware
}

func (inventory GetInventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, middleware := range inventory.middlewares {
		middleware()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"foo": "bar"}`))
}
