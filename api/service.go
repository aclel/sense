package api

// A Service is the API entrypoint. It contains child services
// for different parts of the HTTP API.
type Service struct {
	Nodes NodesService
}

func NewService() *Service {
	s := new(Service)

	s.Nodes = &NodesService{
		Nodes: NewNodesService()
	}

	return s
}
