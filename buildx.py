import codecs

from clang import *
from clang.cindex import *
import sys
Config.set_library_file("C:\\Program Files\\LLVM\\bin\\libclang.dll")
hname="frida-core.h"
hcall="cldef.go"
#hstruct="type1.go"
dep=[
    "__va_start",
    "_GLIB_",
    "json_",
    "g_",
    "gio_",
    "glib_",
    "G_",
    "gobject_",
    "_g_",
    "atexit",
]
funcs=[
    "g_bytes",
    "frida_init",
    "frida_device_manager_new",
    "frida_device_manager_close",
    "frida_shutdown",
    "frida_deinit",
    "g_main_loop_new",
    "frida_device_manager_enumerate_devices_sync",
    "frida_device_list_size",
    "frida_device_list_get",
    "frida_device_get_name",
    "frida_device_get_dtype",
    "frida_device_attach_sync",
    "frida_session_create_script_sync",
    "frida_script_load_sync",
    "g_main_loop_is_running",
    "g_signal_connect_data",
    "g_main_loop_run",
    "frida_script_unload_sync",
    "frida_session_detach_sync",
    "frida_device_manager_close_sync",
    "frida_unref",
    "frida_script_post_sync",
]

depstruct=[
    "Json",
]
def indep(_funcname:str):
    for it in dep:
        if _funcname.startswith(it)==True:
            return True
    return False
def infuncs(_funcname:str):
    for it in funcs:
        if _funcname.startswith(it)==True:
            return True
    return False
def instructdep(_funcname:str):
    for it in depstruct:
        if _funcname.startswith(it)==True:
            return True
    return False


def main():
    index = Index.create()
    root = index.parse(hname)
    f = codecs.open(hcall, "w")
    #f1 = codecs.open(hstruct, "w")
    genCallHead(root.cursor,f)
    genCallloaders(root.cursor,f)
    genCalls(root.cursor,f)
    # genStruct(root.cursor,f1)
    f.close()
    print("生成完成")

def genStruct(node:Cursor,_f):
    for it in node.get_children():  # type:Cursor
        if (it.kind==CursorKind.TYPEDEF_DECL  and (instructdep(it.spelling)==False) and str(it.location.file)==hname):
            kd=it.kind#type:CursorKind
            sys.stdout.write("{}\n".format(it.is_scoped_enum()))

def genCalls(node:Cursor,_f):
    for it in node.get_children():#type:Cursor
        #or str(it.spelling) in funcs
        if infuncs(it.spelling)==True or (it.kind==CursorKind.FUNCTION_DECL and (indep(it.spelling)==False)  and str(it.location.file)==hname):
            print("正在生成函数:{} {}".format(it.spelling,it.type.spelling))
            _f.write("\n")
            genCall(it,_f)
        genCalls(it,_f)
def genCallloaders(node:Cursor,_f):
    for it in node.get_children():#type:Cursor
        #and it.spelling in funcs
        if infuncs(it.spelling)==True or ((it.kind==CursorKind.FUNCTION_DECL ) and (indep(it.spelling)==False) and str(it.location.file)==hname):
            print("正在生成函数:{} {}".format(it.spelling,it.type.spelling))
            genCallloader(it,_f)
        genCallloaders(it,_f)
def genCall(node:Cursor,_f):
    funcname=node.spelling
    functype=build_functype(node.result_type.spelling)
    args=[{"k":it.spelling,"v":it.type.spelling} for it in node.get_arguments()]
    sargs=build_funcargs(args)
    callargs=build_callargs(args)
    ret=build_retval(functype)
    retfunc=build_ret(functype)
    _f.write("func {}({}){}{{\n".format(funcname,sargs,functype))
    _f.write("    {}CL_{}.Call({})\n".format(ret,funcname,callargs))
    _f.write("    {}\n".format(retfunc))
    _f.write("}\n")
def build_retval(_type:str):
    result="r,_,_:="
    if _type=="":
        result=""
    return result
def build_ret(_type:str):
    result=""
    if _type=="":
        result=""
    elif _type=="int":
        result="return int(r)"
    elif _type=="** gchar":
        result = "return uintptr(unsafe.Pointer(r))"
    elif _type=="string":
        result="return C.GoString((*C.char)(unsafe.Pointer(r)))"
    elif _type.startswith("* "):
        result="return {}FormPtr(r)".format(_type.replace("* ",""))
    elif _type=="*[]byte":
        result="return r"
    else:
        result="return {}(r)".format(_type)
    return result
