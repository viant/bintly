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
	return generateCoding(sess, typeName, encodeBaseType, func(sess *session, field *toolbox.FieldInfo) (string, error) {
		return "", fmt.Errorf("unsupported type: %s for field %v.%v", field.TypeName, typeName, field.Name)
	})
}

func generateStructDecoding(sess *session, typeName string) (string, error) {
	return generateCoding(sess, typeName, decodeBaseType, func(session *session, field *toolbox.FieldInfo) (string, error) {
		return "", fmt.Errorf("unsupported type: %s for field %v.%v", field.TypeName, typeName, field.Name)
	})

}

func generateCoding(session *session, typeName string, baseTemplate int, fn fieldGenerator) (string, error) {
	typeInfo := session.Type(typeName)
	if typeInfo == nil {
		return "", fmt.Errorf("failed to lookup '%s'", typeName)
	}
	var codings = make([]string, 0)
	fields := typeInfo.Fields()
	for _, field := range fields {
		if isBaseType(field.TypeName) {
			method := genCodingMethod(field)
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
		code, err := fn(session, field)
		if err != nil {
			return "", err
		}
		codings = append(codings, code)
	}
	return "\t" + strings.Join(codings, "\n\t"), nil
}

func genCodingMethod(field *toolbox.FieldInfo) string {
	codingMethod := strings.Title(field.TypeName)
	if field.IsPointer {
		codingMethod += "Ptr"
	}
	if field.IsSlice {
		codingMethod = codingMethod[2:]
		codingMethod += "s"
	}
	return codingMethod
}
