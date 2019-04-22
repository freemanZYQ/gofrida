package gofrida

import "C"
import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)


func loaddll() *syscall.LazyDLL {
	libname := "frida.dll"
	dll := syscall.NewLazyDLL(libname)
	if err := dll.Load(); err != nil {
		fmt.Println(fmt.Sprintf("加载dll失败,frida.dll和vcruntime140.dll是不是存在,详细错误:%s",err.Error()))
		os.Exit(-1)
		return nil
	}
	return dll
}

func gostr(_p *C.char) string {
	return C.GoString(_p)
}
func cstr(_s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(_s)))
}
func goint(val uintptr) int {
	return int(val)
}
func gobool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}