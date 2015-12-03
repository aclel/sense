package services

import (
	"net/http"

	"github.com/aclel/sense/store"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"golang.org/x/net/context"
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

	ctx := context.Background()
	context.WithValue(ctx, "db", db)

	r.Handle("/api/nodes", defaultChain.Then(AppHandler{ctx, svc.Nodes.Create})).Methods("POST", "OPTIONS")

	return r
}

type AppHandler struct {
	ctx    context.Context
	handle func(context.Context, http.ResponseWriter, *http.Request) error
}

func (appHandler AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := appHandler.handle(appHandler.ctx, w, r); err != nil {
		panic(err)
	}
}
