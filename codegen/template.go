package codegen

import (
	"bytes"
	"fmt"
	"text/template"
)

const (
	decodeBaseType = iota
	encodeBaseType
	decodeDerivedBaseType
	encodeDerivedBaseType
	decodeBaseSliceType
	encodeBaseSliceType
	decodeCustomSliceType
	encodeCustomSliceType
	encodeStructType
	decodeStructType
	encodeSliceStructType
	decodeSliceStructType
	encodeEmbeddedAliasTemplate
	decodeEmbeddedAliasSliceTemplate
	encodeBasicMapTemplate
	decodeBasicMapTemplate
	encodeSliceMapTemplate
	decodeSliceMapTemplate
)

var fieldTemplate = map[int]string{
	encodeBaseType: `	coder.{{.Method}}({{.ReceiverAlias}}.{{.Field}})`,
	decodeBaseType: `	coder.{{.Method}}(&{{.ReceiverAlias}}.{{.Field}})`,
	decodeDerivedBaseType: `	var {{.TransientVar}} {{.BaseType}}
	coder.{{.Method}}(&{{.TransientVar}})
	{{.ReceiverAlias}}.{{.Field}} = {{.FieldType}}({{.TransientVar}})`,
	encodeDerivedBaseType: `	coder.{{.Method}}({{.BaseType}}({{.ReceiverAlias}}.{{.Field}}))`,
	encodeBaseSliceType: `	coder.{{.Method}}(({{.ReceiverAlias}}.{{.Field}}))`,
	decodeBaseSliceType: `	var {{.TransientVar}} []{{.BaseType}}
	coder.{{.Method}}(&{{.TransientVar}})
	{{.ReceiverAlias}}.{{.Field}} = {{.TransientVar}}`,
	encodeCustomSliceType: `	coder.{{.Method}}(*(*[]{{.BaseType}})(unsafe.Pointer(&{{.ReceiverAlias}}.{{.Field}})))`,
	decodeCustomSliceType: `	var {{.TransientVar}} []{{.BaseType}}
	coder.{{.Method}}(&{{.TransientVar}})
	{{.ReceiverAlias}}.{{.Field}} = *(*{{.FieldType}})(unsafe.Pointer(&{{.TransientVar}}))`,
	encodeStructType: `	coder.{{.Method}}({{if .PointerNeeded}}&{{end}}{{.ReceiverAlias}}.{{.Field}})`,
	decodeStructType: `{{if not .PointerNeeded}}	{{.ReceiverAlias}}.{{.Field}} = &{{.FieldType}}{}
{{end}}	coder.{{.Method}}({{if .PointerNeeded}}&{{end}}{{.ReceiverAlias}}.{{.Field}})`,
	encodeSliceStructType: `	var {{.TransientVar}} = len({{.ReceiverAlias}}.{{.Field}})
	coder.Alloc(int32({{.TransientVar}}))
	for i:=0; i < {{.TransientVar}} ; i++ {
		if err := coder.{{.Method}}({{if .PointerNeeded}}&{{end}}{{.ReceiverAlias}}.{{.Field}}[i]);err !=nil {
			return nil
		}
	}`,
	decodeSliceStructType: `	var {{.TransientVar}} = coder.Alloc()
	{{.ReceiverAlias}}.{{.Field}} = make([]{{if not .PointerNeeded}}*{{end}}{{.FieldType}},{{.TransientVar}})
	for i:=0; i < int({{.TransientVar}}) ; i++ {
		if err := coder.{{.Method}}({{if .PointerNeeded}}&{{end}}{{.ReceiverAlias}}.{{.Field}}[i]);err != nil {
			return nil
		}
	}`,
	encodeEmbeddedAliasTemplate: `	{{if not .PointerNeeded}}{{.ReceiverAlias}}.{{.Field}} = &{{.FieldType}}{}{{end}}
	if err := coder.Coder({{if .PointerNeeded}}&{{end}}{{.ReceiverAlias}}.{{.Field}}); err !=nil {
	return err
	}`,
	decodeEmbeddedAliasSliceTemplate: `		if err := coder.Coder({{if .PointerNeeded}}&{{end}}{{.ReceiverAlias}}.{{.Field}}); err != nil {
		return err
	}`,
	encodeBasicMapTemplate: `	coder.Alloc(int32(len({{.ReceiverAlias}}.{{.Field}})))
	for k, v := range {{.ReceiverAlias}}.{{.Field}} {
		coder.{{.KeyMethod}}(k)
		coder.{{.ValueMethod}}({{if .PointerNeeded}}&{{end}}v)
	}`,
	decodeBasicMapTemplate: `	 {
		size := int(coder.Alloc())
		if size == bintly.NilSize {
			return nil
		}
		{{.ReceiverAlias}}.{{.Field}} = make(map[{{.KeyFieldType}}]{{.ValueFieldType}},size)
		for i:=0 ; i < size ; i++ {
			var k {{.KeyFieldType}}
			var v ={{if .PointerMethod}}&{{end}}{{.BaseValueFieldType}}{}
			coder.{{.KeyMethod}}(&k)
			coder.{{.ValueMethod}}({{if .PointerNeeded}}&{{end}}v)
			{{.ReceiverAlias}}.{{.Field}}[k]=v
		}
	}`,
	encodeSliceMapTemplate: ` {	
	coder.Alloc(int32(len({{.ReceiverAlias}}.{{.Field}})))
	for k, v := range {{.ReceiverAlias}}.{{.Field}} {
		coder.{{.KeyMethod}}(k)
		var m1 = len(v)
		coder.Alloc(int32(m1))
		for i:=0; i< m1;i++ {
			if err := coder.{{.ValueMethod}}({{if .PointerNeeded}}&{{end}}v[i]);err != nil {
				return nil
			}
		}
	}
	}`,
	decodeSliceMapTemplate: `	 {
		size := int(coder.Alloc())
		if size == bintly.NilSize {
			return nil
		}
		{{.ReceiverAlias}}.{{.Field}} = make(map[{{.KeyFieldType}}]{{.ValueFieldType}},size)
		for i:=0 ; i < size ; i++ {
			var k {{.KeyFieldType}}
			var v {{.ValueFieldType}}
			coder.{{.KeyMethod}}(&k)
			var m1Size = coder.Alloc()
			v = make({{.ValueFieldType}},m1Size)
			for j:=0; j < int(m1Size); j++ {	
				{{if .PointerMethod}}v[j]=&{{.BaseValueFieldType}}{}{{end}}
				if err := coder.{{.ValueMethod}}({{if .PointerNeeded}}&{{end}}v[j]);err !=nil {	
				 	return nil
				}
			}
			{{.ReceiverAlias}}.{{.Field}}[k]=v
		}
	}`,
}


