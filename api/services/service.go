package services

// A Service is the API entrypoint. It contains child services
// for different parts of the HTTP API.
type Service struct {
	Nodes NodesService
}

func RegisterServices() *Service {
	s := new(Service)

	s.Nodes = NewNodesService()

	return s
}
