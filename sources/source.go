package sources

import "github.com/tmc/mkconfig/services"

// Source provides a way to look up services by label
type Source interface {
	Service(label string) ([]services.Service, error)
}
