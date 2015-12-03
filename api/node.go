package api

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

type NodesService interface {
	// Create creates a new node.
	Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

type nodesService struct {
}

func NewNodesService() NodesService {
	return &nodesService{}
}

func (s *nodesService) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// decode json
	fmt.Println("hello")
	// call database methods from store
	return nil
}
