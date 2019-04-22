package main


import "C"
import (
    "syscall"
    "unsafe"
)

var CL_Dll *syscall.LazyDLL = loaddll()
var CL_g_bytes_new=CL_Dll.NewProc("g_bytes_new")
var CL_g_bytes_new_take=CL_Dll.NewProc("g_bytes_new_take")
var CL_g_bytes_new_static=CL_Dll.NewProc("g_bytes_new_static")
var CL_g_bytes_new_with_free_func=CL_Dll.NewProc("g_bytes_new_with_free_func")
var CL_g_bytes_new_from_bytes=CL_Dll.NewProc("g_bytes_new_from_bytes")
var CL_g_bytes_get_data=CL_Dll.NewProc("g_bytes_get_data")
var CL_g_bytes_get_size=CL_Dll.NewProc("g_bytes_get_size")
var CL_g_bytes_ref=CL_Dll.NewProc("g_bytes_ref")
var CL_g_bytes_unref=CL_Dll.NewProc("g_bytes_unref")
var CL_g_bytes_unref_to_data=CL_Dll.NewProc("g_bytes_unref_to_data")
var CL_g_bytes_unref_to_array=CL_Dll.NewProc("g_bytes_unref_to_array")
var CL_g_bytes_hash=CL_Dll.NewProc("g_bytes_hash")
var CL_g_bytes_equal=CL_Dll.NewProc("g_bytes_equal")
var CL_g_bytes_compare=CL_Dll.NewProc("g_bytes_compare")
var CL_g_main_loop_new=CL_Dll.NewProc("g_main_loop_new")
var CL_g_main_loop_run=CL_Dll.NewProc("g_main_loop_run")
var CL_g_main_loop_is_running=CL_Dll.NewProc("g_main_loop_is_running")
var CL_g_signal_connect_data=CL_Dll.NewProc("g_signal_connect_data")
var CL_g_bytes_get_type=CL_Dll.NewProc("g_bytes_get_type")
var CL_g_bytes_icon_get_type=CL_Dll.NewProc("g_bytes_icon_get_type")
var CL_g_bytes_icon_new=CL_Dll.NewProc("g_bytes_icon_new")
var CL_g_bytes_icon_get_bytes=CL_Dll.NewProc("g_bytes_icon_get_bytes")
var CL_frida_init=CL_Dll.NewProc("frida_init")
var CL_frida_shutdown=CL_Dll.NewProc("frida_shutdown")
var CL_frida_deinit=CL_Dll.NewProc("frida_deinit")
var CL_frida_get_main_context=CL_Dll.NewProc("frida_get_main_context")
var CL_frida_unref=CL_Dll.NewProc("frida_unref")
var CL_frida_version=CL_Dll.NewProc("frida_version")
var CL_frida_version_string=CL_Dll.NewProc("frida_version_string")
var CL_frida_device_manager_new=CL_Dll.NewProc("frida_device_manager_new")
var CL_frida_device_manager_close=CL_Dll.NewProc("frida_device_manager_close")
var CL_frida_device_manager_close_finish=CL_Dll.NewProc("frida_device_manager_close_finish")
var CL_frida_device_manager_close_sync=CL_Dll.NewProc("frida_device_manager_close_sync")
var CL_frida_device_manager_get_device_by_id=CL_Dll.NewProc("frida_device_manager_get_device_by_id")
var CL_frida_device_manager_get_device_by_id_finish=CL_Dll.NewProc("frida_device_manager_get_device_by_id_finish")
var CL_frida_device_manager_get_device_by_id_sync=CL_Dll.NewProc("frida_device_manager_get_device_by_id_sync")
var CL_frida_device_manager_get_device_by_type=CL_Dll.NewProc("frida_device_manager_get_device_by_type")
var CL_frida_device_manager_get_device_by_type_finish=CL_Dll.NewProc("frida_device_manager_get_device_by_type_finish")
var CL_frida_device_manager_get_device_by_type_sync=CL_Dll.NewProc("frida_device_manager_get_device_by_type_sync")
var CL_frida_device_manager_get_device=CL_Dll.NewProc("frida_device_manager_get_device")
var CL_frida_device_manager_get_device_finish=CL_Dll.NewProc("frida_device_manager_get_device_finish")
var CL_frida_device_manager_get_device_sync=CL_Dll.NewProc("frida_device_manager_get_device_sync")
var CL_frida_device_manager_find_device_by_id=CL_Dll.NewProc("frida_device_manager_find_device_by_id")
var CL_frida_device_manager_find_device_by_id_finish=CL_Dll.NewProc("frida_device_manager_find_device_by_id_finish")
var CL_frida_device_manager_find_device_by_id_sync=CL_Dll.NewProc("frida_device_manager_find_device_by_id_sync")
var CL_frida_device_manager_find_device_by_type=CL_Dll.NewProc("frida_device_manager_find_device_by_type")
var CL_frida_device_manager_find_device_by_type_finish=CL_Dll.NewProc("frida_device_manager_find_device_by_type_finish")
var CL_frida_device_manager_find_device_by_type_sync=CL_Dll.NewProc("frida_device_manager_find_device_by_type_sync")
var CL_frida_device_manager_find_device=CL_Dll.NewProc("frida_device_manager_find_device")
var CL_frida_device_manager_find_device_finish=CL_Dll.NewProc("frida_device_manager_find_device_finish")
var CL_frida_device_manager_find_device_sync=CL_Dll.NewProc("frida_device_manager_find_device_sync")
var CL_frida_device_manager_enumerate_devices=CL_Dll.NewProc("frida_device_manager_enumerate_devices")
var CL_frida_device_manager_enumerate_devices_finish=CL_Dll.NewProc("frida_device_manager_enumerate_devices_finish")
var CL_frida_device_manager_enumerate_devices_sync=CL_Dll.NewProc("frida_device_manager_enumerate_devices_sync")
var CL_frida_device_manager_add_remote_device=CL_Dll.NewProc("frida_device_manager_add_remote_device")
var CL_frida_device_manager_add_remote_device_finish=CL_Dll.NewProc("frida_device_manager_add_remote_device_finish")
var CL_frida_device_manager_add_remote_device_sync=CL_Dll.NewProc("frida_device_manager_add_remote_device_sync")
var CL_frida_device_manager_remove_remote_device=CL_Dll.NewProc("frida_device_manager_remove_remote_device")
var CL_frida_device_manager_remove_remote_device_finish=CL_Dll.NewProc("frida_device_manager_remove_remote_device_finish")
var CL_frida_device_manager_remove_remote_device_sync=CL_Dll.NewProc("frida_device_manager_remove_remote_device_sync")
var CL_frida_device_list_size=CL_Dll.NewProc("frida_device_list_size")
var CL_frida_device_list_get=CL_Dll.NewProc("frida_device_list_get")
var CL_frida_device_get_id=CL_Dll.NewProc("frida_device_get_id")
var CL_frida_device_get_name=CL_Dll.NewProc("frida_device_get_name")
var CL_frida_device_get_icon=CL_Dll.NewProc("frida_device_get_icon")
var CL_frida_device_get_dtype=CL_Dll.NewProc("frida_device_get_dtype")
var CL_frida_device_get_manager=CL_Dll.NewProc("frida_device_get_manager")
var CL_frida_device_is_lost=CL_Dll.NewProc("frida_device_is_lost")
var CL_frida_device_get_frontmost_application=CL_Dll.NewProc("frida_device_get_frontmost_application")
var CL_frida_device_get_frontmost_application_finish=CL_Dll.NewProc("frida_device_get_frontmost_application_finish")
var CL_frida_device_get_frontmost_application_sync=CL_Dll.NewProc("frida_device_get_frontmost_application_sync")
var CL_frida_device_enumerate_applications=CL_Dll.NewProc("frida_device_enumerate_applications")
var CL_frida_device_enumerate_applications_finish=CL_Dll.NewProc("frida_device_enumerate_applications_finish")
var CL_frida_device_enumerate_applications_sync=CL_Dll.NewProc("frida_device_enumerate_applications_sync")
var CL_frida_device_get_process_by_pid=CL_Dll.NewProc("frida_device_get_process_by_pid")
var CL_frida_device_get_process_by_pid_finish=CL_Dll.NewProc("frida_device_get_process_by_pid_finish")
var CL_frida_device_get_process_by_pid_sync=CL_Dll.NewProc("frida_device_get_process_by_pid_sync")
var CL_frida_device_get_process_by_name=CL_Dll.NewProc("frida_device_get_process_by_name")
var CL_frida_device_get_process_by_name_finish=CL_Dll.NewProc("frida_device_get_process_by_name_finish")
var CL_frida_device_get_process_by_name_sync=CL_Dll.NewProc("frida_device_get_process_by_name_sync")
var CL_frida_device_get_process=CL_Dll.NewProc("frida_device_get_process")
var CL_frida_device_get_process_finish=CL_Dll.NewProc("frida_device_get_process_finish")
var CL_frida_device_get_process_sync=CL_Dll.NewProc("frida_device_get_process_sync")
var CL_frida_device_find_process_by_pid=CL_Dll.NewProc("frida_device_find_process_by_pid")
var CL_frida_device_find_process_by_pid_finish=CL_Dll.NewProc("frida_device_find_process_by_pid_finish")
var CL_frida_device_find_process_by_pid_sync=CL_Dll.NewProc("frida_device_find_process_by_pid_sync")
var CL_frida_device_find_process_by_name=CL_Dll.NewProc("frida_device_find_process_by_name")
var CL_frida_device_find_process_by_name_finish=CL_Dll.NewProc("frida_device_find_process_by_name_finish")
var CL_frida_device_find_process_by_name_sync=CL_Dll.NewProc("frida_device_find_process_by_name_sync")
var CL_frida_device_find_process=CL_Dll.NewProc("frida_device_find_process")
var CL_frida_device_find_process_finish=CL_Dll.NewProc("frida_device_find_process_finish")
var CL_frida_device_find_process_sync=CL_Dll.NewProc("frida_device_find_process_sync")
var CL_frida_device_enumerate_processes=CL_Dll.NewProc("frida_device_enumerate_processes")
var CL_frida_device_enumerate_processes_finish=CL_Dll.NewProc("frida_device_enumerate_processes_finish")
var CL_frida_device_enumerate_processes_sync=CL_Dll.NewProc("frida_device_enumerate_processes_sync")
var CL_frida_device_enable_spawn_gating=CL_Dll.NewProc("frida_device_enable_spawn_gating")
var CL_frida_device_enable_spawn_gating_finish=CL_Dll.NewProc("frida_device_enable_spawn_gating_finish")
var CL_frida_device_enable_spawn_gating_sync=CL_Dll.NewProc("frida_device_enable_spawn_gating_sync")
var CL_frida_device_disable_spawn_gating=CL_Dll.NewProc("frida_device_disable_spawn_gating")
var CL_frida_device_disable_spawn_gating_finish=CL_Dll.NewProc("frida_device_disable_spawn_gating_finish")
var CL_frida_device_disable_spawn_gating_sync=CL_Dll.NewProc("frida_device_disable_spawn_gating_sync")
var CL_frida_device_enumerate_pending_spawn=CL_Dll.NewProc("frida_device_enumerate_pending_spawn")
var CL_frida_device_enumerate_pending_spawn_finish=CL_Dll.NewProc("frida_device_enumerate_pending_spawn_finish")
var CL_frida_device_enumerate_pending_spawn_sync=CL_Dll.NewProc("frida_device_enumerate_pending_spawn_sync")
var CL_frida_device_enumerate_pending_children=CL_Dll.NewProc("frida_device_enumerate_pending_children")
var CL_frida_device_enumerate_pending_children_finish=CL_Dll.NewProc("frida_device_enumerate_pending_children_finish")
var CL_frida_device_enumerate_pending_children_sync=CL_Dll.NewProc("frida_device_enumerate_pending_children_sync")
var CL_frida_device_spawn=CL_Dll.NewProc("frida_device_spawn")
var CL_frida_device_spawn_finish=CL_Dll.NewProc("frida_device_spawn_finish")
var CL_frida_device_spawn_sync=CL_Dll.NewProc("frida_device_spawn_sync")
var CL_frida_device_input=CL_Dll.NewProc("frida_device_input")
var CL_frida_device_input_finish=CL_Dll.NewProc("frida_device_input_finish")
var CL_frida_device_input_sync=CL_Dll.NewProc("frida_device_input_sync")
var CL_frida_device_resume=CL_Dll.NewProc("frida_device_resume")
var CL_frida_device_resume_finish=CL_Dll.NewProc("frida_device_resume_finish")
var CL_frida_device_resume_sync=CL_Dll.NewProc("frida_device_resume_sync")
var CL_frida_device_kill=CL_Dll.NewProc("frida_device_kill")
var CL_frida_device_kill_finish=CL_Dll.NewProc("frida_device_kill_finish")
var CL_frida_device_kill_sync=CL_Dll.NewProc("frida_device_kill_sync")
var CL_frida_device_attach=CL_Dll.NewProc("frida_device_attach")
var CL_frida_device_attach_finish=CL_Dll.NewProc("frida_device_attach_finish")
var CL_frida_device_attach_sync=CL_Dll.NewProc("frida_device_attach_sync")
var CL_frida_device_inject_library_file=CL_Dll.NewProc("frida_device_inject_library_file")
var CL_frida_device_inject_library_file_finish=CL_Dll.NewProc("frida_device_inject_library_file_finish")
var CL_frida_device_inject_library_file_sync=CL_Dll.NewProc("frida_device_inject_library_file_sync")
var CL_frida_device_inject_library_blob=CL_Dll.NewProc("frida_device_inject_library_blob")
var CL_frida_device_inject_library_blob_finish=CL_Dll.NewProc("frida_device_inject_library_blob_finish")
var CL_frida_device_inject_library_blob_sync=CL_Dll.NewProc("frida_device_inject_library_blob_sync")
var CL_frida_application_list_size=CL_Dll.NewProc("frida_application_list_size")
var CL_frida_application_list_get=CL_Dll.NewProc("frida_application_list_get")
var CL_frida_application_get_identifier=CL_Dll.NewProc("frida_application_get_identifier")
var CL_frida_application_get_name=CL_Dll.NewProc("frida_application_get_name")
var CL_frida_application_get_pid=CL_Dll.NewProc("frida_application_get_pid")
var CL_frida_application_get_small_icon=CL_Dll.NewProc("frida_application_get_small_icon")
var CL_frida_application_get_large_icon=CL_Dll.NewProc("frida_application_get_large_icon")
var CL_frida_process_list_size=CL_Dll.NewProc("frida_process_list_size")
var CL_frida_process_list_get=CL_Dll.NewProc("frida_process_list_get")
var CL_frida_process_get_pid=CL_Dll.NewProc("frida_process_get_pid")
var CL_frida_process_get_name=CL_Dll.NewProc("frida_process_get_name")
var CL_frida_process_get_small_icon=CL_Dll.NewProc("frida_process_get_small_icon")
var CL_frida_process_get_large_icon=CL_Dll.NewProc("frida_process_get_large_icon")
var CL_frida_spawn_options_new=CL_Dll.NewProc("frida_spawn_options_new")
var CL_frida_spawn_options_get_argv=CL_Dll.NewProc("frida_spawn_options_get_argv")
var CL_frida_spawn_options_get_envp=CL_Dll.NewProc("frida_spawn_options_get_envp")
var CL_frida_spawn_options_get_env=CL_Dll.NewProc("frida_spawn_options_get_env")
var CL_frida_spawn_options_get_cwd=CL_Dll.NewProc("frida_spawn_options_get_cwd")
var CL_frida_spawn_options_get_stdio=CL_Dll.NewProc("frida_spawn_options_get_stdio")
var CL_frida_spawn_options_get_aux=CL_Dll.NewProc("frida_spawn_options_get_aux")
var CL_frida_spawn_options_set_argv=CL_Dll.NewProc("frida_spawn_options_set_argv")
var CL_frida_spawn_options_set_envp=CL_Dll.NewProc("frida_spawn_options_set_envp")
var CL_frida_spawn_options_set_env=CL_Dll.NewProc("frida_spawn_options_set_env")
var CL_frida_spawn_options_set_cwd=CL_Dll.NewProc("frida_spawn_options_set_cwd")
var CL_frida_spawn_options_set_stdio=CL_Dll.NewProc("frida_spawn_options_set_stdio")
var CL_frida_spawn_list_size=CL_Dll.NewProc("frida_spawn_list_size")
var CL_frida_spawn_list_get=CL_Dll.NewProc("frida_spawn_list_get")
var CL_frida_spawn_get_pid=CL_Dll.NewProc("frida_spawn_get_pid")
var CL_frida_spawn_get_identifier=CL_Dll.NewProc("frida_spawn_get_identifier")
var CL_frida_child_list_size=CL_Dll.NewProc("frida_child_list_size")
var CL_frida_child_list_get=CL_Dll.NewProc("frida_child_list_get")
var CL_frida_child_get_pid=CL_Dll.NewProc("frida_child_get_pid")
var CL_frida_child_get_parent_pid=CL_Dll.NewProc("frida_child_get_parent_pid")
var CL_frida_child_get_origin=CL_Dll.NewProc("frida_child_get_origin")
var CL_frida_child_get_identifier=CL_Dll.NewProc("frida_child_get_identifier")
var CL_frida_child_get_path=CL_Dll.NewProc("frida_child_get_path")
var CL_frida_child_get_argv=CL_Dll.NewProc("frida_child_get_argv")
var CL_frida_child_get_envp=CL_Dll.NewProc("frida_child_get_envp")
var CL_frida_crash_get_pid=CL_Dll.NewProc("frida_crash_get_pid")
var CL_frida_crash_get_process_name=CL_Dll.NewProc("frida_crash_get_process_name")
var CL_frida_crash_get_summary=CL_Dll.NewProc("frida_crash_get_summary")
var CL_frida_crash_get_report=CL_Dll.NewProc("frida_crash_get_report")
var CL_frida_crash_load_parameters=CL_Dll.NewProc("frida_crash_load_parameters")
var CL_frida_icon_get_width=CL_Dll.NewProc("frida_icon_get_width")
var CL_frida_icon_get_height=CL_Dll.NewProc("frida_icon_get_height")
var CL_frida_icon_get_rowstride=CL_Dll.NewProc("frida_icon_get_rowstride")
var CL_frida_icon_get_pixels=CL_Dll.NewProc("frida_icon_get_pixels")
var CL_frida_session_get_pid=CL_Dll.NewProc("frida_session_get_pid")
var CL_frida_session_get_device=CL_Dll.NewProc("frida_session_get_device")
var CL_frida_session_is_detached=CL_Dll.NewProc("frida_session_is_detached")
var CL_frida_session_detach=CL_Dll.NewProc("frida_session_detach")
var CL_frida_session_detach_finish=CL_Dll.NewProc("frida_session_detach_finish")
var CL_frida_session_detach_sync=CL_Dll.NewProc("frida_session_detach_sync")
var CL_frida_session_enable_child_gating=CL_Dll.NewProc("frida_session_enable_child_gating")
var CL_frida_session_enable_child_gating_finish=CL_Dll.NewProc("frida_session_enable_child_gating_finish")
var CL_frida_session_enable_child_gating_sync=CL_Dll.NewProc("frida_session_enable_child_gating_sync")
var CL_frida_session_disable_child_gating=CL_Dll.NewProc("frida_session_disable_child_gating")
var CL_frida_session_disable_child_gating_finish=CL_Dll.NewProc("frida_session_disable_child_gating_finish")
var CL_frida_session_disable_child_gating_sync=CL_Dll.NewProc("frida_session_disable_child_gating_sync")
var CL_frida_session_create_script=CL_Dll.NewProc("frida_session_create_script")
var CL_frida_session_create_script_finish=CL_Dll.NewProc("frida_session_create_script_finish")
var CL_frida_session_create_script_sync=CL_Dll.NewProc("frida_session_create_script_sync")
var CL_frida_session_create_script_from_bytes=CL_Dll.NewProc("frida_session_create_script_from_bytes")
var CL_frida_session_create_script_from_bytes_finish=CL_Dll.NewProc("frida_session_create_script_from_bytes_finish")
var CL_frida_session_create_script_from_bytes_sync=CL_Dll.NewProc("frida_session_create_script_from_bytes_sync")
var CL_frida_session_compile_script=CL_Dll.NewProc("frida_session_compile_script")
var CL_frida_session_compile_script_finish=CL_Dll.NewProc("frida_session_compile_script_finish")
var CL_frida_session_compile_script_sync=CL_Dll.NewProc("frida_session_compile_script_sync")
var CL_frida_session_enable_debugger=CL_Dll.NewProc("frida_session_enable_debugger")
var CL_frida_session_enable_debugger_finish=CL_Dll.NewProc("frida_session_enable_debugger_finish")
var CL_frida_session_enable_debugger_sync=CL_Dll.NewProc("frida_session_enable_debugger_sync")
var CL_frida_session_disable_debugger=CL_Dll.NewProc("frida_session_disable_debugger")
var CL_frida_session_disable_debugger_finish=CL_Dll.NewProc("frida_session_disable_debugger_finish")
var CL_frida_session_disable_debugger_sync=CL_Dll.NewProc("frida_session_disable_debugger_sync")
var CL_frida_session_enable_jit=CL_Dll.NewProc("frida_session_enable_jit")
var CL_frida_session_enable_jit_finish=CL_Dll.NewProc("frida_session_enable_jit_finish")
var CL_frida_session_enable_jit_sync=CL_Dll.NewProc("frida_session_enable_jit_sync")
var CL_frida_script_get_id=CL_Dll.NewProc("frida_script_get_id")
var CL_frida_script_is_destroyed=CL_Dll.NewProc("frida_script_is_destroyed")
var CL_frida_script_load=CL_Dll.NewProc("frida_script_load")
var CL_frida_script_load_finish=CL_Dll.NewProc("frida_script_load_finish")
var CL_frida_script_load_sync=CL_Dll.NewProc("frida_script_load_sync")
var CL_frida_script_unload=CL_Dll.NewProc("frida_script_unload")
var CL_frida_script_unload_finish=CL_Dll.NewProc("frida_script_unload_finish")
var CL_frida_script_unload_sync=CL_Dll.NewProc("frida_script_unload_sync")
var CL_frida_script_eternalize=CL_Dll.NewProc("frida_script_eternalize")
var CL_frida_script_eternalize_finish=CL_Dll.NewProc("frida_script_eternalize_finish")
var CL_frida_script_eternalize_sync=CL_Dll.NewProc("frida_script_eternalize_sync")
var CL_frida_script_post=CL_Dll.NewProc("frida_script_post")
var CL_frida_script_post_finish=CL_Dll.NewProc("frida_script_post_finish")
var CL_frida_script_post_sync=CL_Dll.NewProc("frida_script_post_sync")
var CL_frida_injector_new=CL_Dll.NewProc("frida_injector_new")
var CL_frida_injector_new_inprocess=CL_Dll.NewProc("frida_injector_new_inprocess")
var CL_frida_injector_close=CL_Dll.NewProc("frida_injector_close")
var CL_frida_injector_close_finish=CL_Dll.NewProc("frida_injector_close_finish")
var CL_frida_injector_close_sync=CL_Dll.NewProc("frida_injector_close_sync")
var CL_frida_injector_inject_library_file=CL_Dll.NewProc("frida_injector_inject_library_file")
var CL_frida_injector_inject_library_file_finish=CL_Dll.NewProc("frida_injector_inject_library_file_finish")
var CL_frida_injector_inject_library_file_sync=CL_Dll.NewProc("frida_injector_inject_library_file_sync")
var CL_frida_injector_inject_library_blob=CL_Dll.NewProc("frida_injector_inject_library_blob")
var CL_frida_injector_inject_library_blob_finish=CL_Dll.NewProc("frida_injector_inject_library_blob_finish")
var CL_frida_injector_inject_library_blob_sync=CL_Dll.NewProc("frida_injector_inject_library_blob_sync")
var CL_frida_injector_demonitor_and_clone_state=CL_Dll.NewProc("frida_injector_demonitor_and_clone_state")
var CL_frida_injector_demonitor_and_clone_state_finish=CL_Dll.NewProc("frida_injector_demonitor_and_clone_state_finish")
var CL_frida_injector_demonitor_and_clone_state_sync=CL_Dll.NewProc("frida_injector_demonitor_and_clone_state_sync")
var CL_frida_injector_recreate_thread=CL_Dll.NewProc("frida_injector_recreate_thread")
var CL_frida_injector_recreate_thread_finish=CL_Dll.NewProc("frida_injector_recreate_thread_finish")
var CL_frida_injector_recreate_thread_sync=CL_Dll.NewProc("frida_injector_recreate_thread_sync")
var CL_frida_file_monitor_new=CL_Dll.NewProc("frida_file_monitor_new")
var CL_frida_file_monitor_get_path=CL_Dll.NewProc("frida_file_monitor_get_path")
var CL_frida_file_monitor_enable=CL_Dll.NewProc("frida_file_monitor_enable")
var CL_frida_file_monitor_enable_finish=CL_Dll.NewProc("frida_file_monitor_enable_finish")
var CL_frida_file_monitor_enable_sync=CL_Dll.NewProc("frida_file_monitor_enable_sync")
var CL_frida_file_monitor_disable=CL_Dll.NewProc("frida_file_monitor_disable")
var CL_frida_file_monitor_disable_finish=CL_Dll.NewProc("frida_file_monitor_disable_finish")
var CL_frida_file_monitor_disable_sync=CL_Dll.NewProc("frida_file_monitor_disable_sync")
var CL_frida_error_quark=CL_Dll.NewProc("frida_error_quark")
var CL_frida_device_type_get_type=CL_Dll.NewProc("frida_device_type_get_type")
var CL_frida_child_origin_get_type=CL_Dll.NewProc("frida_child_origin_get_type")
var CL_frida_session_detach_reason_get_type=CL_Dll.NewProc("frida_session_detach_reason_get_type")
var CL_frida_stdio_get_type=CL_Dll.NewProc("frida_stdio_get_type")
var CL_frida_unload_policy_get_type=CL_Dll.NewProc("frida_unload_policy_get_type")
var CL_frida_device_manager_get_type=CL_Dll.NewProc("frida_device_manager_get_type")
var CL_frida_device_list_get_type=CL_Dll.NewProc("frida_device_list_get_type")
var CL_frida_device_get_type=CL_Dll.NewProc("frida_device_get_type")
var CL_frida_application_list_get_type=CL_Dll.NewProc("frida_application_list_get_type")
var CL_frida_application_get_type=CL_Dll.NewProc("frida_application_get_type")
var CL_frida_process_list_get_type=CL_Dll.NewProc("frida_process_list_get_type")
var CL_frida_process_get_type=CL_Dll.NewProc("frida_process_get_type")
var CL_frida_spawn_options_get_type=CL_Dll.NewProc("frida_spawn_options_get_type")
var CL_frida_spawn_list_get_type=CL_Dll.NewProc("frida_spawn_list_get_type")
var CL_frida_spawn_get_type=CL_Dll.NewProc("frida_spawn_get_type")
var CL_frida_child_list_get_type=CL_Dll.NewProc("frida_child_list_get_type")
var CL_frida_child_get_type=CL_Dll.NewProc("frida_child_get_type")
var CL_frida_crash_get_type=CL_Dll.NewProc("frida_crash_get_type")
var CL_frida_icon_get_type=CL_Dll.NewProc("frida_icon_get_type")
var CL_frida_session_get_type=CL_Dll.NewProc("frida_session_get_type")
var CL_frida_script_get_type=CL_Dll.NewProc("frida_script_get_type")
var CL_frida_injector_get_type=CL_Dll.NewProc("frida_injector_get_type")
var CL_frida_file_monitor_get_type=CL_Dll.NewProc("frida_file_monitor_get_type")

