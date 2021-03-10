package codegen

import (
	"fmt"
	"github.com/viant/toolbox"
	"io/ioutil"
	"strings"
)

type fieldGenerator func(session *session, field *toolbox.FieldInfo) (string, error)

type templateParameters struct {
	Method        string
	Field         string
	FieldType     string
	ReceiverAlias string
	TransientVar  string
	BaseType      string
	PointerNeeded bool
}

func Generate(options *Options) error {

	if err := options.Validate(); err != nil {
		return err
	}
	session := newSession(options)

	//
	session.addImport("github.com/viant/bintly")
	err := session.readPackageCode()
	if err != nil {
		return err
	}

	// then we generate code for the types given
	for _, rootType := range options.Types {
		if err := generateStructCoding(session, rootType); err != nil {
			return err
		}
	}

	prefix, err := expandBlockTemplate(fileCode, struct {
		Pkg     string
		Code    string
		Imports string
	}{session.pkg, "", session.getImports()})
	if err != nil {
		return err
	}

	dest := session.Dest
	err = ioutil.WriteFile(dest, []byte(prefix+strings.Join(session.structCodingCode, "")), 0644)
	session.structCodingCode = []string{}
	return err

}

func generateStructCoding(session *session, typeName string) error {
	if ok := session.shallGenerateCode(typeName); !ok {
		return nil
	}
	enc, err := generateStructEncoding(session, typeName)
	if err != nil {
		return err
	}
	dec, err := generateStructDecoding(session, typeName)
	if err != nil {
		return err
	}

	receiver := strings.ToLower(typeName[0:1]) + " *" + typeName

	code, err := expandBlockTemplate(codingStructType, struct {
		Receiver      string
		EncodingCases string
		DecodingCases string
	}{receiver, enc, dec})
	if err != nil {
		return err
	}
	session.structCodingCode = append(session.structCodingCode, code)

	return nil
}

func generateStructEncoding(sess *session, typeName string) (string, error) {
	return generateCoding(sess, typeName, false, func(sess *session, field *toolbox.FieldInfo) (string, error) {
		return "", fmt.Errorf("unsupported type: %s for field %v.%v", field.TypeName, typeName, field.Name)
	})
}

func generateStructDecoding(sess *session, typeName string) (string, error) {
	return generateCoding(sess, typeName, true, func(session *session, field *toolbox.FieldInfo) (string, error) {
		return "", fmt.Errorf("unsupported type: %s for field %v.%v", field.TypeName, typeName, field.Name)
	})

}

func generateCoding(sess *session, typeName string, isDecoder bool, fn fieldGenerator) (string, error) {

	baseTemplate := encodeBaseType
	derivedTemplate := encodeDerivedBaseType
	baseSliceTemplate := encodeBaseSliceType
	derivedSliceTemplate := encodeCustomSliceType
	structTemplate := encodeStructType
	customSliceTemplate := encodeSliceStructType
	enbeddedAliasSliceTemplate := encodeEmbeddedAliasTemplate
	if isDecoder {
		baseTemplate = decodeBaseType
		derivedTemplate = decodeDerivedBaseType
		baseSliceTemplate = decodeBaseSliceType
		derivedSliceTemplate = decodeCustomSliceType
		structTemplate = decodeStructType
		customSliceTemplate = decodeSliceStructType
		enbeddedAliasSliceTemplate = decodeEmbeddedAliasSliceTemplate
	}
	typeInfo := sess.Type(typeName)
	if typeInfo == nil {
		return "", fmt.Errorf("failed to lookup '%s'", typeName)
	}
	var codings = make([]string, 0)
	fields := typeInfo.Fields()
	for _, field := range fields {

		receiverAlias := strings.ToLower(typeName[0:1])
		var generated bool
		var err error
		// base type
		if generated, err = generateBaseType(field, baseTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}

		// derived type
		if generated, err = generateDerivedType(sess, field, derivedTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}

		// base slice type
		if generated, err = generateBaseSliceType(sess, field, baseSliceTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}
		// derived slice type
		if generated, err = generateDerivedSliceType(sess, field, derivedSliceTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}

		// struct type
		fieldType := sess.Type(getBaseFieldType(field.TypeName))
		if fieldType == nil {
			return "", fmt.Errorf("unsupported field type %v for field %v", fieldType, field)
		}
		if generated, err = generateStructType(sess, fieldType, field, structTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}

		// struct slice
		if generated, err = generateSliceOfStruct(sess, field, customSliceTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}

		// alias slice
		if generated, err = generateSliceAlias(sess, fieldType, field, enbeddedAliasSliceTemplate, receiverAlias, &codings); err != nil {
			return "", err
		}
		if generated {
			continue
		}
		code, err := fn(sess, field)
		if err != nil {
			return "", err
		}
		codings = append(codings, code)

	}
	return strings.Join(codings, "\n"), nil
}

