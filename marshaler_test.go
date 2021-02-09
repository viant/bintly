package bintly

import (
	"encoding/binary"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	bin "github.com/viant/bintly/binary"
	"math"
	"reflect"
	"testing"
)

type benchStruct struct {
	A1 int
	A2 string
	A3 bool
	A4 float64
	A5 []int
	A6 []string
	A7 []float64
	A8 []byte
}

var b1 = &benchStruct{
	A1: math.MaxInt32,
	A2: "this is benchmark test",
	A3: true,
	A4: 3.3333,
	A5: intSlice,
	A6: stringSlice,
	A7: float64Slice,
	A8: uint8Slice,
}

func (m *benchStruct) DecodeBinary(stream *Reader) error {
	stream.Int(&m.A1)
	stream.String(&m.A2)
	stream.Bool(&m.A3)
	stream.Float64(&m.A4)
	stream.Ints(&m.A5)
	stream.Strings(&m.A6)
	stream.Float64s(&m.A7)
	stream.Uint8s(&m.A8)
	return nil
}

func (m *benchStruct) EncodeBinary(stream *Writer) error {
	stream.Int(m.A1)
	stream.String(m.A2)
	stream.Bool(m.A3)
	stream.Float64(m.A4)
	stream.Ints(m.A5)
	stream.Strings(m.A6)
	stream.Float64s(m.A7)
	stream.Uint8s(m.A8)
	return nil
}

func (m *benchStruct) ToBytes() ([]byte, error) {
	writer := bin.NewWriter(binary.LittleEndian)
	if err := writer.Int(m.A1); err != nil {
		return nil, err
	}
	if err := writer.String(m.A2); err != nil {
		return nil, err
	}
	if err := writer.Bool(m.A3); err != nil {
		return nil, err
	}
	if err := writer.Float64(m.A4); err != nil {
		return nil, err
	}
	if err := writer.Ints(m.A5); err != nil {
		return nil, err
	}
	if err := writer.Strings(m.A6); err != nil {
		return nil, err
	}
	if err := writer.Float64s(m.A7); err != nil {
		return nil, err
	}
	if err := writer.Bytes(m.A8); err != nil {
		return nil, err
	}
	return writer.ToBytes(), nil
}

func (m *benchStruct) FromBytes(bs []byte) {
	reader := bin.NewReader(bs, binary.LittleEndian)
	m.A1 = reader.Int()
	m.A2 = reader.String()
	m.A3 = reader.Bool()
	m.A4 = reader.Float64()
	m.A5 = reader.Ints()
	m.A6 = reader.Strings()
	m.A7 = reader.Float64s()
	m.A8 = reader.Bytes()
}

func Test_BenchStruct(t *testing.T) {
	data, err := b1.ToBytes()
	assert.Nil(t, err)
	b2 := &benchStruct{}
	b2.FromBytes(data)
	assert.EqualValues(t, b2, b1)
}

func BenchmarkUnmarshal(b *testing.B) {
	data, err := Marshal(b1)
	assert.Nil(b, err)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := benchStruct{}
		err = Unmarshal(data, &c1)
		if err != nil {
			assert.Nil(b, err)
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(b1)
	}
}

func BenchmarkUnmarshalBinary(b *testing.B) {
	data, _ := b1.ToBytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := benchStruct{}
		c1.FromBytes(data)
	}
}
func BenchmarkMarshalBinary(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b1.ToBytes()
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	data, err := json.Marshal(b1)
	assert.Nil(b, err)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := benchStruct{}
		err = json.Unmarshal(data, &c1)
		if err != nil {
			assert.Nil(b, err)
		}
	}
}
func BenchmarkJSONMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(b1)
	}
}

type marshStruct01 struct {
	ID          int
	OptionalID  *int
	Name        string
	Options     []string
	Elements    []int
	Coordinates []float64
}

func (m *marshStruct01) DecodeBinary(stream *Reader) error {
	stream.Int(&m.ID)
	stream.IntPtr(&m.OptionalID)
	stream.String(&m.Name)
	stream.Strings(&m.Options)
	stream.Ints(&m.Elements)
	stream.Float64s(&m.Coordinates)
	return nil
}

func (m *marshStruct01) EncodeBinary(stream *Writer) error {
	stream.Int(m.ID)
	stream.IntPtr(m.OptionalID)
	stream.String(m.Name)
	stream.Strings(m.Options)
	stream.Ints(m.Elements)
	stream.Float64s(m.Coordinates)
	return nil
}

type marshStruct02 struct {
	A1 int
	B1 *int
	C1 []int
	D1 uint
	E1 *uint
	F1 []uint

	A2 int64
	B2 *int64
	C2 []int64
	D2 uint64
	E2 *uint64
	F2 []uint64

	A3 int32
	B3 *int32
	C3 []int32
	D3 uint32
	E3 *uint32
	F3 []uint32

	A4 int16
	B4 *int16
	C4 []int16
	D4 uint16
	E4 *uint16
	F4 []uint16
	A5 int8
	B5 *int8
	C5 []int8
	D5 uint8
	E5 *uint8
	F5 []uint8
	G  bool
	H  string
}

