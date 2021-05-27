package apiserver

// APIServer implements properties of the server
type APIServer struct {
	config *Config
}

// New creates a new instance of API server
func New(config *Config) *APIServer {

	return &APIServer{
		config: config,
	}
}

// Start actually starts the server
func (s *APIServer) Start() error {
	return nil
}
