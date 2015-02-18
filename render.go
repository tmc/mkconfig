package mkconfig

import (
	"io"
	"io/ioutil"
	"os"

	"text/template"

	"github.com/tmc/mkconfig/sources"
)

func RenderPath(source sources.Source, templatePath, outputPath string, dryRun bool) error {
	f, err := os.Open(templatePath)
	if err != nil {
		return err
	}
	var output io.Writer = os.Stdout
	if outputPath != "" && dryRun == false {
		output, err = os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
	}
	return Render(source, f, output, dryRun)
}

func Render(source sources.Source, templateSrc io.Reader, output io.Writer, dryRun bool) error {
	src, err := ioutil.ReadAll(templateSrc)
	if err != nil {
		return err
	}
	t, err := template.New("mkconfig").Funcs(tmplFuncs(source)).Parse(string(src))
	if err != nil {
		return err
	}
	return t.Execute(output, nil)
}

func tmplFuncs(source sources.Source) template.FuncMap {
	return map[string]interface{}{
		"service": source.Service,
	}
}
