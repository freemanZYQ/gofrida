package main
/*
 #include "frida-core.h"
 */
import "C"
import (
	"reflect"
	"unsafe"
)

type GType=C.GType
type GQuark=C.GQuark
type JsonNodeType=C.JsonNodeType
type guint=C.guint
type gboolean=C.gboolean
type gint64=C.gint64
type gdouble=C.gdouble
type gconstpointer=C.gconstpointer
type double=C.double
type JsonObjectForeach=C.JsonObjectForeach
type gpointer=C.gpointer
type gchar=C.gchar
type JsonArrayForeach=C.JsonArrayForeach
type gunichar=C.gunichar
type gsize=C.gsize
type gssize=C.gssize
type GAsyncReadyCallback=C.GAsyncReadyCallback
type gint=C.gint
type JsonBoxedSerializeFunc=C.JsonBoxedSerializeFunc
type JsonBoxedDeserializeFunc=C.JsonBoxedDeserializeFunc
type FridaDeviceType=C.FridaDeviceType
type FridaDeviceManagerPredicate=C.FridaDeviceManagerPredicate
type FridaDeviceProcessPredicate=C.FridaDeviceProcessPredicate
type FridaStdio=C.FridaStdio
type FridaChildOrigin=C.FridaChildOrigin
type guint16=C.guint16
type GCallback=C.GCallback
type GClosureNotify=C.GClosureNotify
type GConnectFlags=C.GConnectFlags
type GDestroyNotify=C.GDestroyNotify
type wstring=string

type CObj interface{
	CPtr()uintptr
}

func GByte(_byte *[]byte)uintptr{
	if _byte==nil{
		return 0
	}
	return uintptr(unsafe.Pointer(&_byte))
}
func CObjPtr(_obj CObj)uintptr{
	if reflect.ValueOf(_obj).IsNil()==false{
		return _obj.CPtr()
	}
	return 0
}


type JsonObject struct{
	ptr uintptr
}
func JsonObjectFormPtr(_ptr uintptr)*JsonObject{
	return &JsonObject{_ptr}
}
func (this* JsonObject)CPtr()uintptr{
	return this.ptr
}

type JsonNode struct{
	ptr uintptr
}
func JsonNodeFormPtr(_ptr uintptr)*JsonNode{
	return &JsonNode{_ptr}
}
func (this* JsonNode)CPtr()uintptr{
	return this.ptr
}

type JsonArray struct{
	ptr uintptr
}
func JsonArrayFormPtr(_ptr uintptr)*JsonArray{
	return &JsonArray{_ptr}
}
func (this* JsonArray)CPtr()uintptr{
	return this.ptr
}

type JsonBuilder struct{
	ptr uintptr
}
func JsonBuilderFormPtr(_ptr uintptr)*JsonBuilder{
	return &JsonBuilder{_ptr}
}
func (this* JsonBuilder)CPtr()uintptr{
	return this.ptr
}


type JsonGenerator struct{
	ptr uintptr
}
func JsonGeneratorFormPtr(_ptr uintptr)*JsonGenerator{
	return &JsonGenerator{_ptr}
}
func (this* JsonGenerator)CPtr()uintptr{
	return this.ptr
}


type GList struct{
	ptr uintptr
}
func GListFormPtr(_ptr uintptr)*GList{
	return &GList{_ptr}
}
func (this* GList)CPtr()uintptr{
	return this.ptr
}

type JsonObjectIter struct{
	ptr uintptr
}
func JsonObjectIterFormPtr(_ptr uintptr)*JsonObjectIter{
	return &JsonObjectIter{_ptr}
}
func (this* JsonObjectIter)CPtr()uintptr{
	return this.ptr
}

type GOutputStream struct{
	ptr uintptr
}
func GOutputStreamFormPtr(_ptr uintptr)*GOutputStream{
	return &GOutputStream{_ptr}
}
func (this* GOutputStream)CPtr()uintptr{
	return this.ptr
}