func g_bytes_new( _data uintptr, _size gsize,)uintptr{
    r,_,_:=CL_g_bytes_new.Call(uintptr(int(_data)),uintptr(int(_size)),)
    return uintptr(r)
}

func g_bytes_new_take( _data uintptr, _size gsize,)uintptr{
    r,_,_:=CL_g_bytes_new_take.Call(uintptr(int(_data)),uintptr(int(_size)),)
    return uintptr(r)
}

func g_bytes_new_static( _data uintptr, _size gsize,)uintptr{
    r,_,_:=CL_g_bytes_new_static.Call(uintptr(int(_data)),uintptr(int(_size)),)
    return uintptr(r)
}

func g_bytes_new_with_free_func( _data uintptr, _size gsize, _free_func GDestroyNotify, _user_data uintptr,)uintptr{
    r,_,_:=CL_g_bytes_new_with_free_func.Call(uintptr(int(_data)),uintptr(int(_size)),uintptr(unsafe.Pointer(_free_func)),uintptr(int(_user_data)),)
    return uintptr(r)
}

func g_bytes_new_from_bytes( _bytes uintptr, _offset gsize, _length gsize,)uintptr{
    r,_,_:=CL_g_bytes_new_from_bytes.Call(uintptr(int(_bytes)),uintptr(int(_offset)),uintptr(int(_length)),)
    return uintptr(r)
}

