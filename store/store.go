package store

type Store struct {
	Nodes NodesStore
}

func NewStore() *Store {
	s := new(Store)

	s.Nodes = NewNodesStore()

	return s
}
