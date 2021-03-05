package codegen

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"path"
	"testing"
)

func TestGen(t *testing.T) {

	parent := path.Join(toolbox.CallerDirectory(3), "test_data")

	var useCases = []struct {
		description string
		options     *Options
		hasError    bool
	}{
		{
			description: "message  types",
			options: &Options{
				Source: path.Join(parent, "messages"),
				Types:  []string{"Message"},
				Dest:   path.Join(parent, "messages"),
			},
		},
		{
			description: "basic struct code generation",
			options: &Options{
				Source: path.Join(parent, "basic_struct"),
				Types:  []string{"Message"},
				Dest:   path.Join(parent, "basic_struct", "encoding.go"),
			},
		},
		{
			description: "basic aliased types",
			options: &Options{
				Source: path.Join(parent, "primitive_alias"),
				Types:  []string{"Message"},
				Dest:   path.Join(parent, "primitive_alias", "encoding.go"),
			},
		},
	}

	for _, useCase := range useCases[:1] {
		err := Generate(useCase.options)
		assert.Nil(t, err, useCase.hasError)

	}
}