func g_bytes_get_data( _bytes uintptr, _size *int,)uintptr{
    r,_,_:=CL_g_bytes_get_data.Call(uintptr(int(_bytes)),uintptr(unsafe.Pointer(_size)),)
    return uintptr(r)
}

func g_bytes_get_size( _bytes uintptr,)gsize{
    r,_,_:=CL_g_bytes_get_size.Call(uintptr(int(_bytes)),)
    return gsize(r)
}

func g_bytes_ref( _bytes uintptr,)uintptr{
    r,_,_:=CL_g_bytes_ref.Call(uintptr(int(_bytes)),)
    return uintptr(r)
}

func g_bytes_unref( _bytes uintptr,){
    CL_g_bytes_unref.Call(uintptr(int(_bytes)),)
    
}

func g_bytes_unref_to_data( _bytes uintptr, _size *int,)uintptr{
    r,_,_:=CL_g_bytes_unref_to_data.Call(uintptr(int(_bytes)),uintptr(unsafe.Pointer(_size)),)
    return uintptr(r)
}

func g_bytes_unref_to_array( _bytes uintptr,)* GByteArray{
    r,_,_:=CL_g_bytes_unref_to_array.Call(uintptr(int(_bytes)),)
    return GByteArrayFormPtr(r)
}

func g_bytes_hash( _bytes uintptr,)int{
    r,_,_:=CL_g_bytes_hash.Call(uintptr(int(_bytes)),)
    return int(r)
}

