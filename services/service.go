package services

// Service represents a network-accessible service.
type Service struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	Addr string `json:"addr"`
}
