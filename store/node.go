package store

import "golang.org/x/net/context"

type Node struct {
	name string
}

type NodesStore interface {
	// Create creates a new node.
	Create(ctx context.Context, in *Node) (*Node, error)
}

type nodesStore struct {
}

func NewNodesStore() NodesStore {
	return &nodesStore{}
}

func (s *nodesStore) Create(ctx context.Context, in *Node) (*Node, error) {
	// call the database
	return nil, nil
}