const (
	fileCode = iota
	codingStructType
	codingSliceType
)

var blockTemplate = map[int]string{
	fileCode: `// Code generated by bintly codegen. DO NOT EDIT.\n\n

package {{.Pkg}}

import (
{{.Imports}}
)
{{.Code}}

`,
	codingStructType: `
func ({{.Receiver}}) EncodeBinary(coder *bintly.Writer) error {
{{.EncodingCases}}
	return nil
}
func ({{.Receiver}}) DecodeBinary(coder *bintly.Reader) error {
{{.DecodingCases}}	
	return nil
}
`,
	codingSliceType: `func ({{.ReceiverAlias}} *{{.SliceType}}) EncodeBinary(coder *bintly.Writer) error {
	var size = len(*{{.ReceiverAlias}})
	coder.Alloc(int32(size))
	for i:=0; i < size ; i++ {
		if err := coder.Coder({{if not .IsPointer}}&{{end}}(*{{.ReceiverAlias}})[i]);err !=nil {
			return nil
		}
	}
	return nil
}

func ({{.ReceiverAlias}} *{{.SliceType}}) DecodeBinary(coder *bintly.Reader) error  {
	var tmp = coder.Alloc()
	*{{.ReceiverAlias}} = make([]{{if .IsPointer}}*{{end}}{{.ComponentType}},tmp)
	for i:=0; i < int(tmp) ; i++ {
		tmp := 	{{if .IsPointer}}&{{end}}{{.ComponentType}}{}
		if err := coder.Coder({{if not .IsPointer}}&{{end}}(*{{.ReceiverAlias}})[i]);err != nil {
			return nil
		}
		(*{{.ReceiverAlias}})[i] = tmp
	}
	return nil
}
`,
}

//expandTemplate replaces templates parameters with actual data
func expandTemplate(namespace string, dictionary map[int]string, key int, data interface{}) (string, error) {
	var id = fmt.Sprintf("%v_%v", namespace, key)
	textTemplate, ok := dictionary[key]
	if !ok {
		return "", fmt.Errorf("failed to lookup template for %v.%v", namespace, key)
	}
	temlate, err := template.New(id).Parse(textTemplate)
	if err != nil {
		return "", fmt.Errorf("fiailed to parse template %v %v, due to %v", namespace, key, err)
	}
	writer := new(bytes.Buffer)
	err = temlate.Execute(writer, data)
	return writer.String(), err
}

//expandFieldTemplate replaces template fields with data
func expandFieldTemplate(key int, data interface{}) (string, error) {
	return expandTemplate("fieldTemplate", fieldTemplate, key, data)
}

//expandBlockTemplate replaces template block with data
func expandBlockTemplate(key int, data interface{}) (string, error) {
	return expandTemplate("blockTemplate", blockTemplate, key, data)
}