func generateSliceAlias(sess *session, fieldType *toolbox.TypeInfo, field *toolbox.FieldInfo, templateId int, receiverAlias string, codings *[]string) (bool, error) {

	if fieldType.IsSlice && !isInlineSliceType(field.TypeName) {
		if err := generateStructCoding(sess, fieldType.ComponentType); err != nil {
			return false, err
		}

		var code string
		err := generateSliceCoding(sess, field.TypeName, fieldType.IsPointerComponentType, fieldType.ComponentType)
		if err != nil {
			return false, err
		}

		code, err = expandFieldTemplate(templateId, templateParameters{
			Method:        "Coder",
			Field:         field.Name,
			ReceiverAlias: receiverAlias,
			PointerNeeded: !field.IsPointer,
			FieldType:     field.TypeName,
		})
		if err != nil {
			return false, err
		}
		*codings = append(*codings, code)
		return true, nil
	}
	return false, nil
}

func generateSliceCoding(sess *session, typeName string, isPointer bool, componentType string) error {
	if ok := sess.shallGenerateCode(typeName); !ok {
		return nil
	}
	code, err := expandBlockTemplate(codingSliceType, struct {
		ReceiverAlias string
		SliceType     string
		IsPointer     bool
		ComponentType string
	}{
		ReceiverAlias: strings.ToLower(typeName[0:1]),
		SliceType:     typeName,
		IsPointer:     isPointer,
		ComponentType: componentType,
	})
	if err != nil {
		return err
	}
	sess.structCodingCode = append(sess.structCodingCode, code)
	return nil
}

func generateSliceOfStruct(sess *session, field *toolbox.FieldInfo, customSliceTemplate int, receiverAlias string, codings *[]string) (bool, error) {
	if field.IsSlice && isInlineSliceType(field.TypeName) {
		if err := generateStructCoding(sess, field.ComponentType); err != nil {
			return false, err
		}
		fieldTypeName := field.ComponentType
		code, err := expandFieldTemplate(customSliceTemplate, templateParameters{
			Method:        "Coder",
			Field:         field.Name,
			FieldType:     fieldTypeName,
			ReceiverAlias: receiverAlias,
			TransientVar:  toolbox.ToCaseFormat(field.Name, toolbox.CaseUpperCamel, toolbox.CaseLowerCamel),
			PointerNeeded: !field.IsPointerComponent,
		})
		if err != nil {
			return false, err
		}
		*codings = append(*codings, code)
		return true, nil
	}
	return false, nil
}

func generateStructType(sess *session, fieldType *toolbox.TypeInfo, field *toolbox.FieldInfo, structTemplate int, receiverAlias string, codings *[]string) (bool, error) {
	if isStruct(fieldType) && !field.IsSlice && !field.IsPointerComponent {
		if err := generateStructCoding(sess, fieldType.Name); err != nil {
			return false, err
		}
		code, err := expandFieldTemplate(structTemplate, templateParameters{
			Method:        "Coder",
			Field:         field.Name,
			FieldType:     field.TypeName,
			ReceiverAlias: receiverAlias,
			PointerNeeded: !field.IsPointer,
		})
		if err != nil {
			return false, err
		}
		*codings = append(*codings, code)
		return true, nil
	}
	return false, nil
}

func generateDerivedSliceType(sess *session, field *toolbox.FieldInfo, derivedSliceTemplate int, receiverAlias string, codings *[]string) (bool, error) {
	customSliceType, err := getDerivedSliceType(sess, field.TypeName)
	if customSliceType == "" {
		return false, nil
	}
	sess.addImport("unsafe")
	method := genCodingMethod("[]"+customSliceType, false, true)
	code, err := expandFieldTemplate(derivedSliceTemplate, templateParameters{
		Method:        method,
		Field:         field.Name,
		FieldType:     field.TypeName,
		ReceiverAlias: receiverAlias,
		TransientVar:  toolbox.ToCaseFormat(field.Name, toolbox.CaseUpperCamel, toolbox.CaseLowerCamel),
		BaseType:      customSliceType,
	})
	if err != nil {
		return false, err
	}
	*codings = append(*codings, code)
	return true, nil
}