func (m *marshStruct02) EncodeBinary(stream *Writer) error {
	stream.Int(m.A1)
	stream.IntPtr(m.B1)
	stream.Ints(m.C1)
	stream.Uint(m.D1)
	stream.UintPtr(m.E1)
	stream.Uints(m.F1)

	stream.Int64(m.A2)
	stream.Int64Ptr(m.B2)
	stream.Int64s(m.C2)
	stream.Uint64(m.D2)
	stream.Uint64Ptr(m.E2)
	stream.Uint64s(m.F2)

	stream.Int32(m.A3)
	stream.Int32Ptr(m.B3)
	stream.Int32s(m.C3)
	stream.Uint32(m.D3)
	stream.Uint32Ptr(m.E3)
	stream.Uint32s(m.F3)

	stream.Int16(m.A4)
	stream.Int16Ptr(m.B4)
	stream.Int16s(m.C4)
	stream.Uint16(m.D4)
	stream.Uint16Ptr(m.E4)
	stream.Uint16s(m.F4)

	stream.Int8(m.A5)
	stream.Int8Ptr(m.B5)
	stream.Int8s(m.C5)
	stream.Uint8(m.D5)
	stream.Uint8Ptr(m.E5)
	stream.Uint8s(m.F5)
	stream.Bool(m.G)
	stream.String(m.H)
	return nil
}

func (m *marshStruct02) DecodeBinary(stream *Reader) error {
	stream.Int(&m.A1)
	stream.IntPtr(&m.B1)
	stream.Ints(&m.C1)
	stream.Uint(&m.D1)
	stream.UintPtr(&m.E1)
	stream.Uints(&m.F1)

	stream.Int64(&m.A2)
	stream.Int64Ptr(&m.B2)
	stream.Int64s(&m.C2)
	stream.Uint64(&m.D2)
	stream.Uint64Ptr(&m.E2)
	stream.Uint64s(&m.F2)

	stream.Int32(&m.A3)
	stream.Int32Ptr(&m.B3)
	stream.Int32s(&m.C3)
	stream.Uint32(&m.D3)
	stream.Uint32Ptr(&m.E3)
	stream.Uint32s(&m.F3)

	stream.Int16(&m.A4)
	stream.Int16Ptr(&m.B4)
	stream.Int16s(&m.C4)
	stream.Uint16(&m.D4)
	stream.Uint16Ptr(&m.E4)
	stream.Uint16s(&m.F4)

	stream.Int8(&m.A5)
	stream.Int8Ptr(&m.B5)
	stream.Int8s(&m.C5)
	stream.Uint8(&m.D5)
	stream.Uint8Ptr(&m.E5)
	stream.Uint8s(&m.F5)
	stream.Bool(&m.G)
	stream.String(&m.H)
	return nil
}

func TestMarshal(t *testing.T) {

	var useCases = []struct {
		description string
		source      interface{}
	}{
		{
			description: "basic struct",
			source: &marshStruct01{
				ID:   100,
				Name: "test Me",
				Options: []string{
					"1",
					"20",
					"300",
				},
				Elements:    []int{30, 40, 50, 100, 1000, 2000},
				Coordinates: []float64{1.2, 0.2, 5.3, 6.3},
			},
		},
		{
			description: "large struct",
			source: &marshStruct02{
				A1: 101,
				B1: &intSlice[0],
				C1: intSlice,
				D1: 201,
				E1: &uintSlice[0],
				F1: uintSlice,
				A2: math.MinInt32,
				B2: &int64Slice[0],
				C2: int64Slice,
				D2: uint64slice[0],
				E2: &uint64slice[0],
				F2: uint64slice,
				A3: int32Slice[0],
				B3: &int32Slice[0],
				C3: int32Slice,
				D3: uint32Slice[0],
				E3: &uint32Slice[0],
				F3: uint32Slice,
				A4: 0,
				B4: nil,
				C4: nil,
				D4: uint16Slice[0],
				E4: &uint16Slice[0],
				F4: uint16Slice,
				A5: int8Slice[0],
				B5: &int8Slice[0],
				C5: int8Slice,
				D5: uint8Slice[0],
				E5: &uint8Slice[0],
				F5: uint8Slice,
				G:  false,
				H:  "test me",
			},
		},
	}

	for _, useCase := range useCases {
		data, err := Marshal(useCase.source)
		if !assert.Nil(t, err, useCase.description) {
			continue
		}
		var actual interface{}
		vType := reflect.TypeOf(useCase.source)
		if vType.Kind() == reflect.Ptr {
			vType = vType.Elem()
		}
		actual = reflect.New(vType).Interface()
		err = Unmarshal(data, actual)
		if !assert.Nil(t, err, useCase.description) {
			continue
		}
		assert.EqualValues(t, useCase.source, actual, useCase.description)
	}
}
