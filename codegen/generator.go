package codegen

import (
	"github.com/viant/toolbox"
	"os"
	"path/filepath"
)

type Generator struct {
	fileInfo *toolbox.FileSetInfo
	options  *Options
	Pkg      string
}

func New(options *Options) *Generator {

	return &Generator{
		options: options,
	}
}

func (g *Generator) Generate() error {
	if err := g.options.Validate(); err != nil {
		return err
	}
	err := g.readPackageCode()
	if err != nil {
		return err
	}

	// then we generate code for the types given
	for _, rootType := range g.options.Types {
		if err := g.generateStructCode(rootType); err != nil {
			return err
		}
	}

	return nil

}

func (g *Generator) readPackageCode() error {
	p, err := filepath.Abs(g.options.Source)
	if err != nil {
		return err
	}

	var f os.FileInfo
	if f, err = os.Stat(p); err != nil {
		// path/to/whatever does not exist
		return err
	}

	if !f.IsDir() {
		g.Pkg = filepath.Dir(p)
		dir, _ := filepath.Split(p)
		g.fileInfo, err = toolbox.NewFileSetInfo(dir)

	} else {
		g.Pkg = filepath.Base(p)
		g.fileInfo, err = toolbox.NewFileSetInfo(p)
	}

	// if Pkg flag is set use it
	if g.options.Pkg != "" {
		g.Pkg = g.options.Pkg
	}

	return err
}

func (g *Generator) generateStructCode(rootType string) error {
	return nil
}
