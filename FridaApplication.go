package main

type FridaApplication struct{
	Identifier string
	Name string
	Pid int
	ptr uintptr
}
func FridaApplicationFormPtr(_ptr uintptr)*FridaApplication{
	return &FridaApplication{ptr:_ptr}
}
func (this* FridaApplication)CPtr()uintptr{
	return this.ptr
}

func (this* FridaApplication)Fill(){
	this.Identifier=frida_application_get_identifier(this)
	this.Name=frida_application_get_name(this)
	this.Pid=frida_application_get_pid(this)
}