func generateBaseSliceType(sess *session, field *toolbox.FieldInfo, baseSliceTemplate int, receiverAlias string, codings *[]string) (bool, error) {
	sliceType, err := getBaseSliceType(sess, field.TypeName)
	if sliceType == "" {
		return false, nil
	}
	method := genCodingMethod("[]"+sliceType, false, true)
	code, err := expandFieldTemplate(baseSliceTemplate, templateParameters{
		Method:        method,
		Field:         field.Name,
		FieldType:     field.TypeName,
		ReceiverAlias: receiverAlias,
		TransientVar:  toolbox.ToCaseFormat(field.Name, toolbox.CaseUpperCamel, toolbox.CaseLowerCamel),
		BaseType:      sliceType,
	})
	if err != nil {
		return false, err
	}
	*codings = append(*codings, code)
	return true, nil
}

func generateDerivedType(sess *session, field *toolbox.FieldInfo, derivedTemplate int, receiverAlias string, codings *[]string) (bool, error) {
	baseType, err := getBaseDerivedType(sess, field.TypeName)
	if baseType == "" {
		return false, nil
	}
	method := genCodingMethod(baseType, field.IsPointer, field.IsSlice)
	code, err := expandFieldTemplate(derivedTemplate, templateParameters{
		Method:        method,
		Field:         field.Name,
		FieldType:     field.TypeName,
		ReceiverAlias: receiverAlias,
		TransientVar:  toolbox.ToCaseFormat(field.Name, toolbox.CaseUpperCamel, toolbox.CaseLowerCamel),
		BaseType:      baseType,
	})
	if err != nil {
		return false, err
	}
	*codings = append(*codings, code)
	return true, nil
}

func generateBaseType(field *toolbox.FieldInfo, baseTemplate int, receiverAlias string, codings *[]string) (bool, error) {
	if !isBaseType(field.TypeName) {
		return false, nil
	}
	method := genCodingMethod(field.TypeName, field.IsPointer, field.IsSlice)
	code, err := expandFieldTemplate(baseTemplate, templateParameters{
		Method:        method,
		Field:         field.Name,
		ReceiverAlias: receiverAlias,
	})
	if err != nil {
		return false, err
	}
	*codings = append(*codings, code)
	return true, nil
}

func isInlineSliceType(fieldType string) bool {
	return strings.Contains(fieldType, "[]")
}

func getBaseDerivedType(s *session, typeName string) (string, error) {
	aType := s.Type(typeName)
	if aType == nil {
		return "", fmt.Errorf("alias type name %v is nil for type %v ", aType, typeName)
	}
	if aType.IsDerived {
		derived := aType.Derived
		if isBaseType(derived) {
			return derived, nil
		}
		derived, err := getBaseDerivedType(s, derived)
		if err != nil {
			return "", err
		}
		if isBaseType(derived) {
			return derived, nil
		}
	}
	return "", nil
}

func getBaseSliceType(s *session, typeName string) (string, error) {
	aType := s.Type(typeName)
	if aType == nil {
		return "", fmt.Errorf("alias type name %v is nil for type %v ", aType, typeName)
	}
	if aType.IsSlice && isBaseType(aType.ComponentType) {
		return aType.ComponentType, nil
	}
	return "", nil
}

func getDerivedSliceType(s *session, typeName string) (string, error) {
	aType := s.Type(typeName)
	if aType == nil {
		return "", fmt.Errorf("alias type name %v is nil for type %v ", aType, typeName)
	}

	if aType.IsSlice {
		cType, err := getBaseDerivedType(s, aType.ComponentType)
		if err != nil {
			return "", fmt.Errorf("can't find base type %v for componentType %v ", aType, aType.ComponentType)
		}
		return cType, nil
	}

	return "", nil
}

func genCodingMethod(baseType string, IsPointer bool, IsSlice bool) string {
	if strings.Contains(baseType, "time.Time") {
		baseType = strings.Replace(baseType, "time.", "", 1)
	}
	codingMethod := strings.Title(baseType)
	if IsPointer {
		codingMethod += "Ptr"
	}
	if IsSlice {
		codingMethod = codingMethod[2:]
		codingMethod += "s"
	}
	return codingMethod

}

func getBaseFieldType(fieldType string) string {
	if !strings.Contains(fieldType, "[]") {
		return fieldType
	}
	if fieldType[0:3] == "[]*" {
		return fieldType[3:]
	}
	if fieldType[0:2] == "[]" {
		return fieldType[2:]
	}
	return fieldType
}
