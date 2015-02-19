// Package sources provides an abstraction of services and a registry for additional source backends
//
// A source is defined by the Source interface:
//
//   type Source interface {
//   	Service(name string) ([]services.Service, error)
//   }
package sources