func g_bytes_equal( _bytes1 uintptr, _bytes2 uintptr,)int{
    r,_,_:=CL_g_bytes_equal.Call(uintptr(int(_bytes1)),uintptr(int(_bytes2)),)
    return int(r)
}

func g_bytes_compare( _bytes1 uintptr, _bytes2 uintptr,)int{
    r,_,_:=CL_g_bytes_compare.Call(uintptr(int(_bytes1)),uintptr(int(_bytes2)),)
    return int(r)
}

func g_main_loop_new( _context * GMainContext, _is_running int,)* GMainLoop{
    r,_,_:=CL_g_main_loop_new.Call(CObjPtr(_context),uintptr(C.int(_is_running)),)
    return GMainLoopFormPtr(r)
}

func g_main_loop_run( _loop * GMainLoop,){
    CL_g_main_loop_run.Call(CObjPtr(_loop),)
    
}

func g_main_loop_is_running( _loop * GMainLoop,)int{
    r,_,_:=CL_g_main_loop_is_running.Call(CObjPtr(_loop),)
    return int(r)
}

func g_signal_connect_data( _instance uintptr, _detailed_signal string, _c_handler GCallback, _data uintptr, _destroy_data GClosureNotify, _connect_flags GConnectFlags,)int{
    r,_,_:=CL_g_signal_connect_data.Call(uintptr(int(_instance)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_detailed_signal))),uintptr(unsafe.Pointer(_c_handler)),uintptr(int(_data)),uintptr(unsafe.Pointer(_destroy_data)),uintptr(int(_connect_flags)),)
    return int(r)
}

func g_bytes_get_type()GType{
    r,_,_:=CL_g_bytes_get_type.Call()
    return GType(r)
}

func g_bytes_icon_get_type()GType{
    r,_,_:=CL_g_bytes_icon_get_type.Call()
    return GType(r)
}

func g_bytes_icon_new( _bytes uintptr,)* GIcon{
    r,_,_:=CL_g_bytes_icon_new.Call(uintptr(int(_bytes)),)
    return GIconFormPtr(r)
}

func g_bytes_icon_get_bytes( _icon * GBytesIcon,)uintptr{
    r,_,_:=CL_g_bytes_icon_get_bytes.Call(CObjPtr(_icon),)
    return uintptr(r)
}

func frida_init(){
    CL_frida_init.Call()
    
}

func frida_shutdown(){
    CL_frida_shutdown.Call()
    
}

func frida_deinit(){
    CL_frida_deinit.Call()
    
}

func frida_get_main_context()* GMainContext{
    r,_,_:=CL_frida_get_main_context.Call()
    return GMainContextFormPtr(r)
}

func frida_unref( _obj uintptr,){
    CL_frida_unref.Call(uintptr(int(_obj)),)
    
}

func frida_version( _major *int, _minor *int, _micro *int, _nano *int,){
    CL_frida_version.Call(uintptr(unsafe.Pointer(_major)),uintptr(unsafe.Pointer(_minor)),uintptr(unsafe.Pointer(_micro)),uintptr(unsafe.Pointer(_nano)),)
    
}

