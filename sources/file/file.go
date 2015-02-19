// Package file provides a mkconfig.sources backend that reads from a very basic flat file format
//
// Example contents of example.cfg:
//
//   host1,10.0.0.1,8080
//   host2,10.0.0.2,8080
package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/tmc/mkconfig/services"
	"github.com/tmc/mkconfig/sources"
)

type fileSource struct {
	path     string
	services map[string][]services.Service
}

func init() {
	sources.Register("file", mkFileSource)
}

func mkFileSource(path string) (sources.Source, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fs := &fileSource{
		path:     path,
		services: make(map[string][]services.Service),
	}
	return fs, fs.parseHosts(f)
}

func (fs *fileSource) Service(serviceName string) ([]services.Service, error) {
	services, ok := fs.services[serviceName]
	if !ok {
		return nil, fmt.Errorf("file backend (%s): service '%s' not found.", fs.path, serviceName)
	}
	return services, nil
}

func (fs *fileSource) parseHosts(src io.Reader) error {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(buf), "\n") {
		var svc services.Service
		parts := strings.SplitN(line, ",", 3)
		switch len(parts) {
		case 0, 1:
			continue
		case 3:
			svc.Port, _ = strconv.Atoi(parts[2])
			fallthrough
		default:
			svc.Addr = parts[1]
			svc.Name = parts[0]
		}

		serviceName := strings.TrimRight(svc.Name, "0123456789")
		if _, ok := fs.services[serviceName]; !ok {
			fs.services[serviceName] = make([]services.Service, 0)
		}

		fs.services[serviceName] = append(fs.services[serviceName], svc)

	}
	return nil

}
