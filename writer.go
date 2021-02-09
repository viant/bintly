package bintly

import (
	"fmt"
	"sync"
	"time"
)

type (
	//Writer represents binary writer
	Writer struct {
		alloc encUint32s
		encInts
		encUints
		encInt64s
		encUint64s
		encInt32s
		encUint32s
		encInt16s
		encUint16s
		encInt8s
		encUint8s
		encFloat64s
		encFloat32s
	}

	encInts     []int
	encUints    []uint
	encInt64s   []int64
	encUint64s  []uint64
	encInt32s   []int32
	encUint32s  []uint32
	encInt16s   []int16
	encUint16s  []uint16
	encInt8s    []int8
	encUint8s   []uint8
	encFloat64s []float64
	encFloat32s []float32
)

//Any writes any supported writer type
func (w *Writer) Any(v interface{}) error {
	switch actual := v.(type) {
	case int:
		w.Int(actual)
	case *int:
		w.IntPtr(actual)
	case []int:
		w.Ints(actual)
	case uint:
		w.Uint(actual)
	case *uint:
		w.UintPtr(actual)
	case []uint:
		w.Uints(actual)
	case int64:
		w.Int64(actual)
	case *int64:
		w.Int64Ptr(actual)
	case []int64:
		w.Int64s(actual)
	case uint64:
		w.Uint64(actual)
	case *uint64:
		w.Uint64Ptr(actual)
	case []uint64:
		w.Uint64s(actual)
	case int32:
		w.Int32(actual)
	case *int32:
		w.Int32Ptr(actual)
	case []int32:
		w.Int32s(actual)
	case uint32:
		w.Uint32(actual)
	case *uint32:
		w.Uint32Ptr(actual)
	case []uint32:
		w.Uint32s(actual)
	case int16:
		w.Int16(actual)
	case *int16:
		w.Int16Ptr(actual)
	case []int16:
		w.Int16s(actual)
	case uint16:
		w.Uint16(actual)
	case *uint16:
		w.Uint16Ptr(actual)
	case []uint16:
		w.Uint16s(actual)
	case int8:
		w.Int8(actual)
	case *int8:
		w.Int8Ptr(actual)
	case []int8:
		w.Int8s(actual)
	case uint8:
		w.Uint8(actual)
	case *uint8:
		w.Uint8Ptr(actual)
	case []uint8:
		w.Uint8s(actual)
	case float64:
		w.Float64(actual)
	case *float64:
		w.Float64Ptr(actual)
	case []float64:
		w.Float64s(actual)
	case float32:
		w.Float32(actual)
	case *float32:
		w.Float32Ptr(actual)
	case []float32:
		w.Float32s(actual)
	case bool:
		w.Bool(actual)
	case *bool:
		w.BoolPtr(actual)
	case string:
		w.String(actual)
	case *string:
		w.StringPtr(actual)
	case []string:
		w.Strings(actual)
	case time.Time:
		w.Time(actual)
	case *time.Time:
		w.TimePtr(actual)
	default:
		encoder, ok := v.(Encoder)
		if ok {
			return w.Coder(encoder)
		}
		return fmt.Errorf("unsupproted writer type: %T", v)
	}
	return nil
}

//Alloc append data allocation size for repeater or pointers(0,1) types
func (w *Writer) Alloc(size uint32) {
	w.alloc.Uint32(size)
}

//IntPtr writes *int
func (w *Writer) IntPtr(v *int) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Int(*v)
}

//Ints writes []int
func (w *Writer) Ints(vs []int) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendInts(vs)
}

//UintPtr writes *uint
func (w *Writer) UintPtr(v *uint) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Uint(*v)
}

//Uints writes []uint
func (w *Writer) Uints(vs []uint) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendUints(vs)
}

//Int64Ptr writes *int64
func (w *Writer) Int64Ptr(v *int64) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Int64(*v)
}

//Int64s writes []int64
func (w *Writer) Int64s(vs []int64) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendInt64s(vs)
}

//Uint64Ptr writes *uint64
func (w *Writer) Uint64Ptr(v *uint64) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Uint64(*v)
}

