package sources

import (
	"errors"
	"net/url"
)

var (
	ErrAlreadyRegistered = errors.New("sources: already registered")
	ErrNotRegistered     = errors.New("sources: source not registered")

	registry map[string]InitFunc
)

type InitFunc func(path string) (Source, error)

func Register(name string, initFunc InitFunc) error {
	if _, ok := registry[name]; ok {
		return ErrAlreadyRegistered
	}
	registry[name] = initFunc
	return nil
}

func Get(source string) (Source, error) {
	path, err := url.Parse(source)
	if err != nil {
		return nil, err
	}
	sourceName, sourcePath := path.Scheme, path.Host+path.Path

	fn, ok := registry[sourceName]
	if !ok {
		return nil, ErrNotRegistered
	}
	return fn(sourcePath)
}

func init() {
	registry = map[string]InitFunc{}
}
