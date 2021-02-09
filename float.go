package bintly

import "unsafe"

//PutFloat64 copy float64 into []byte
func PutFloat64(dest []byte, v float64) {
	copy(dest, (*(*[size64bits]byte)(unsafe.Pointer(&v)))[:])
}

//Float64 copy []byte  to float64
func Float64(b []byte) float64 {
	v := float64(0)
	copy((*(*[size64bits]byte)(unsafe.Pointer(&v)))[:], b)
	return v
}

//PutFloat32 copy float32 into []byte
func PutFloat32(dest []byte, v float32) {
	copy(dest, (*(*[size32bits]byte)(unsafe.Pointer(&v)))[:])
}

//Float32 copy []byte to float32
func Float32(b []byte) float32 {
	v := float32(0)
	copy((*(*[size32bits]byte)(unsafe.Pointer(&v)))[:], b)
	return v
}
