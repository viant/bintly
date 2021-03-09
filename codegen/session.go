package codegen

import (
	"github.com/viant/toolbox"
	"os"
	"path/filepath"
	"strings"
)

type session struct {
	*Options
	*toolbox.FileSetInfo
	pkg              string
	structCodingCode []string
	generatedTypes   map[string]bool
	imports          map[string]bool
	isBlockTemplateDone bool
}

func (s *session) shallGenerateCode(typeName string) bool {
	if _, ok := s.generatedTypes[typeName]; ok {
		return false
	}
	s.generatedTypes[typeName] = true
	return true
}

func (s *session) readPackageCode() error {
	p, err := filepath.Abs(s.Source)
	if err != nil {
		return err
	}

	var f os.FileInfo
	if f, err = os.Stat(p); err != nil {
		// path/to/whatever does not exist
		return err
	}

	if !f.IsDir() {
		s.Pkg = filepath.Dir(p)
		dir, _ := filepath.Split(p)
		s.FileSetInfo, err = toolbox.NewFileSetInfo(dir)

	} else {
		s.Pkg = filepath.Base(p)
		s.FileSetInfo, err = toolbox.NewFileSetInfo(p)
	}

	// if Pkg flag is set use it
	if s.Pkg != "" {
		s.pkg = s.Pkg
	}

	return err
}

func newSession(option *Options) *session {
	return &session{Options: option,
		structCodingCode: make([]string, 0),
		generatedTypes:   make(map[string]bool),
		imports:          make(map[string]bool),
	}
}

func (s *session) addImport(pkg string) {
	s.imports[`"`+pkg+`"`] = true
}

func (s *session) getImports() string {
	return "\t" + strings.Join(toolbox.MapKeysToStringSlice(s.imports), "\n\t")
}
