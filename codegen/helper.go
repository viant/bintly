package codegen

import (
	"github.com/viant/toolbox"
	"strings"
)

func firstLetterToUppercase(text string) string {
	return strings.ToUpper(string(text[0:1])) + string(text[1:])
}

func firstLetterToLowercase(text string) string {
	return strings.ToLower(string(text[0:1])) + string(text[1:])
}

func extractReceiverAlias(structType string) string {
	var result = string(structType[0])
	for i := len(structType) - 1; i > 0; i-- {
		aChar := string(structType[i])
		lowerChar := strings.ToLower(aChar)
		if lowerChar != aChar {
			result = lowerChar
			break
		}
	}
	return strings.ToLower(result)
}

func getSliceHelperTypeName(typeName string, isPointer bool) string {
	if typeName == "" {
		return ""
	}
	var pluralName = firstLetterToUppercase(typeName) + "s"
	if isPointer {
		pluralName += "Ptr"
	}
	return pluralName
}



func wrapperIfNeeded(text, wrappingChar string) string {
	if strings.HasPrefix(text, wrappingChar) {
		return text
	}
	return wrappingChar + text + wrappingChar
}


func getJSONKey(options *Options, field *toolbox.FieldInfo) string {
	var key = field.Name
	return key
}

func normalizeTypeName(typeName string) string {
	return strings.Replace(typeName, "*", "", strings.Count(typeName, "*"))
}