def convert_type(_type:str):
    result=_type.replace("const ","")
    result=result.replace("__LineNo ","")
    if result=="void":
        result=""
    elif result=="gint":
        result="int"
    elif result=="guint":
        result = "int"
    elif result=="unsigned long":
        result="int"
    elif result=="gboolean":
        result="int"
    elif result=="gconstpointer":
        result="uintptr"
    elif result=="int":
        result = "uintptr"
    elif result=="gsize *":
        result = "*int"
    elif result=="gpointer":
        result="uintptr"
    elif result=="volatile void *":
        result="uintptr"
    elif result=="gulong":
        result="int"
    elif result=="unsigned int":
        result="int"
    elif result=="guint16":
        result = "int"
    elif result=="gchar **":
        result="uintptr"
    elif result=="uintptr_t":
        result="uintptr"
    elif result=="void (*)(void)":
        result = "uintptr"
    elif result=="GBytes *":
        return "uintptr"
    elif result=="wchar_t *":
        return "wstring"
    elif result=="gchar *":
        return "string"
    elif result=="guint *":
        return "*int"
    elif result=="gint *":
        return "*int"
    elif result=="const gchar *":
        result="string"
    elif result.endswith(" *"):
        result="* {}".format(result.replace(" *",""))
    elif result.endswith(" **"):
        result = "** {}".format(result.replace(" **", ""))
    return result
def build_functype(_functype:str):
    return convert_type(_functype)
def build_funcargs(_args:list):
    sargs=""
    for k in _args:
        argname="_{}".format(k["k"])
        argval=convert_type(k["v"])
        sargs="{} {} {},".format(sargs,argname,argval)

    return sargs
def build_callargs(_args:list):
    sargs=""
    for k in _args:
        argname="_{}".format(k["k"])
        argval = convert_type(k["v"])
        val=build_callarg(argname,argval)
        sargs="{}{},".format(sargs,val)
    return sargs
def build_callarg(_argname,_argtype):
    result=_argname
    if _argtype=="int":
        result="uintptr(C.int({}))".format(_argname)
    elif _argtype=="string":
        result= "uintptr(unsafe.Pointer(syscall.StringBytePtr({})))".format(_argname)
    elif _argtype=="wstring":
        result= "uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr({})))".format(_argname)
    elif _argtype.startswith("* "):
        result="CObjPtr({})".format(_argname)
    elif _argtype=="GAsyncReadyCallback":
        result="uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype=="GCallback":
        result = "uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype=="GClosureNotify":
        result = "uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype == "GDestroyNotify":
        result = "uintptr(unsafe.Pointer({}))".format(_argname)

    elif _argtype=="FridaDeviceManagerPredicate":
        result = "uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype=="FridaDeviceProcessPredicate":
        result = "uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype.startswith("** "):
        result = "uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype=="*[]byte":
        result ="uintptr(unsafe.Pointer({}))".format(_argname)
    elif _argtype=="*int":
        result ="uintptr(unsafe.Pointer({}))".format(_argname)
    else:
        result="uintptr(int({}))".format(_argname)
        #result="uintptr(unsafe.Pointer(C.{}({})))".format(_argtype,_argname)
    return result

def genCallloader(node:Cursor,_f):
    funcname=node.spelling
    functype=node.result_type.spelling
    args=[{it.spelling:it.type.spelling} for it in node.get_arguments()]
    _f.writelines("var CL_{}=CL_Dll.NewProc(\"{}\")\n".format(funcname, funcname))



def genCallHead(node,_f):
    _f.writelines("package main\n")
    _f.writelines("\n\n")
    _f.write("import \"C\"\n")
    _f.writelines("import (\n")
    _f.writelines("    \"syscall\"\n")
    _f.writelines("    \"unsafe\"\n")
    _f.writelines(")")
    _f.writelines("\n\n")
    _f.writelines("var CL_Dll *syscall.LazyDLL = loaddll()\n")

if __name__ == '__main__':
    main()
