// Package json_source provides a mkconfig.sources backend that reads from a very basic json format
//
// Example contents of example.json:
//
//   {
//     "dummysvc": [
//       {
//         "name": "dummysvc1",
//         "port": 8080,
//         "addr": "10.0.0.1"
//       },
//       {
//         "name": "dummysvc2",
//         "port": 8080,
//         "addr": "10.0.0.2"
//       },
//       {
//         "name": "dummysvc3",
//         "port": 8080,
//         "addr": "10.0.0.3"
//       },
//       {
//         "name": "dummysvc4",
//         "port": 8080,
//         "addr": "10.0.0.4"
//       }
//     ]
//   }
package json_source

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tmc/mkconfig/services"
	"github.com/tmc/mkconfig/sources"
)

type jsonSource struct {
	path     string
	services map[string][]services.Service
}

func init() {
	sources.Register("json", mkJSONSource)
}

func mkJSONSource(path string) (sources.Source, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fs := &jsonSource{
		path:     path,
		services: make(map[string][]services.Service),
	}
	return fs, fs.parseHosts(f)
}

func (fs *jsonSource) Service(serviceName string) ([]services.Service, error) {
	services, ok := fs.services[serviceName]
	if !ok {
		return nil, fmt.Errorf("file backend (%s): service '%s' not found.", fs.path, serviceName)
	}
	return services, nil
}

func (fs *jsonSource) parseHosts(src io.Reader) error {
	return json.NewDecoder(src).Decode(&fs.services)
}
