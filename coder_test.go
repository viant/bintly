package bintly

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStructCoder_DecodeBinary(t *testing.T) {

	type intAlias int
	ia := intAlias(3023)

	var useCases = []struct {
		description string
		value       interface{}
	}{
		{
			description: "int/uint types",
			value: struct {
				I   int
				U   uint
				Is  []int
				Uis []uint
			}{1, 2, intSlice, uintSlice},
		},
		{
			description: "int64/uint64 types",
			value: struct {
				I int64
				U uint64
			}{1000, 3000},
		},
		{
			description: "int32/uint32 types",
			value: struct {
				I int32
				U uint32
			}{-1500, 30777},
		},
		{
			description: "int16/uint16 types",
			value: struct {
				I int16
				U uint16
			}{-1544, 664},
		},
		{
			description: "int16/uint16 types",
			value: struct {
				I int16
				U uint16
			}{-15, 255},
		},
		{
			description: "float64/float32 types",
			value: struct {
				F1 float64
				F2 float32
			}{0.1, 3.222},
		},
		{
			description: "bool type",
			value: struct {
				B    bool
				BPtr *bool
			}{true, &boolSlice[0]},
		},
		{
			description: "alias type",
			value: struct {
				I     intAlias
				Bytes []byte
				Z     *intAlias
			}{102, []byte("123"), &ia},
		},
	}

	for _, useCase := range useCases {
		data, err := Marshal(useCase.value)
		if !assert.Nil(t, err, useCase.description) {
			continue
		}
		actualPtr := reflect.New(reflect.TypeOf(useCase.value))
		err = Unmarshal(data, actualPtr.Interface())
		if !assert.Nil(t, err, useCase.description) {
			continue
		}
		assert.EqualValues(t, useCase.value, actualPtr.Elem().Interface(), useCase.description)

	}

}