package server

type ServiceName = string

// Server main structure
type Server struct {
	EndpointServices map[ServiceName]string
}