//Uint64s writes []uint64
func (w *Writer) Uint64s(vs []uint64) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendUint64s(vs)
}

//Int32Ptr writes *int32
func (w *Writer) Int32Ptr(v *int32) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Int32(*v)
}

//Int32s writes []int32
func (w *Writer) Int32s(vs []int32) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendInt32s(vs)
}

//Uint32Ptr writes *uint32
func (w *Writer) Uint32Ptr(v *uint32) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Uint32(*v)
}

//Uint32s writes []uint32
func (w *Writer) Uint32s(vs []uint32) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendUint32s(vs)
}

//Int16Ptr writes *int16
func (w *Writer) Int16Ptr(v *int16) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Int16(*v)
}

//Int16s writes []int16
func (w *Writer) Int16s(vs []int16) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendInt16s(vs)
}

//Uint16Ptr writes *uint16
func (w *Writer) Uint16Ptr(v *uint16) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Uint16(*v)
}

//Uint16s writes []uint16
func (w *Writer) Uint16s(vs []uint16) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendUint16s(vs)
}

//Int8Ptr writes *int8
func (w *Writer) Int8Ptr(v *int8) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Int8(*v)
}

//Int8s writes []int8
func (w *Writer) Int8s(vs []int8) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendInt8s(vs)
}

//Uint8Ptr writes *uint8
func (w *Writer) Uint8Ptr(v *uint8) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Uint8(*v)
}

//Uint8s writes []uint8
func (w *Writer) Uint8s(vs []uint8) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendUint8s(vs)
}

//Float64Ptr writes *float64
func (w *Writer) Float64Ptr(v *float64) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Float64(*v)
}

//Float64s writes []float64
func (w *Writer) Float64s(vs []float64) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendFloat64s(vs)
}

//Float32Ptr writes *float32
func (w *Writer) Float32Ptr(v *float32) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	w.Float32(*v)
}

//Float32s writes []float32
func (w *Writer) Float32s(vs []float32) {
	w.alloc.Uint32(uint32(len(vs)))
	w.appendFloat32s(vs)
}

//Bool writes bool
func (w *Writer) Bool(v bool) {
	i := uint8(0)
	if v {
		i = 1
	}
	w.Uint8(i)
}

//BoolPtr writes *bool
func (w *Writer) BoolPtr(v *bool) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	i := uint8(0)
	if *v {
		i = 1
	}
	w.Uint8(i)
}

//String writes string
func (w *Writer) String(v string) {
	b := unsafeGetBytes(v)
	w.Uint8s(b)
}

//StringPtr writes *string
func (w *Writer) StringPtr(v *string) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	b := unsafeGetBytes(*v)
	w.Uint8s(b)
}

//Strings writes []string
func (w *Writer) Strings(v []string) {
	size := len(v)
	w.alloc.Uint32(uint32(size))
	for i := range v {
		w.String(v[i])
	}
}

//Time writes time.Time
func (w *Writer) Time(v time.Time) {
	n := v.UnixNano()
	w.Int64(n)
}

//TimePtr writes *time.Time
func (w *Writer) TimePtr(v *time.Time) {
	if v == nil {
		w.alloc.Uint32(0)
		return
	}
	w.alloc.Uint32(1)
	n := v.UnixNano()
	w.Int64(n)
}

