package gofrida

/*
 #include "frida-core.h"
*/
import "C"
import (
	"unsafe"
)

type GError struct{
	ptr uintptr
}
func (this* GError) Message()string{
	cerr:=(*C.GError)(unsafe.Pointer(this))
	serr:=C.GoString((*C.char)(cerr.message))
	return serr
}
func (this* GError)Code()int{
	cerr:=(*C.GError)(unsafe.Pointer(this))
	return int(cerr.code)
}
func GErrorFormPtr(_ptr uintptr)*GError{
	return &GError{_ptr}
}
func (this* GError)CPtr()uintptr{
	return this.ptr
}
