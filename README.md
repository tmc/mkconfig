mkconfig
========

mkconfig is a way to render configuration files for services.

It is designed to be extensible in terms of sources of service information.

See github.com/tmc/mkconfig/sources/json/json.go for a basic json file backend.


## mkconfig
    go get github.com/tmc/mkconfig

Example invocation:


	mkconfig -source=json://./example/dummy.json -template=./example/dummy.tmpl

Output:


	# example template
	dummysvc:
	 - 10.0.0.1:8080
	 - 10.0.0.2:8080
	 - 10.0.0.3:8080
	 - 10.0.0.4:8080

dummy.tmpl:


	# example template
	dummysvc:
	{{ range service "dummysvc"}} - {{.Addr}}:{{.Port}}
	{{ end }}

dummy.json:


	{
	  "dummysvc": [
	    {
	      "name": "dummysvc1",
	      "port": 8080,
	      "addr": "10.0.0.1"
	    },
	    {
	      "name": "dummysvc2",
	      "port": 8080,
	      "addr": "10.0.0.2"
	    },
	    {
	      "name": "dummysvc3",
	      "port": 8080,
	      "addr": "10.0.0.3"
	    },
	    {
	      "name": "dummysvc4",
	      "port": 8080,
	      "addr": "10.0.0.4"
	    }
	  ]
	}

