package codegen

import (
	"fmt"
	"github.com/viant/toolbox"
	"io/ioutil"
	"strings"
)

type fieldGenerator func(session *session, field *toolbox.FieldInfo) (string, error)

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

	if options.Dest == "" {
		fmt.Print(session.structCodingCode)
		return nil
	}
	return ioutil.WriteFile(options.Dest, []byte(strings.Join(session.structCodingCode, "")), 0644)

	return nil

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
	code, err = expandBlockTemplate(fileCode, struct {
		Pkg     string
		Code    string
		Imports string
	}{session.pkg, code, session.getImports()})

	session.structCodingCode = append(session.structCodingCode, code)
	//
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
	baseDerivedTemplate := encodeDerivedBaseType
	if isDecoder {
		baseTemplate = decodeBaseType
		baseDerivedTemplate = decodeDerivedBaseType
	}
	typeInfo := sess.Type(typeName)
	if typeInfo == nil {
		return "", fmt.Errorf("failed to lookup '%s'", typeName)
	}
	var codings = make([]string, 0)
	fields := typeInfo.Fields()
	for _, field := range fields {
		if isBaseType(field.TypeName) {
			method := genCodingMethod(field.TypeName, field.IsPointer, field.IsSlice)
			receiverAlias := strings.ToLower(typeName[0:1])
			code, err := expandFieldTemplate(baseTemplate, struct {
				Method        string
				Field         string
				ReceiverAlias string
			}{method, field.Name, receiverAlias})
			if err != nil {
				return "", err
			}
			codings = append(codings, code)
			continue
		}

		// base type aliasing
		baseType, err := getBaseDerivedType(sess, field.TypeName)
		if baseType != "" {

			method := genCodingMethod(baseType, field.IsPointer, field.IsSlice)
			receiverAlias := strings.ToLower(typeName[0:1])
			code, err := expandFieldTemplate(baseDerivedTemplate, struct {
				Method        string
				Field         string
				FieldType     string
				ReceiverAlias string
				TransientVar  string
				BaseType      string
			}{
				Method:        method,
				Field:         field.Name,
				FieldType:     field.TypeName,
				ReceiverAlias: receiverAlias,
				TransientVar:  toolbox.ToCaseFormat(field.Name, toolbox.CaseUpperCamel, toolbox.CaseLowerCamel),
				BaseType:      baseType,
			})
			if err != nil {
				return "", err
			}
			codings = append(codings, code)
			continue
		}

		// base slice type aliasing
		sliceType, err := getBaseSliceType(sess, field.TypeName)
		if sliceType != "" {

			method := genCodingMethod("[]"+sliceType, false, true)
			receiverAlias := strings.ToLower(typeName[0:1])
			code, err := expandFieldTemplate(baseDerivedTemplate, struct {
				Method        string
				Field         string
				FieldType     string
				ReceiverAlias string
				TransientVar  string
				BaseType      string
			}{
				Method:        method,
				Field:         field.Name,
				FieldType:     field.TypeName,
				ReceiverAlias: receiverAlias,
				TransientVar:  toolbox.ToCaseFormat(field.Name, toolbox.CaseUpperCamel, toolbox.CaseLowerCamel),
				BaseType:      sliceType,
			})
			if err != nil {
				return "", err
			}
			codings = append(codings, code)
			continue
		}

		code, err := fn(sess, field)
		if err != nil {
			return "", err
		}
		codings = append(codings, code)
	}
	return "\t" + strings.Join(codings, "\n\t"), nil
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

func genCodingMethod(baseType string, IsPointer bool, IsSlice bool) string {
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
