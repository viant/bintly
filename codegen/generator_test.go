package codegen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"path"
	"strings"
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
			description: "basic struct code generation",
			options: &Options{
				Source: path.Join(parent, "basic_struct"),
				Types:  []string{"Message"},
				Dest:   path.Join(parent, "basic_struct", "encoding.go"),
			},
		},
	}

	for _, useCase := range useCases {
		filesetInfo, err := toolbox.NewFileSetInfo(useCase.options.Source)
		assert.Nil(t, err, useCase.hasError)
		GenFields(filesetInfo.Type("Message"))

	}
}

var typeMap = map[string]string{
	"int":      "Int",
	"[]int":    "Ints",
	"uint":     "Int",
	"[]uint":   "Ints",
	"string":   "String",
	"[]string": "Strings",
}

func GenFields(myType *toolbox.TypeInfo) {
	fmt.Printf("func (r *Receiver) EncodeBinary(stream *bintly.Writer) error {\n")
	fields := myType.Fields()

	for _, f := range fields {
		typeName := strings.Title(f.TypeName)
		if typeName[0:3] == "*[]" {
			//continue
			//TODO

			typeName = strings.Replace(typeName, "[]", "", 1)
			fmt.Printf(`size := int(stream.Alloc())
	r.%s = make([]*%s, size)
	for i := 0; i < size; i++ {
		%s := &%s{}
		if err := stream.Coder(%s); err != nil {
			return err
		}
		r.%s[i] = %s
	}`, f.Name, typeName, typeName, typeName, typeName, typeName, typeName)
		} else {
			if typeName[0:2] == "[]" {
				typeName = strings.Replace(typeName, "[]", "", 1) + "s"
			}
			if f.IsPointer {
				typeName += "Ptr"
			}
			fmt.Printf("\tenc.%s(r.%s)\n", typeName, f.Name)
			//		fmt.Printf("dec.%s(&p.%s)\n", typeName, f.Name)
		}
	}
	fmt.Printf("\treturn nil\n}\n")
}