type GCancellable struct{
	ptr uintptr
}
func GCancellableFormPtr(_ptr uintptr)*GCancellable{
	return &GCancellable{_ptr}
}
func (this* GCancellable)CPtr()uintptr{
	return this.ptr
}


type JsonParser struct{
	ptr uintptr
}
func JsonParserFormPtr(_ptr uintptr)*JsonParser{
	return &JsonParser{_ptr}
}
func (this* JsonParser)CPtr()uintptr{
	return this.ptr
}


type GInputStream struct{
	ptr uintptr
}
func GInputStreamFormPtr(_ptr uintptr)*GInputStream{
	return &GInputStream{_ptr}
}
func (this* GInputStream)CPtr()uintptr{
	return this.ptr
}

type GAsyncResult struct{
	ptr uintptr
}
func GAsyncResultFormPtr(_ptr uintptr)*GAsyncResult{
	return &GAsyncResult{_ptr}
}
func (this* GAsyncResult)CPtr()uintptr{
	return this.ptr
}

type JsonPath struct{
	ptr uintptr
}
func JsonPathFormPtr(_ptr uintptr)*JsonPath{
	return &JsonPath{_ptr}
}
func (this* JsonPath)CPtr()uintptr{
	return this.ptr
}

type JsonReader struct{
	ptr uintptr
}
func JsonReaderFormPtr(_ptr uintptr)*JsonReader{
	return &JsonReader{_ptr}
}
func (this* JsonReader)CPtr()uintptr{
	return this.ptr
}

type JsonSerializable struct{
	ptr uintptr
}
func JsonSerializableFormPtr(_ptr uintptr)*JsonSerializable{
	return &JsonSerializable{_ptr}
}
func (this* JsonSerializable)CPtr()uintptr{
	return this.ptr
}

type GParamSpec struct{
	ptr uintptr
}
func GParamSpecFormPtr(_ptr uintptr)*GParamSpec{
	return &GParamSpec{_ptr}
}
func (this* GParamSpec)CPtr()uintptr{
	return this.ptr
}

type GObject struct{
	ptr uintptr
}
func GObjectFormPtr(_ptr uintptr)*GObject{
	return &GObject{_ptr}
}
func (this* GObject)CPtr()uintptr{
	return this.ptr
}


type GVariant struct{
	ptr uintptr
}
func GVariantFormPtr(_ptr uintptr)*GVariant{
	return &GVariant{_ptr}
}
func (this* GVariant)CPtr()uintptr{
	return this.ptr
}

type GMainContext struct{
	ptr uintptr
}
func GMainContextFormPtr(_ptr uintptr)*GMainContext{
	return &GMainContext{_ptr}
}
func (this* GMainContext)CPtr()uintptr{
	return this.ptr
}



type FridaDeviceList struct{
	ptr uintptr
}
func FridaDeviceListFormPtr(_ptr uintptr)*FridaDeviceList{
	return &FridaDeviceList{_ptr}
}
func (this* FridaDeviceList)CPtr()uintptr{
	return this.ptr
}


type FridaIcon struct{
	ptr uintptr
}
func FridaIconFormPtr(_ptr uintptr)*FridaIcon{
	return &FridaIcon{_ptr}
}
func (this* FridaIcon)CPtr()uintptr{
	return this.ptr
}



type FridaApplicationList struct{
	ptr uintptr
}
func FridaApplicationListFormPtr(_ptr uintptr)*FridaApplicationList{
	return &FridaApplicationList{_ptr}
}
func (this* FridaApplicationList)CPtr()uintptr{
	return this.ptr
}



type GBytes struct{
	ptr uintptr
}
func GBytesFormPtr(_ptr uintptr)*GBytes{
	return &GBytes{_ptr}
}
func (this* GBytes)CPtr()uintptr{
	return this.ptr
}



