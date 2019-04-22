package main


type FridaProcess struct{
	Name string
	Pid int
	ptr uintptr
}
func FridaProcessFormPtr(_ptr uintptr)*FridaProcess{
	return &FridaProcess{ptr:_ptr}
}
func (this* FridaProcess)CPtr()uintptr{
	return this.ptr
}

func (this* FridaProcess)Fill(){
	this.Name=frida_process_get_name(this)
	this.Pid=frida_process_get_pid(this)
}