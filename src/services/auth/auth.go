package auth

// Auth interface
type Auth interface {
	Encode(obj interface{}) (string, error)
	Decode(token string) (interface{}, error)
}
