package stress

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"github.com/fxamacker/cbor"
	"github.com/vmihailenco/msgpack"
	"github.com/stretchr/testify/assert"
	"github.com/viant/bintly"
	bin "github.com/viant/bintly/binary"
	"math"
	"testing"
)

var b1 = &BenchStruct{
	A1: math.MaxInt32,
	A2: "this is benchmark test",
	A3: true,
	A4: 3.3333,
	A5: intSlice,
	A6: stringSlice,
	A7: float64Slice,
	A8: uint8Slice,
}

//DecodeBinary decode bindly stream
func (m *BenchStruct) DecodeBinary(stream *bintly.Reader) error {
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

//EncodeBinary encodes bintly stream
func (m *BenchStruct) EncodeBinary(stream *bintly.Writer) error {
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

//ToBytes converts to bytes with wrapped encoding/binary ByteOrder
func (m *BenchStruct) ToBytes() ([]byte, error) {
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

//ToBytes converts from bytes with wrapped encoding/binary ByteOrder
func (m *BenchStruct) FromBytes(bs []byte) {
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
	b2 := &BenchStruct{}
	b2.FromBytes(data)
	assert.EqualValues(t, b2, b1)

	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(b1)
	assert.Nil(t, err)
	c1 := &BenchStruct{}
	err = dec.Decode(c1)
	assert.Nil(t, err)
	assert.EqualValues(t, b1, c1)

}

func BenchmarkUnmarshalBintly(b *testing.B) {
	data, err := bintly.Marshal(b1)
	assert.Nil(b, err)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := BenchStruct{}
		err = bintly.Unmarshal(data, &c1)
		if err != nil {
			assert.Nil(b, err)
		}
	}
}

func BenchmarkMarshalBintly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = bintly.Marshal(b1)
	}
}

func BenchmarkUnmarshalBinary(b *testing.B) {
	data, _ := b1.ToBytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := BenchStruct{}
		c1.FromBytes(data)
	}
}

func BenchmarkMarshalBinary(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = b1.ToBytes()
	}
}

func BenchmarkMarshalGob(b *testing.B) {
	var buf bytes.Buffer
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(b1)
		if err != nil {
			assert.NotNil(b, err)
		}
	}
}

func BenchmarkMarshalCbor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = cbor.Marshal(b1)
	}
}

func BenchmarkUnmarshalCbor(b *testing.B) {
	data, err := cbor.Marshal(b1)
	if !assert.Nil(b, err) {
		return
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := BenchStruct{}
		cbor.Unmarshal(data, &c1)
	}
}

func BenchmarkMarshalMsgPack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(b1)
	}
}

func BenchmarkUnmarshalMsgPack(b *testing.B) {
	data, err := msgpack.Marshal(b1)
	if !assert.Nil(b, err) {
		return
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := BenchStruct{}
		msgpack.Unmarshal(data, &c1)
	}
}

func BenchmarkUnMarshalGob(b *testing.B) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b1)
	if err != nil {
		assert.NotNil(b, err)
	}
	data := buf.Bytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := BenchStruct{}
		dec := gob.NewDecoder(bytes.NewReader(data))
		err := dec.Decode(&c1)
		if err != nil {
			assert.NotNil(b, err)
		}
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	data, err := json.Marshal(b1)
	assert.Nil(b, err)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1 := BenchStruct{}
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
