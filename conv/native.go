package conv

import "reflect"

var intType = reflect.TypeOf(0)
var uintType = reflect.TypeOf(uint(0))
var int64Type = reflect.TypeOf(int64(0))
var uint64Type = reflect.TypeOf(uint64(0))
var int32Type = reflect.TypeOf(int32(0))
var uint32Type = reflect.TypeOf(uint32(0))
var int16Type = reflect.TypeOf(int16(0))
var uint16Type = reflect.TypeOf(uint16(0))
var int8Type = reflect.TypeOf(int8(0))
var uint8Type = reflect.TypeOf(uint8(0))
var float64Type = reflect.TypeOf(float64(0))
var float32Type = reflect.TypeOf(float32(0))
var boolType = reflect.TypeOf(true)
var stringType = reflect.TypeOf("")
var bytesType = reflect.TypeOf([]byte{})

func IsNative(t reflect.Type) bool {
	switch t {
	case intType, uintType, int64Type, uint64Type, int32Type, uint32Type,
		int16Type, uint16Type, int8Type, uint8Type,
		float64Type, float32Type, boolType, stringType, bytesType:
		return true
	}
	return false
}

//MatchNative returns matched alias native type pointer or nil
func MatchNative(t reflect.Type) *reflect.Type {
	switch t.Kind() {
	case reflect.Int:
		return &intType
	case reflect.Uint:
		return &uintType
	case reflect.Int64:
		return &int64Type
	case reflect.Uint64:
		return &uint64Type
	case reflect.Int32:
		return &int32Type
	case reflect.Uint32:
		return &uint32Type
	case reflect.Int16:
		return &int16Type
	case reflect.Uint16:
		return &uint16Type
	case reflect.Int8:
		return &int8Type
	case reflect.Uint8:
		return &uint8Type
	case reflect.Float64:
		return &float64Type
	case reflect.Float32:
		return &float32Type
	case reflect.Bool:
		return &boolType
	case reflect.String:
		return &stringType
	}
	if IsBytesConvertibleTo(t) {
		return &bytesType
	}
	return nil
}

//IsIntConvertibleTo return true if int convertibleTo
func IsIntConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(intType)
}

//IsUintConvertibleTo return true if uint convertibleTo
func IsUintConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(uintType)
}

//IsInt64ConvertibleTo return true if int64 convertibleTo
func IsInt64ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(int64Type)
}

//IsUint64ConvertibleTo return true if uint64 convertibleTo
func IsUint64ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(uint64Type)
}

//IsInt32ConvertibleTo return true if int32 convertibleTo
func IsInt32ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(int32Type)
}

//IsUint32ConvertibleTo return true if uint32 convertibleTo
func IsUint32ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(uint32Type)
}

//IsInt16ConvertibleTo return true if int16 convertibleTo
func IsInt16ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(int16Type)
}

//IsUint16ConvertibleTo return true if uint16 convertibleTo
func IsUint16ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(uint16Type)
}

//IsInt8ConvertibleTo return true if int16 convertibleTo
func IsInt8ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(int8Type)
}

//IsUint8ConvertibleTo return true if uint16 convertibleTo
func IsUint8ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(uint8Type)
}

//IsFloat64ConvertibleTo return true if float64 convertibleTo
func IsFloat64ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(float64Type)
}

//IsFloat32ConvertibleTo return true if float32 convertibleTo
func IsFloat32ConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(float32Type)
}

//IsBoolConvertibleTo return true if bool convertibleTo
func IsBoolConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(boolType)
}

//IsByteConvertibleTo return true if []byte ConvertibleTo
func IsBytesConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(bytesType)
}

//IsStringConvertibleTo return true if string ConvertibleTo
func IsStringConvertibleTo(t reflect.Type) bool {
	return t.ConvertibleTo(stringType)
}
