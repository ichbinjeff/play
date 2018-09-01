package main

import (
	"net/http"
	"time"

	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/play/rest"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	for _, ep := range rest.RegisteredEndpoints {
		router.Methods(ep.Method).Name(ep.Name).Handler(ep.Handler)
	}
	fmt.Println("starting server")
	handler := (handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.ExposedHeaders([]string{"Content-Disposition"}),
		handlers.AllowCredentials(),
		handlers.MaxAge(600),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})))(router)

	httpServer := http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:9200",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server started, my friend")
	httpServer.ListenAndServe()
}
