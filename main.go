package main

import (
	"fmt"
	"time"
)

func main() {


	//gmain:=g_main_loop_new(nil,1)
	//defer frida_deinit()
	manager:=NewFridaDeviceManager()
	defer manager.Close()
	d,err:=manager.GetDevice_with_id_milltimeout("8cd45fb4bbc44bcccfc10416ed3cd4b9e312dac6",1000)
	if err!=nil{
		panic(err)
	}
	a,err:=d.Get_Application_with_Identifier("com.meelive.ingkee")
	if err!=nil{
		panic(err)
	}

	session,err:=d.Attach(a.Pid)
	if err!=nil{
		panic(err)
	}
	sc,err:=session.Create_Script_with_path("./test.js")
	if err!=nil{
		panic(err)
	}
	sc.On("message", func(_script *FridaScript,_message map[string]interface{},_data []byte,_userdata uintptr) {
		fmt.Println(_message)
		fmt.Println(_data)
		//if jsobj["type"]=="log"{
		//	fmt.Println(jsobj["payload"])
		//}
	})
	err=sc.Load()
	if err!=nil{
		panic(err)
	}
	//fmt.Println(g_main_loop_is_running(gmain))


	//g_main_loop_run(gmain)
	time.Sleep(time.Second*4)






}