package sources

import "github.com/tmc/mkconfig/services"

type Source interface {
	Service(name string) ([]services.Service, error)
}
