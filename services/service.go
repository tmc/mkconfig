package services

type Service struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	Addr string `json:"addr"`
}