func frida_version_string()string{
    r,_,_:=CL_frida_version_string.Call()
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_device_manager_new()* FridaDeviceManager{
    r,_,_:=CL_frida_device_manager_new.Call()
    return FridaDeviceManagerFormPtr(r)
}

func frida_device_manager_close( _self * FridaDeviceManager, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_close.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_close_finish( _self * FridaDeviceManager, _result * GAsyncResult,){
    CL_frida_device_manager_close_finish.Call(CObjPtr(_self),CObjPtr(_result),)
    
}

func frida_device_manager_close_sync( _self * FridaDeviceManager,){
    CL_frida_device_manager_close_sync.Call(CObjPtr(_self),)
    
}

func frida_device_manager_get_device_by_id( _self * FridaDeviceManager, _id string, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_get_device_by_id.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_id))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_get_device_by_id_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_get_device_by_id_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_get_device_by_id_sync( _self * FridaDeviceManager, _id string, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_get_device_by_id_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_id))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_get_device_by_type( _self * FridaDeviceManager, _type FridaDeviceType, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_get_device_by_type.Call(CObjPtr(_self),uintptr(int(_type)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_get_device_by_type_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_get_device_by_type_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_get_device_by_type_sync( _self * FridaDeviceManager, _type FridaDeviceType, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_get_device_by_type_sync.Call(CObjPtr(_self),uintptr(int(_type)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_get_device( _self * FridaDeviceManager, _predicate FridaDeviceManagerPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_get_device.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_get_device_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_get_device_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_get_device_sync( _self * FridaDeviceManager, _predicate FridaDeviceManagerPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_get_device_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_find_device_by_id( _self * FridaDeviceManager, _id string, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_find_device_by_id.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_id))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_find_device_by_id_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_find_device_by_id_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_find_device_by_id_sync( _self * FridaDeviceManager, _id string, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_find_device_by_id_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_id))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_find_device_by_type( _self * FridaDeviceManager, _type FridaDeviceType, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_find_device_by_type.Call(CObjPtr(_self),uintptr(int(_type)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_find_device_by_type_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_find_device_by_type_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_find_device_by_type_sync( _self * FridaDeviceManager, _type FridaDeviceType, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_find_device_by_type_sync.Call(CObjPtr(_self),uintptr(int(_type)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_find_device( _self * FridaDeviceManager, _predicate FridaDeviceManagerPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_find_device.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_find_device_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_find_device_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_find_device_sync( _self * FridaDeviceManager, _predicate FridaDeviceManagerPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_find_device_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_enumerate_devices( _self * FridaDeviceManager, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_enumerate_devices.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_enumerate_devices_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDeviceList{
    r,_,_:=CL_frida_device_manager_enumerate_devices_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceListFormPtr(r)
}

func frida_device_manager_enumerate_devices_sync( _self * FridaDeviceManager, _error ** GError,)* FridaDeviceList{
    r,_,_:=CL_frida_device_manager_enumerate_devices_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceListFormPtr(r)
}

func frida_device_manager_add_remote_device( _self * FridaDeviceManager, _host string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_add_remote_device.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_host))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_add_remote_device_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_add_remote_device_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_add_remote_device_sync( _self * FridaDeviceManager, _host string, _error ** GError,)* FridaDevice{
    r,_,_:=CL_frida_device_manager_add_remote_device_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_host))),uintptr(unsafe.Pointer(_error)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_manager_remove_remote_device( _self * FridaDeviceManager, _host string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_manager_remove_remote_device.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_host))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_manager_remove_remote_device_finish( _self * FridaDeviceManager, _result * GAsyncResult, _error ** GError,){
    CL_frida_device_manager_remove_remote_device_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_manager_remove_remote_device_sync( _self * FridaDeviceManager, _host string, _error ** GError,){
    CL_frida_device_manager_remove_remote_device_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_host))),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_list_size( _self * FridaDeviceList,)int{
    r,_,_:=CL_frida_device_list_size.Call(CObjPtr(_self),)
    return int(r)
}

func frida_device_list_get( _self * FridaDeviceList, _index int,)* FridaDevice{
    r,_,_:=CL_frida_device_list_get.Call(CObjPtr(_self),uintptr(C.int(_index)),)
    return FridaDeviceFormPtr(r)
}

func frida_device_get_id( _self * FridaDevice,)string{
    r,_,_:=CL_frida_device_get_id.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_device_get_name( _self * FridaDevice,)string{
    r,_,_:=CL_frida_device_get_name.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_device_get_icon( _self * FridaDevice,)* FridaIcon{
    r,_,_:=CL_frida_device_get_icon.Call(CObjPtr(_self),)
    return FridaIconFormPtr(r)
}

func frida_device_get_dtype( _self * FridaDevice,)FridaDeviceType{
    r,_,_:=CL_frida_device_get_dtype.Call(CObjPtr(_self),)
    return FridaDeviceType(r)
}

func frida_device_get_manager( _self * FridaDevice,)* FridaDeviceManager{
    r,_,_:=CL_frida_device_get_manager.Call(CObjPtr(_self),)
    return FridaDeviceManagerFormPtr(r)
}

func frida_device_is_lost( _self * FridaDevice,)int{
    r,_,_:=CL_frida_device_is_lost.Call(CObjPtr(_self),)
    return int(r)
}

func frida_device_get_frontmost_application( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_get_frontmost_application.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_get_frontmost_application_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaApplication{
    r,_,_:=CL_frida_device_get_frontmost_application_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaApplicationFormPtr(r)
}

func frida_device_get_frontmost_application_sync( _self * FridaDevice, _error ** GError,)* FridaApplication{
    r,_,_:=CL_frida_device_get_frontmost_application_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    return FridaApplicationFormPtr(r)
}

func frida_device_enumerate_applications( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_enumerate_applications.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_enumerate_applications_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaApplicationList{
    r,_,_:=CL_frida_device_enumerate_applications_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaApplicationListFormPtr(r)
}

func frida_device_enumerate_applications_sync( _self * FridaDevice, _error ** GError,)* FridaApplicationList{
    r,_,_:=CL_frida_device_enumerate_applications_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    return FridaApplicationListFormPtr(r)
}

func frida_device_get_process_by_pid( _self * FridaDevice, _pid int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_get_process_by_pid.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_get_process_by_pid_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_get_process_by_pid_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_get_process_by_pid_sync( _self * FridaDevice, _pid int, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_get_process_by_pid_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_get_process_by_name( _self * FridaDevice, _name string, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_get_process_by_name.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_get_process_by_name_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_get_process_by_name_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_get_process_by_name_sync( _self * FridaDevice, _name string, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_get_process_by_name_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_get_process( _self * FridaDevice, _predicate FridaDeviceProcessPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_get_process.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_get_process_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_get_process_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_get_process_sync( _self * FridaDevice, _predicate FridaDeviceProcessPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_get_process_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_find_process_by_pid( _self * FridaDevice, _pid int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_find_process_by_pid.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_find_process_by_pid_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_find_process_by_pid_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_find_process_by_pid_sync( _self * FridaDevice, _pid int, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_find_process_by_pid_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_find_process_by_name( _self * FridaDevice, _name string, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_find_process_by_name.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_find_process_by_name_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_find_process_by_name_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_find_process_by_name_sync( _self * FridaDevice, _name string, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_find_process_by_name_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_find_process( _self * FridaDevice, _predicate FridaDeviceProcessPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_find_process.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_find_process_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_find_process_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_find_process_sync( _self * FridaDevice, _predicate FridaDeviceProcessPredicate, _predicate_target uintptr, _timeout int, _cancellable * GCancellable, _error ** GError,)* FridaProcess{
    r,_,_:=CL_frida_device_find_process_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_predicate)),uintptr(int(_predicate_target)),uintptr(C.int(_timeout)),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessFormPtr(r)
}

func frida_device_enumerate_processes( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_enumerate_processes.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_enumerate_processes_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaProcessList{
    r,_,_:=CL_frida_device_enumerate_processes_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessListFormPtr(r)
}

func frida_device_enumerate_processes_sync( _self * FridaDevice, _error ** GError,)* FridaProcessList{
    r,_,_:=CL_frida_device_enumerate_processes_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    return FridaProcessListFormPtr(r)
}

func frida_device_enable_spawn_gating( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_enable_spawn_gating.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_enable_spawn_gating_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,){
    CL_frida_device_enable_spawn_gating_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_enable_spawn_gating_sync( _self * FridaDevice, _error ** GError,){
    CL_frida_device_enable_spawn_gating_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_disable_spawn_gating( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_disable_spawn_gating.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_disable_spawn_gating_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,){
    CL_frida_device_disable_spawn_gating_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_disable_spawn_gating_sync( _self * FridaDevice, _error ** GError,){
    CL_frida_device_disable_spawn_gating_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_enumerate_pending_spawn( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_enumerate_pending_spawn.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_enumerate_pending_spawn_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaSpawnList{
    r,_,_:=CL_frida_device_enumerate_pending_spawn_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaSpawnListFormPtr(r)
}

func frida_device_enumerate_pending_spawn_sync( _self * FridaDevice, _error ** GError,)* FridaSpawnList{
    r,_,_:=CL_frida_device_enumerate_pending_spawn_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    return FridaSpawnListFormPtr(r)
}

func frida_device_enumerate_pending_children( _self * FridaDevice, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_enumerate_pending_children.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_enumerate_pending_children_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaChildList{
    r,_,_:=CL_frida_device_enumerate_pending_children_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaChildListFormPtr(r)
}

func frida_device_enumerate_pending_children_sync( _self * FridaDevice, _error ** GError,)* FridaChildList{
    r,_,_:=CL_frida_device_enumerate_pending_children_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    return FridaChildListFormPtr(r)
}

func frida_device_spawn( _self * FridaDevice, _program string, _options * FridaSpawnOptions, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_spawn.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_program))),CObjPtr(_options),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_spawn_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)int{
    r,_,_:=CL_frida_device_spawn_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_device_spawn_sync( _self * FridaDevice, _program string, _options * FridaSpawnOptions, _error ** GError,)int{
    r,_,_:=CL_frida_device_spawn_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_program))),CObjPtr(_options),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_device_input( _self * FridaDevice, _pid int, _data uintptr, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_input.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(int(_data)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_input_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,){
    CL_frida_device_input_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_input_sync( _self * FridaDevice, _pid int, _data uintptr, _error ** GError,){
    CL_frida_device_input_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(int(_data)),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_resume( _self * FridaDevice, _pid int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_resume.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_resume_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,){
    CL_frida_device_resume_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_resume_sync( _self * FridaDevice, _pid int, _error ** GError,){
    CL_frida_device_resume_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_kill( _self * FridaDevice, _pid int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_kill.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_kill_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,){
    CL_frida_device_kill_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_kill_sync( _self * FridaDevice, _pid int, _error ** GError,){
    CL_frida_device_kill_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_device_attach( _self * FridaDevice, _pid int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_attach.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_attach_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)* FridaSession{
    r,_,_:=CL_frida_device_attach_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaSessionFormPtr(r)
}

func frida_device_attach_sync( _self * FridaDevice, _pid int, _error ** GError,)* FridaSession{
    r,_,_:=CL_frida_device_attach_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(_error)),)
    return FridaSessionFormPtr(r)
}

func frida_device_inject_library_file( _self * FridaDevice, _pid int, _path string, _entrypoint string, _data string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_inject_library_file.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_path))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_inject_library_file_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)int{
    r,_,_:=CL_frida_device_inject_library_file_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_device_inject_library_file_sync( _self * FridaDevice, _pid int, _path string, _entrypoint string, _data string, _error ** GError,)int{
    r,_,_:=CL_frida_device_inject_library_file_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_path))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_device_inject_library_blob( _self * FridaDevice, _pid int, _blob uintptr, _entrypoint string, _data string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_device_inject_library_blob.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(int(_blob)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_device_inject_library_blob_finish( _self * FridaDevice, _result * GAsyncResult, _error ** GError,)int{
    r,_,_:=CL_frida_device_inject_library_blob_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_device_inject_library_blob_sync( _self * FridaDevice, _pid int, _blob uintptr, _entrypoint string, _data string, _error ** GError,)int{
    r,_,_:=CL_frida_device_inject_library_blob_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(int(_blob)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_application_list_size( _self * FridaApplicationList,)int{
    r,_,_:=CL_frida_application_list_size.Call(CObjPtr(_self),)
    return int(r)
}

func frida_application_list_get( _self * FridaApplicationList, _index int,)* FridaApplication{
    r,_,_:=CL_frida_application_list_get.Call(CObjPtr(_self),uintptr(C.int(_index)),)
    return FridaApplicationFormPtr(r)
}

func frida_application_get_identifier( _self * FridaApplication,)string{
    r,_,_:=CL_frida_application_get_identifier.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_application_get_name( _self * FridaApplication,)string{
    r,_,_:=CL_frida_application_get_name.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_application_get_pid( _self * FridaApplication,)int{
    r,_,_:=CL_frida_application_get_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_application_get_small_icon( _self * FridaApplication,)* FridaIcon{
    r,_,_:=CL_frida_application_get_small_icon.Call(CObjPtr(_self),)
    return FridaIconFormPtr(r)
}

func frida_application_get_large_icon( _self * FridaApplication,)* FridaIcon{
    r,_,_:=CL_frida_application_get_large_icon.Call(CObjPtr(_self),)
    return FridaIconFormPtr(r)
}

func frida_process_list_size( _self * FridaProcessList,)int{
    r,_,_:=CL_frida_process_list_size.Call(CObjPtr(_self),)
    return int(r)
}

func frida_process_list_get( _self * FridaProcessList, _index int,)* FridaProcess{
    r,_,_:=CL_frida_process_list_get.Call(CObjPtr(_self),uintptr(C.int(_index)),)
    return FridaProcessFormPtr(r)
}

func frida_process_get_pid( _self * FridaProcess,)int{
    r,_,_:=CL_frida_process_get_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_process_get_name( _self * FridaProcess,)string{
    r,_,_:=CL_frida_process_get_name.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_process_get_small_icon( _self * FridaProcess,)* FridaIcon{
    r,_,_:=CL_frida_process_get_small_icon.Call(CObjPtr(_self),)
    return FridaIconFormPtr(r)
}

func frida_process_get_large_icon( _self * FridaProcess,)* FridaIcon{
    r,_,_:=CL_frida_process_get_large_icon.Call(CObjPtr(_self),)
    return FridaIconFormPtr(r)
}

func frida_spawn_options_new()* FridaSpawnOptions{
    r,_,_:=CL_frida_spawn_options_new.Call()
    return FridaSpawnOptionsFormPtr(r)
}

func frida_spawn_options_get_argv( _self * FridaSpawnOptions, _result_length *int,)uintptr{
    r,_,_:=CL_frida_spawn_options_get_argv.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_result_length)),)
    return uintptr(r)
}

func frida_spawn_options_get_envp( _self * FridaSpawnOptions, _result_length *int,)uintptr{
    r,_,_:=CL_frida_spawn_options_get_envp.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_result_length)),)
    return uintptr(r)
}

func frida_spawn_options_get_env( _self * FridaSpawnOptions, _result_length *int,)uintptr{
    r,_,_:=CL_frida_spawn_options_get_env.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_result_length)),)
    return uintptr(r)
}

func frida_spawn_options_get_cwd( _self * FridaSpawnOptions,)string{
    r,_,_:=CL_frida_spawn_options_get_cwd.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_spawn_options_get_stdio( _self * FridaSpawnOptions,)FridaStdio{
    r,_,_:=CL_frida_spawn_options_get_stdio.Call(CObjPtr(_self),)
    return FridaStdio(r)
}

func frida_spawn_options_get_aux( _self * FridaSpawnOptions,)* GVariantDict{
    r,_,_:=CL_frida_spawn_options_get_aux.Call(CObjPtr(_self),)
    return GVariantDictFormPtr(r)
}

func frida_spawn_options_set_argv( _self * FridaSpawnOptions, _value uintptr, _value_length int,){
    CL_frida_spawn_options_set_argv.Call(CObjPtr(_self),uintptr(int(_value)),uintptr(C.int(_value_length)),)
    
}

func frida_spawn_options_set_envp( _self * FridaSpawnOptions, _value uintptr, _value_length int,){
    CL_frida_spawn_options_set_envp.Call(CObjPtr(_self),uintptr(int(_value)),uintptr(C.int(_value_length)),)
    
}

func frida_spawn_options_set_env( _self * FridaSpawnOptions, _value uintptr, _value_length int,){
    CL_frida_spawn_options_set_env.Call(CObjPtr(_self),uintptr(int(_value)),uintptr(C.int(_value_length)),)
    
}

func frida_spawn_options_set_cwd( _self * FridaSpawnOptions, _value string,){
    CL_frida_spawn_options_set_cwd.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_value))),)
    
}

func frida_spawn_options_set_stdio( _self * FridaSpawnOptions, _value FridaStdio,){
    CL_frida_spawn_options_set_stdio.Call(CObjPtr(_self),uintptr(int(_value)),)
    
}

func frida_spawn_list_size( _self * FridaSpawnList,)int{
    r,_,_:=CL_frida_spawn_list_size.Call(CObjPtr(_self),)
    return int(r)
}

func frida_spawn_list_get( _self * FridaSpawnList, _index int,)* FridaSpawn{
    r,_,_:=CL_frida_spawn_list_get.Call(CObjPtr(_self),uintptr(C.int(_index)),)
    return FridaSpawnFormPtr(r)
}

func frida_spawn_get_pid( _self * FridaSpawn,)int{
    r,_,_:=CL_frida_spawn_get_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_spawn_get_identifier( _self * FridaSpawn,)string{
    r,_,_:=CL_frida_spawn_get_identifier.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_child_list_size( _self * FridaChildList,)int{
    r,_,_:=CL_frida_child_list_size.Call(CObjPtr(_self),)
    return int(r)
}

func frida_child_list_get( _self * FridaChildList, _index int,)* FridaChild{
    r,_,_:=CL_frida_child_list_get.Call(CObjPtr(_self),uintptr(C.int(_index)),)
    return FridaChildFormPtr(r)
}

func frida_child_get_pid( _self * FridaChild,)int{
    r,_,_:=CL_frida_child_get_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_child_get_parent_pid( _self * FridaChild,)int{
    r,_,_:=CL_frida_child_get_parent_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_child_get_origin( _self * FridaChild,)FridaChildOrigin{
    r,_,_:=CL_frida_child_get_origin.Call(CObjPtr(_self),)
    return FridaChildOrigin(r)
}

func frida_child_get_identifier( _self * FridaChild,)string{
    r,_,_:=CL_frida_child_get_identifier.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_child_get_path( _self * FridaChild,)string{
    r,_,_:=CL_frida_child_get_path.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_child_get_argv( _self * FridaChild, _result_length *int,)uintptr{
    r,_,_:=CL_frida_child_get_argv.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_result_length)),)
    return uintptr(r)
}

func frida_child_get_envp( _self * FridaChild, _result_length *int,)uintptr{
    r,_,_:=CL_frida_child_get_envp.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_result_length)),)
    return uintptr(r)
}

func frida_crash_get_pid( _self * FridaCrash,)int{
    r,_,_:=CL_frida_crash_get_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_crash_get_process_name( _self * FridaCrash,)string{
    r,_,_:=CL_frida_crash_get_process_name.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_crash_get_summary( _self * FridaCrash,)string{
    r,_,_:=CL_frida_crash_get_summary.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_crash_get_report( _self * FridaCrash,)string{
    r,_,_:=CL_frida_crash_get_report.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_crash_load_parameters( _self * FridaCrash,)* GVariantDict{
    r,_,_:=CL_frida_crash_load_parameters.Call(CObjPtr(_self),)
    return GVariantDictFormPtr(r)
}

func frida_icon_get_width( _self * FridaIcon,)int{
    r,_,_:=CL_frida_icon_get_width.Call(CObjPtr(_self),)
    return int(r)
}

func frida_icon_get_height( _self * FridaIcon,)int{
    r,_,_:=CL_frida_icon_get_height.Call(CObjPtr(_self),)
    return int(r)
}

func frida_icon_get_rowstride( _self * FridaIcon,)int{
    r,_,_:=CL_frida_icon_get_rowstride.Call(CObjPtr(_self),)
    return int(r)
}

func frida_icon_get_pixels( _self * FridaIcon,)uintptr{
    r,_,_:=CL_frida_icon_get_pixels.Call(CObjPtr(_self),)
    return uintptr(r)
}

func frida_session_get_pid( _self * FridaSession,)int{
    r,_,_:=CL_frida_session_get_pid.Call(CObjPtr(_self),)
    return int(r)
}

func frida_session_get_device( _self * FridaSession,)* FridaDevice{
    r,_,_:=CL_frida_session_get_device.Call(CObjPtr(_self),)
    return FridaDeviceFormPtr(r)
}

func frida_session_is_detached( _self * FridaSession,)int{
    r,_,_:=CL_frida_session_is_detached.Call(CObjPtr(_self),)
    return int(r)
}

func frida_session_detach( _self * FridaSession, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_detach.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_detach_finish( _self * FridaSession, _result * GAsyncResult,){
    CL_frida_session_detach_finish.Call(CObjPtr(_self),CObjPtr(_result),)
    
}

func frida_session_detach_sync( _self * FridaSession,){
    CL_frida_session_detach_sync.Call(CObjPtr(_self),)
    
}

func frida_session_enable_child_gating( _self * FridaSession, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_enable_child_gating.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_enable_child_gating_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,){
    CL_frida_session_enable_child_gating_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_enable_child_gating_sync( _self * FridaSession, _error ** GError,){
    CL_frida_session_enable_child_gating_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_disable_child_gating( _self * FridaSession, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_disable_child_gating.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_disable_child_gating_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,){
    CL_frida_session_disable_child_gating_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_disable_child_gating_sync( _self * FridaSession, _error ** GError,){
    CL_frida_session_disable_child_gating_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_create_script( _self * FridaSession, _name string, _source string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_create_script.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_source))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_create_script_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,)* FridaScript{
    r,_,_:=CL_frida_session_create_script_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaScriptFormPtr(r)
}

func frida_session_create_script_sync( _self * FridaSession, _name string, _source string, _error ** GError,)* FridaScript{
    r,_,_:=CL_frida_session_create_script_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_source))),uintptr(unsafe.Pointer(_error)),)
    return FridaScriptFormPtr(r)
}

func frida_session_create_script_from_bytes( _self * FridaSession, _bytes uintptr, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_create_script_from_bytes.Call(CObjPtr(_self),uintptr(int(_bytes)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_create_script_from_bytes_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,)* FridaScript{
    r,_,_:=CL_frida_session_create_script_from_bytes_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return FridaScriptFormPtr(r)
}

func frida_session_create_script_from_bytes_sync( _self * FridaSession, _bytes uintptr, _error ** GError,)* FridaScript{
    r,_,_:=CL_frida_session_create_script_from_bytes_sync.Call(CObjPtr(_self),uintptr(int(_bytes)),uintptr(unsafe.Pointer(_error)),)
    return FridaScriptFormPtr(r)
}

func frida_session_compile_script( _self * FridaSession, _name string, _source string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_compile_script.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_source))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_compile_script_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,)uintptr{
    r,_,_:=CL_frida_session_compile_script_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return uintptr(r)
}

func frida_session_compile_script_sync( _self * FridaSession, _name string, _source string, _error ** GError,)uintptr{
    r,_,_:=CL_frida_session_compile_script_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_name))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_source))),uintptr(unsafe.Pointer(_error)),)
    return uintptr(r)
}

func frida_session_enable_debugger( _self * FridaSession, _port int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_enable_debugger.Call(CObjPtr(_self),uintptr(C.int(_port)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_enable_debugger_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,){
    CL_frida_session_enable_debugger_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_enable_debugger_sync( _self * FridaSession, _port int, _error ** GError,){
    CL_frida_session_enable_debugger_sync.Call(CObjPtr(_self),uintptr(C.int(_port)),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_disable_debugger( _self * FridaSession, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_disable_debugger.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_disable_debugger_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,){
    CL_frida_session_disable_debugger_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_disable_debugger_sync( _self * FridaSession, _error ** GError,){
    CL_frida_session_disable_debugger_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_enable_jit( _self * FridaSession, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_session_enable_jit.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_session_enable_jit_finish( _self * FridaSession, _result * GAsyncResult, _error ** GError,){
    CL_frida_session_enable_jit_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_session_enable_jit_sync( _self * FridaSession, _error ** GError,){
    CL_frida_session_enable_jit_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_get_id( _self * FridaScript,)int{
    r,_,_:=CL_frida_script_get_id.Call(CObjPtr(_self),)
    return int(r)
}

func frida_script_is_destroyed( _self * FridaScript,)int{
    r,_,_:=CL_frida_script_is_destroyed.Call(CObjPtr(_self),)
    return int(r)
}

func frida_script_load( _self * FridaScript, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_script_load.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_script_load_finish( _self * FridaScript, _result * GAsyncResult, _error ** GError,){
    CL_frida_script_load_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_load_sync( _self * FridaScript, _error ** GError,){
    CL_frida_script_load_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_unload( _self * FridaScript, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_script_unload.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_script_unload_finish( _self * FridaScript, _result * GAsyncResult, _error ** GError,){
    CL_frida_script_unload_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_unload_sync( _self * FridaScript, _error ** GError,){
    CL_frida_script_unload_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_eternalize( _self * FridaScript, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_script_eternalize.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_script_eternalize_finish( _self * FridaScript, _result * GAsyncResult, _error ** GError,){
    CL_frida_script_eternalize_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_eternalize_sync( _self * FridaScript, _error ** GError,){
    CL_frida_script_eternalize_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_post( _self * FridaScript, _message string, _data uintptr, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_script_post.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_message))),uintptr(int(_data)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_script_post_finish( _self * FridaScript, _result * GAsyncResult, _error ** GError,){
    CL_frida_script_post_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_script_post_sync( _self * FridaScript, _message string, _data uintptr, _error ** GError,){
    CL_frida_script_post_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(syscall.StringBytePtr(_message))),uintptr(int(_data)),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_injector_new()* FridaInjector{
    r,_,_:=CL_frida_injector_new.Call()
    return FridaInjectorFormPtr(r)
}

func frida_injector_new_inprocess()* FridaInjector{
    r,_,_:=CL_frida_injector_new_inprocess.Call()
    return FridaInjectorFormPtr(r)
}

func frida_injector_close( _self * FridaInjector, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_injector_close.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_injector_close_finish( _self * FridaInjector, _result * GAsyncResult,){
    CL_frida_injector_close_finish.Call(CObjPtr(_self),CObjPtr(_result),)
    
}

func frida_injector_close_sync( _self * FridaInjector,){
    CL_frida_injector_close_sync.Call(CObjPtr(_self),)
    
}

func frida_injector_inject_library_file( _self * FridaInjector, _pid int, _path string, _entrypoint string, _data string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_injector_inject_library_file.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_path))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_injector_inject_library_file_finish( _self * FridaInjector, _result * GAsyncResult, _error ** GError,)int{
    r,_,_:=CL_frida_injector_inject_library_file_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_injector_inject_library_file_sync( _self * FridaInjector, _pid int, _path string, _entrypoint string, _data string, _error ** GError,)int{
    r,_,_:=CL_frida_injector_inject_library_file_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_path))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_injector_inject_library_blob( _self * FridaInjector, _pid int, _blob uintptr, _entrypoint string, _data string, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_injector_inject_library_blob.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(int(_blob)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_injector_inject_library_blob_finish( _self * FridaInjector, _result * GAsyncResult, _error ** GError,)int{
    r,_,_:=CL_frida_injector_inject_library_blob_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_injector_inject_library_blob_sync( _self * FridaInjector, _pid int, _blob uintptr, _entrypoint string, _data string, _error ** GError,)int{
    r,_,_:=CL_frida_injector_inject_library_blob_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(int(_blob)),uintptr(unsafe.Pointer(syscall.StringBytePtr(_entrypoint))),uintptr(unsafe.Pointer(syscall.StringBytePtr(_data))),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_injector_demonitor_and_clone_state( _self * FridaInjector, _id int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_injector_demonitor_and_clone_state.Call(CObjPtr(_self),uintptr(C.int(_id)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_injector_demonitor_and_clone_state_finish( _self * FridaInjector, _result * GAsyncResult, _error ** GError,)int{
    r,_,_:=CL_frida_injector_demonitor_and_clone_state_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_injector_demonitor_and_clone_state_sync( _self * FridaInjector, _id int, _error ** GError,)int{
    r,_,_:=CL_frida_injector_demonitor_and_clone_state_sync.Call(CObjPtr(_self),uintptr(C.int(_id)),uintptr(unsafe.Pointer(_error)),)
    return int(r)
}

func frida_injector_recreate_thread( _self * FridaInjector, _pid int, _id int, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_injector_recreate_thread.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(C.int(_id)),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_injector_recreate_thread_finish( _self * FridaInjector, _result * GAsyncResult, _error ** GError,){
    CL_frida_injector_recreate_thread_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_injector_recreate_thread_sync( _self * FridaInjector, _pid int, _id int, _error ** GError,){
    CL_frida_injector_recreate_thread_sync.Call(CObjPtr(_self),uintptr(C.int(_pid)),uintptr(C.int(_id)),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_file_monitor_new( _path string,)* FridaFileMonitor{
    r,_,_:=CL_frida_file_monitor_new.Call(uintptr(unsafe.Pointer(syscall.StringBytePtr(_path))),)
    return FridaFileMonitorFormPtr(r)
}

func frida_file_monitor_get_path( _self * FridaFileMonitor,)string{
    r,_,_:=CL_frida_file_monitor_get_path.Call(CObjPtr(_self),)
    return C.GoString((*C.char)(unsafe.Pointer(r)))
}

func frida_file_monitor_enable( _self * FridaFileMonitor, _cancellable * GCancellable, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_file_monitor_enable.Call(CObjPtr(_self),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_file_monitor_enable_finish( _self * FridaFileMonitor, _result * GAsyncResult, _error ** GError,){
    CL_frida_file_monitor_enable_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_file_monitor_enable_sync( _self * FridaFileMonitor, _cancellable * GCancellable, _error ** GError,){
    CL_frida_file_monitor_enable_sync.Call(CObjPtr(_self),CObjPtr(_cancellable),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_file_monitor_disable( _self * FridaFileMonitor, _callback GAsyncReadyCallback, _user_data uintptr,){
    CL_frida_file_monitor_disable.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_callback)),uintptr(int(_user_data)),)
    
}

func frida_file_monitor_disable_finish( _self * FridaFileMonitor, _result * GAsyncResult, _error ** GError,){
    CL_frida_file_monitor_disable_finish.Call(CObjPtr(_self),CObjPtr(_result),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_file_monitor_disable_sync( _self * FridaFileMonitor, _error ** GError,){
    CL_frida_file_monitor_disable_sync.Call(CObjPtr(_self),uintptr(unsafe.Pointer(_error)),)
    
}

func frida_error_quark()GQuark{
    r,_,_:=CL_frida_error_quark.Call()
    return GQuark(r)
}

func frida_device_type_get_type()GType{
    r,_,_:=CL_frida_device_type_get_type.Call()
    return GType(r)
}

func frida_child_origin_get_type()GType{
    r,_,_:=CL_frida_child_origin_get_type.Call()
    return GType(r)
}

func frida_session_detach_reason_get_type()GType{
    r,_,_:=CL_frida_session_detach_reason_get_type.Call()
    return GType(r)
}

func frida_stdio_get_type()GType{
    r,_,_:=CL_frida_stdio_get_type.Call()
    return GType(r)
}

func frida_unload_policy_get_type()GType{
    r,_,_:=CL_frida_unload_policy_get_type.Call()
    return GType(r)
}

func frida_device_manager_get_type()GType{
    r,_,_:=CL_frida_device_manager_get_type.Call()
    return GType(r)
}

func frida_device_list_get_type()GType{
    r,_,_:=CL_frida_device_list_get_type.Call()
    return GType(r)
}

func frida_device_get_type()GType{
    r,_,_:=CL_frida_device_get_type.Call()
    return GType(r)
}

func frida_application_list_get_type()GType{
    r,_,_:=CL_frida_application_list_get_type.Call()
    return GType(r)
}

func frida_application_get_type()GType{
    r,_,_:=CL_frida_application_get_type.Call()
    return GType(r)
}

func frida_process_list_get_type()GType{
    r,_,_:=CL_frida_process_list_get_type.Call()
    return GType(r)
}

func frida_process_get_type()GType{
    r,_,_:=CL_frida_process_get_type.Call()
    return GType(r)
}

func frida_spawn_options_get_type()GType{
    r,_,_:=CL_frida_spawn_options_get_type.Call()
    return GType(r)
}

func frida_spawn_list_get_type()GType{
    r,_,_:=CL_frida_spawn_list_get_type.Call()
    return GType(r)
}

func frida_spawn_get_type()GType{
    r,_,_:=CL_frida_spawn_get_type.Call()
    return GType(r)
}

func frida_child_list_get_type()GType{
    r,_,_:=CL_frida_child_list_get_type.Call()
    return GType(r)
}

func frida_child_get_type()GType{
    r,_,_:=CL_frida_child_get_type.Call()
    return GType(r)
}

func frida_crash_get_type()GType{
    r,_,_:=CL_frida_crash_get_type.Call()
    return GType(r)
}

func frida_icon_get_type()GType{
    r,_,_:=CL_frida_icon_get_type.Call()
    return GType(r)
}

func frida_session_get_type()GType{
    r,_,_:=CL_frida_session_get_type.Call()
    return GType(r)
}

func frida_script_get_type()GType{
    r,_,_:=CL_frida_script_get_type.Call()
    return GType(r)
}

func frida_injector_get_type()GType{
    r,_,_:=CL_frida_injector_get_type.Call()
    return GType(r)
}

func frida_file_monitor_get_type()GType{
    r,_,_:=CL_frida_file_monitor_get_type.Call()
    return GType(r)
}
