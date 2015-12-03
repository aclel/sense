package main

import (
	"github.com/aclel/sense/store"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func NewRouter(db *store.DB, svc *Service) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Enable CORS to allow Cross Origin requests
	c := cors.New(cors.Options{
		// This can be uncommented to restrict CORS to only localhost:8080 and teamneptune.co
		//AllowedOrigins:   []string{"http://localhost:8080", "http://teamneptune.co"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	defaultChain := alice.New(c.Handler)
	r.Handle("/api/nodes", defaultChain.Then(svc.Nodes.Create))
	return r
}