//Coder encodes data with encoder
func (w *Writer) Coder(v Encoder) error {
	if v == nil {
		w.alloc.Uint32(0)
		return nil
	}
	size := uint32(1)
	if allocator, ok := v.(Alloc); ok {
		size = allocator.Alloc()
	}
	w.alloc.Uint32(size)
	switch size {
	case 0:
		return nil
	case 1:
		return v.EncodeBinary(w)
	}
	for i := 0; i < int(size); i++ {
		if err := v.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) reset() {
	w.alloc = w.alloc[:0]
	w.encInts = w.encInts[:0]
	w.encUints = w.encUints[:0]
	w.encInt64s = w.encInt64s[:0]
	w.encUint64s = w.encUint64s[:0]
	w.encInt32s = w.encInt32s[:0]
	w.encUint32s = w.encUint32s[:0]
	w.encInt16s = w.encInt16s[:0]
	w.encUint16s = w.encUint16s[:0]
	w.encInt8s = w.encInt8s[:0]
	w.encUint8s = w.encUint8s[:0]
	w.encFloat64s = w.encFloat64s[:0]
	w.encFloat32s = w.encFloat32s[:0]
}

//Size returns data size
func (w *Writer) Size() int {
	result := size8bitsInBytes + w.alloc.size()
	result += w.encInts.size() + w.encUints.size()
	result += w.encInt64s.size() + w.encUint64s.size()
	result += w.encInt32s.size() + w.encUint32s.size()
	result += w.encInt16s.size() + w.encUint16s.size()
	result += w.encInt8s.size() + w.encUint8s.size()
	result += w.encFloat64s.size() + w.encFloat32s.size()
	return result
}

//Bytes returns writer bytes and resets stream
func (w *Writer) Bytes() []byte {
	var data = make([]byte, w.Size())
	offset := 0
	offset = w.alloc.store(data, offset, codecAlloc)
	offset = w.encInts.store(data, offset)
	offset = w.encFloat64s.store(data, offset)
	offset = w.encUint8s.store(data, offset)
	offset = w.encFloat32s.store(data, offset)
	offset = w.encUints.store(data, offset)
	offset = w.encInt64s.store(data, offset)
	offset = w.encUint64s.store(data, offset)
	offset = w.encInt32s.store(data, offset)
	offset = w.encUint32s.store(data, offset, codecUint32s)
	offset = w.encInt16s.store(data, offset)
	offset = w.encUint16s.store(data, offset)
	offset = w.encInt8s.store(data, offset)
	w.reset()
	data[offset] = codecEOF
	return data
}

//Int writes int
func (s *encInts) Int(v int) {
	*s = append(*s, v)
}

func (s *encInts) appendInts(vs []int) {
	*s = append(*s, vs...)
}

func (s *encInts) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecInts
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutInts(data[offset:], *s)
	offset += sizeIntInBytes * size
	return offset
}

func (s *encInts) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * sizeIntInBytes)
	}
	return 0
}

//Uint writes uint
func (s *encUints) Uint(v uint) {
	*s = append(*s, v)
}

func (s *encUints) appendUints(vs []uint) {
	*s = append(*s, vs...)
}

func (s *encUints) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecUints
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutUints(data[offset:], *s)
	offset += sizeIntInBytes * size
	return offset
}

func (s *encUints) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * sizeIntInBytes)
	}
	return 0
}

//Int64 writes int64
func (s *encInt64s) Int64(v int64) {
	*s = append(*s, v)
}

func (s *encInt64s) appendInt64s(vs []int64) {
	*s = append(*s, vs...)
}

func (s *encInt64s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecInt64s
	offset += size8bitsInBytes
	PutInt32(data[offset:], int32(size))
	offset += size32bitsInBytes
	PutInt64s(data[offset:], *s)
	offset += size64bitsInBytes * size
	return offset
}

func (s *encInt64s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size64bitsInBytes)
	}
	return 0
}

//Uint64 write uint64
func (s *encUint64s) Uint64(v uint64) {
	*s = append(*s, v)
}

func (s *encUint64s) appendUint64s(vs []uint64) {
	*s = append(*s, vs...)
}

func (s *encUint64s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecUint64s
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutUint64s(data[offset:], *s)
	offset += size64bitsInBytes * size
	return offset
}

func (s *encUint64s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size64bitsInBytes)
	}
	return 0
}

//Int32 writes int32
func (s *encInt32s) Int32(v int32) {
	*s = append(*s, v)
}

func (s *encInt32s) appendInt32s(v []int32) {
	*s = append(*s, v...)
}

func (s *encInt32s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecInt32s
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutInt32s(data[offset:], *s)
	offset += size32bitsInBytes * size
	return offset
}

func (s *encInt32s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size32bitsInBytes)
	}
	return 0
}

//Uint32 writes uint32
func (s *encUint32s) Uint32(v uint32) {
	*s = append(*s, v)
}

