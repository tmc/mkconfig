package sources

import "errors"

var (
	ErrAlreadyRegistered = errors.New("sources: already registered")
	ErrNotRegistered     = errors.New("sources: source not registered")

	registry map[string]InitFunc
)

type InitFunc func() (Source, error)

func Register(name string, initFunc InitFunc) error {
	if _, ok := registry[name]; ok {
		return ErrAlreadyRegistered
	}
	registry[name] = initFunc
	return nil
}

func Get(sourceName string) (Source, error) {
	fn, ok := registry[sourceName]
	if !ok {
		return nil, ErrNotRegistered
	}
	return fn()
}

func init() {
	registry = map[string]InitFunc{}
}