type FridaProcessList struct{
	ptr uintptr
}
func FridaProcessListFormPtr(_ptr uintptr)*FridaProcessList{
	return &FridaProcessList{_ptr}
}
func (this* FridaProcessList)CPtr()uintptr{
	return this.ptr
}


type FridaSpawnOptions struct{
	ptr uintptr
}
func FridaSpawnOptionsFormPtr(_ptr uintptr)*FridaSpawnOptions{
	return &FridaSpawnOptions{_ptr}
}
func (this* FridaSpawnOptions)CPtr()uintptr{
	return this.ptr
}

type FridaSpawnList struct{
	ptr uintptr
}
func FridaSpawnListFormPtr(_ptr uintptr)*FridaSpawnList{
	return &FridaSpawnList{_ptr}
}
func (this* FridaSpawnList)CPtr()uintptr{
	return this.ptr
}

type FridaSpawn struct{
	ptr uintptr
}
func FridaSpawnFormPtr(_ptr uintptr)*FridaSpawn{
	return &FridaSpawn{_ptr}
}
func (this* FridaSpawn)CPtr()uintptr{
	return this.ptr
}

type GVariantDict struct{
	ptr uintptr
}
func GVariantDictFormPtr(_ptr uintptr)*GVariantDict{
	return &GVariantDict{_ptr}
}
func (this* GVariantDict)CPtr()uintptr{
	return this.ptr
}

type FridaChildList struct{
	ptr uintptr
}
func FridaChildListFormPtr(_ptr uintptr)*FridaChildList{
	return &FridaChildList{_ptr}
}
func (this* FridaChildList)CPtr()uintptr{
	return this.ptr
}


type FridaChild struct{
	ptr uintptr
}
func FridaChildFormPtr(_ptr uintptr)*FridaChild{
	return &FridaChild{_ptr}
}
func (this* FridaChild)CPtr()uintptr{
	return this.ptr
}



type FridaCrash struct{
	ptr uintptr
}
func FridaCrashFormPtr(_ptr uintptr)*FridaCrash{
	return &FridaCrash{_ptr}
}
func (this* FridaCrash)CPtr()uintptr{
	return this.ptr
}

type FridaInjector struct{
	ptr uintptr
}
func FridaInjectorFormPtr(_ptr uintptr)*FridaInjector{
	return &FridaInjector{_ptr}
}
func (this* FridaInjector)CPtr()uintptr{
	return this.ptr
}

type GValue struct{
	ptr uintptr
}
func GValueFormPtr(_ptr uintptr)*GValue{
	return &GValue{_ptr}
}
func (this* GValue)CPtr()uintptr{
	return this.ptr
}

type GMainLoop struct{
	ptr uintptr
}
func GMainLoopFormPtr(_ptr uintptr)*GMainLoop{
	return &GMainLoop{_ptr}
}
func (this* GMainLoop)CPtr()uintptr{
	return this.ptr
}

type FridaFileMonitor struct{
	ptr uintptr
}
func FridaFileMonitorFormPtr(_ptr uintptr)*FridaFileMonitor{
	return &FridaFileMonitor{_ptr}
}
func (this* FridaFileMonitor)CPtr()uintptr{
	return this.ptr
}

type GIcon struct{
	ptr uintptr
}
func GIconFormPtr(_ptr uintptr)*GIcon{
	return &GIcon{_ptr}
}
func (this* GIcon)CPtr()uintptr{
	return this.ptr
}

type GBytesIcon struct{
	ptr uintptr
}
func GBytesIconFormPtr(_ptr uintptr)*GBytesIcon{
	return &GBytesIcon{_ptr}
}
func (this* GBytesIcon)CPtr()uintptr{
	return this.ptr
}

type GByteArray struct{
	ptr uintptr
}
func GByteArrayFormPtr(_ptr uintptr)*GByteArray{
	return &GByteArray{_ptr}
}
func (this* GByteArray)CPtr()uintptr{
	return this.ptr
}