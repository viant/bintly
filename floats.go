package bintly

import (
	"reflect"
	"unsafe"
)

//PutFloat64s copy []float64 into []byte
func PutFloat64s(bs []byte, vs []float64) {
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

//GetFloat64s copy []byte  into []float64
func GetFloat64s(bs []byte, vs []float64) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []float64
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}

//PutFloat32s copy []float32 into []byte
func PutFloat32s(bs []byte, vs []float32) {
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

//GetFloat32s copy []byte  into []float32
func GetFloat32s(bs []byte, vs []float32) {
	destLen := len(vs)
	if destLen == 0 {
		return
	}
	var data []float32
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(unsafe.Pointer(&bs[0]))
	sh.Len = destLen
	sh.Cap = destLen
	copy(vs, data)
}
