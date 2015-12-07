package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"golang.org/x/net/context"
)

// NewRouter registers all HTTP routes and returns a mux Router
func NewRouter(ctx context.Context) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Enable CORS to allow Cross Origin requests
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	// Default middleware chain. CORS ->
	defaultChain := alice.New(c.Handler)

	// Register http handlers
	services := RegisterServices()

	// Setup routes
	r.Handle("/api/nodes", defaultChain.Then(appHandler{ctx, services.Nodes.Create})).Methods("POST", "OPTIONS")

	return r
}

type appHandler struct {
	ctx    context.Context
	handle func(context.Context, http.ResponseWriter, *http.Request) error
}

func (appHandler appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := appHandler.handle(appHandler.ctx, w, r); err != nil {
		panic(err)
	}
}
