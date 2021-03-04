package bintly

import (
	"reflect"
	"unsafe"
)

const (
	n = 2048
)


//PutInts copy []int into []byte
func PutInts(bs []byte, vs []int) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//Ints copy []byte  into []int
func Ints(bs []byte) []int {
	vs := make([]int, len(bs)/sizeIntInBytes)
	GetInts(bs, vs)
	return vs
}

//GetInts copy []byte  into []int
func GetInts(bs []byte, vs []int) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []int
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//PutUints copy []uint into []byte
func PutUints(bs []byte, vs []uint) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetUints copy []byte  into []uint
func GetUints(bs []byte, vs []uint) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []uint
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//PutUint64s copy []uint64 into []byte
func PutUint64s(bs []byte, vs []uint64) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetUint64s copy []byte  into []uint64
func GetUint64s(bs []byte, vs []uint64) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []uint64
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//Uint64s copy []byte  into []uint64
func Uint64s(bs []byte) []uint64 {
	vs := make([]uint64, len(bs)/size64bitsInBytes)
	GetUint64s(bs, vs)
	return vs
}

//PutUint32s copy []uint32 into []byte
func PutUint32s(bs []byte, vs []uint32) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetUint32s copy []byte  into []uint32
func GetUint32s(bs []byte, vs []uint32) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []uint32
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//Uint32s copy []byte  into []uint32
func Uint32s(bs []byte) []uint32 {
	vs := make([]uint32, len(bs)/size32bitsInBytes)
	GetUint32s(bs, vs)
	return vs
}

//PutUint16s copy []uint16 into []byte
func PutUint16s(bs []byte, vs []uint16) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetUint16s copy []byte  into []uint16
func GetUint16s(bs []byte, vs []uint16) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []uint16
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//Uint16s copy []byte  into []uint16
func Uint16s(bs []byte) []uint16 {
	vs := make([]uint16, len(bs)/size16bitsInBytes)
	GetUint16s(bs, vs)
	return vs
}

//PutInt64s copy []int64 into []byte
func PutInt64s(bs []byte, vs []int64) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetInt64s copy []byte  into []int64
func GetInt64s(bs []byte, vs []int64) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []int64
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//Int64s copy []byte  into []int64
func Int64s(bs []byte) []int64 {
	vs := make([]int64, len(bs)/size64bitsInBytes)
	GetInt64s(bs, vs)
	return vs
}

//PutInt32s copy []int32 into []byte
func PutInt32s(bs []byte, vs []int32) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetInt32s copy []byte  into []int32
func GetInt32s(bs []byte, vs []int32) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []int32
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//Int32s copy []byte  into []int32
func Int32s(bs []byte) []int32 {
	vs := make([]int32, len(bs)/size32bitsInBytes)
	GetInt32s(bs, vs)
	return vs
}

//PutInt16s copy []int16 into []byte
func PutInt16s(bs []byte, vs []int16) {
	destLen := len(bs)
	if destLen == 0 {
		return
	}
	var data []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&vs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(bs, data)
}

//GetInt16s copy []byte  into []int16
func GetInt16s(bs []byte, vs []int16) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []int16
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//Int16s copy []byte  into []int16
func Int16s(bs []byte) []int16 {
	vs := make([]int16, len(bs)/size16bitsInBytes)
	GetInt16s(bs, vs)
	return vs
}

//PutInt8s copy []int8 into []byte
func PutInt8s(bs []byte, vs []int8) {
	copy(bs, *(*[]byte)(unsafe.Pointer(&vs)))
}

//GetInt8s copy []byte  into []int8
func GetInt8s(bs []byte, vs []int8) {
	copy(vs, *(*[]int8)(unsafe.Pointer(&bs)))
}

//Int8s copy []byte  into []int8
func Int8s(bs []byte) []int8 {
	vs := make([]int8, len(bs))
	GetInt8s(bs, vs)
	return vs
}

//PutUint8s copy []uint8 into []byte
func PutUint8s(bs []byte, vs []uint8) {
	copy(bs, vs)
}

//GetUint8s copy []byte  into []uint8
func GetUint8s(bs []byte, vs []uint8) {
	copy(vs, bs)
}

//Uint8s copy []byte  into []uint8
func Uint8s(bs []byte) []uint8 {
	vs := make([]uint8, len(bs))
	GetUint8s(bs, vs)
	return vs
}