func (s *encUint32s) appendUint32s(v []uint32) {
	*s = append(*s, v...)
}

func (s *encUint32s) store(data []byte, offset int, codec uint8) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codec
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutUint32s(data[offset:], *s)
	offset += size32bitsInBytes * size
	return offset
}

func (s *encUint32s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size32bitsInBytes)
	}
	return 0
}

//Int16 writes int16
func (s *encInt16s) Int16(v int16) {
	*s = append(*s, v)
}

func (s *encInt16s) appendInt16s(v []int16) {
	*s = append(*s, v...)
}

func (s *encInt16s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecInt16s
	offset += size8bitsInBytes
	PutInt32(data[offset:], int32(size))
	offset += size32bitsInBytes
	PutInt16s(data[offset:], *s)
	offset += size16bitsInBytes * size
	return offset
}

func (s *encInt16s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size16bitsInBytes)
	}
	return 0
}

//Uint16 writes uint16
func (s *encUint16s) Uint16(v uint16) {
	*s = append(*s, v)
}

func (s *encUint16s) appendUint16s(v []uint16) {
	*s = append(*s, v...)
}

func (s *encUint16s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecUint16s
	offset += size8bitsInBytes
	PutInt32(data[offset:], int32(size))
	offset += size32bitsInBytes
	PutUint16s(data[offset:], *s)
	offset += size16bitsInBytes * size
	return offset
}

func (s *encUint16s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size16bitsInBytes + (size * size64bitsInBytes)
	}
	return 0
}

//Int8 writes int8
func (s *encInt8s) Int8(v int8) {
	*s = append(*s, v)
}

func (s *encInt8s) appendInt8s(v []int8) {
	*s = append(*s, v...)
}

func (s *encInt8s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecInt8s
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutInt8s(data[offset:], *s)
	offset += size8bitsInBytes * size
	return offset
}

func (s *encInt8s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size8bitsInBytes)
	}
	return 0
}

//Uint8 uint8
func (s *encUint8s) Uint8(v uint8) {
	*s = append(*s, v)
}

func (s *encUint8s) appendUint8s(v []uint8) {
	*s = append(*s, v...)
}

func (s *encUint8s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecUint8s
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutUint8s(data[offset:], *s)
	offset += size8bitsInBytes * size
	return offset
}

func (s *encUint8s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size8bitsInBytes)
	}
	return 0
}

//Float64 float64
func (s *encFloat64s) Float64(v float64) {
	*s = append(*s, v)
}

func (s *encFloat64s) appendFloat64s(v []float64) {
	*s = append(*s, v...)
}

func (s *encFloat64s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecFloat64s
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutFloat64s(data[offset:], *s)
	offset += size64bitsInBytes * size
	return offset
}

func (s *encFloat64s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size64bitsInBytes)
	}
	return 0
}

//Float32 writes float32
func (s *encFloat32s) Float32(v float32) {
	*s = append(*s, v)
}

func (s *encFloat32s) appendFloat32s(v []float32) {
	*s = append(*s, v...)
}

func (s *encFloat32s) store(data []byte, offset int) int {
	size := len(*s)
	if size == 0 {
		return offset
	}
	data[offset] = codecFloat32s
	offset += size8bitsInBytes
	PutUint32(data[offset:], uint32(size))
	offset += size32bitsInBytes
	PutFloat32s(data[offset:], *s)
	offset += size32bitsInBytes * size
	return offset
}

func (s *encFloat32s) size() int {
	if size := len(*s); size > 0 {
		return size8bitsInBytes + size32bitsInBytes + (size * size32bitsInBytes)
	}
	return 0
}

//Writers represents writer pool
type Writers struct {
	sync.Pool
}

//Get returns a writer
func (p *Writers) Get() *Writer {
	codec := p.Pool.Get()
	return codec.(*Writer)
}

//NewWriters creates writer pool
func NewWriters() *Writers {
	return &Writers{
		Pool: sync.Pool{
			New: func() interface{} {
				return &Writer{}
			},
		},
	}
}
