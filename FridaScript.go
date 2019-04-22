package gofrida

/*
#include "frida-core.h"
void on_message(FridaScript * script, gchar * message, GBytes * data, gpointer user_data);
 */
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"unsafe"
)

var calls=sync.Map{}

//export on_message
func  on_message(script *C.FridaScript,message *C.gchar,data *C.GBytes,user_data C.gpointer){
	defer g_bytes_unref(uintptr(unsafe.Pointer(data)))
	id:=frida_script_get_id(FridaScriptFormPtr(uintptr(unsafe.Pointer(script))))
	key:=fmt.Sprintf("%d_%s",id,"message")
	jsobj:=make(map[string]interface{})
	err:=json.Unmarshal([]byte(C.GoString(message)),&jsobj)
	if err!=nil{
		return
	}
	fv,ok:=calls.Load(key)
	if ok==false{
		return
	}
	var gobytes []byte
	if uintptr(unsafe.Pointer(data))!=uintptr(0){
		var nsize int
		pbuf:=g_bytes_get_data(uintptr(unsafe.Pointer(data)),&nsize)
		gobytes=C.GoBytes(unsafe.Pointer(pbuf),C.int(nsize))
	}

	f:=fv.(func(_script *FridaScript,_message map[string]interface{},_data []byte,_userdata uintptr))
	f(FridaScriptFormPtr(uintptr(unsafe.Pointer(script))),jsobj,gobytes,uintptr(unsafe.Pointer(user_data)))

}

type FridaScript struct{
	ptr uintptr
}
func FridaScriptFormPtr(_ptr uintptr)*FridaScript{
	return &FridaScript{_ptr}
}
func (this* FridaScript)CPtr()uintptr{
	return this.ptr
}

func (this* FridaScript)Post(_message map[string]interface{},_data uintptr)error{
	var gerr *GError
	bt,err:=json.Marshal(_message)
	if err!=nil{
		return err
	}
	frida_script_post_sync(this,string(bt),_data,&gerr)
	if gerr!=nil{
		return errors.New(gerr.Message())
	}
	return nil
}

func (this* FridaScript)Id()int{
	return frida_script_get_id(this)
}

func (this* FridaScript)On(_signal string,_func func(_script *FridaScript,_message map[string]interface{},_data []byte,_userdata uintptr)){
	key:=fmt.Sprintf("%d_%s",this.Id(),_signal)
	calls.Store(key,_func)
	g_signal_connect_data(this.CPtr(),_signal,C.GCallback(C.on_message),0,C.GClosureNotify(nil),0)
}

func (this* FridaScript)UnLoad()error{
	var gerr *GError
	frida_script_unload_sync(this,&gerr)
	if gerr!=nil{
		return errors.New(gerr.Message())
	}
	frida_unref(this.CPtr())
	return nil
}

func (this* FridaScript)Load()error{
	var gerr *GError
	frida_script_load_sync(this,&gerr)
	if gerr!=nil{
		return errors.New(gerr.Message())
	}
	return nil
}