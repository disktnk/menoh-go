// Code generated by the command above; DO NOT EDIT.

package external

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modmenoh = syscall.NewLazyDLL("menoh.dll")

	procmenoh_delete_model_data                                       = modmenoh.NewProc("menoh_delete_model_data")
	procmenoh_make_model_data_from_onnx                               = modmenoh.NewProc("menoh_make_model_data_from_onnx")
	procmenoh_model_data_optimize                                     = modmenoh.NewProc("menoh_model_data_optimize")
	procmenoh_delete_variable_profile_table_builder                   = modmenoh.NewProc("menoh_delete_variable_profile_table_builder")
	procmenoh_make_variable_profile_table_builder                     = modmenoh.NewProc("menoh_make_variable_profile_table_builder")
	procmenoh_variable_profile_table_builder_add_input_profile_dims_2 = modmenoh.NewProc("menoh_variable_profile_table_builder_add_input_profile_dims_2")
	procmenoh_variable_profile_table_builder_add_input_profile_dims_4 = modmenoh.NewProc("menoh_variable_profile_table_builder_add_input_profile_dims_4")
	procmenoh_variable_profile_table_builder_add_output_profile       = modmenoh.NewProc("menoh_variable_profile_table_builder_add_output_profile")
	procmenoh_build_variable_profile_table                            = modmenoh.NewProc("menoh_build_variable_profile_table")
	procmenoh_delete_variable_profile_table                           = modmenoh.NewProc("menoh_delete_variable_profile_table")
	procmenoh_variable_profile_table_get_dtype                        = modmenoh.NewProc("menoh_variable_profile_table_get_dtype")
	procmenoh_variable_profile_table_get_dims_size                    = modmenoh.NewProc("menoh_variable_profile_table_get_dims_size")
	procmenoh_variable_profile_table_get_dims_at                      = modmenoh.NewProc("menoh_variable_profile_table_get_dims_at")
	procmenoh_delete_model_builder                                    = modmenoh.NewProc("menoh_delete_model_builder")
	procmenoh_make_model_builder                                      = modmenoh.NewProc("menoh_make_model_builder")
	procmenoh_model_builder_attach_external_buffer                    = modmenoh.NewProc("menoh_model_builder_attach_external_buffer")
	procmenoh_build_model                                             = modmenoh.NewProc("menoh_build_model")
	procmenoh_delete_model                                            = modmenoh.NewProc("menoh_delete_model")
	procmenoh_model_get_variable_dtype                                = modmenoh.NewProc("menoh_model_get_variable_dtype")
	procmenoh_model_get_variable_dims_size                            = modmenoh.NewProc("menoh_model_get_variable_dims_size")
	procmenoh_model_get_variable_dims_at                              = modmenoh.NewProc("menoh_model_get_variable_dims_at")
	procmenoh_model_get_variable_buffer_handle                        = modmenoh.NewProc("menoh_model_get_variable_buffer_handle")
	procmenoh_model_run                                               = modmenoh.NewProc("menoh_model_run")
	procmenoh_get_last_error_message                                  = modmenoh.NewProc("menoh_get_last_error_message")
)

func MenohDeleteModelData(mdHandle uintptr) {
	syscall.Syscall(procmenoh_delete_model_data.Addr(), 1, uintptr(mdHandle), 0, 0)
	return
}

func MenohMakeModelDataFromONNX(path string, mdHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(path)
	if err != nil {
		return
	}
	return _MenohMakeModelDataFromONNX(_p0, mdHandle)
}

func _MenohMakeModelDataFromONNX(path *byte, mdHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_make_model_data_from_onnx.Addr(), 2, uintptr(unsafe.Pointer(path)), uintptr(mdHandle), 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohModelDataOptimize(mdHandle uintptr, vptHandle uintptr) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_model_data_optimize.Addr(), 2, uintptr(mdHandle), uintptr(vptHandle), 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohDeleteVariableProfileTableBuilder(vptHandle uintptr) {
	syscall.Syscall(procmenoh_delete_variable_profile_table_builder.Addr(), 1, uintptr(vptHandle), 0, 0)
	return
}

func MenohMakeVariableProfileTableBuilder(vptbHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_make_variable_profile_table_builder.Addr(), 1, uintptr(vptbHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohVariableProfileTableBuilderAddInputProfileDims2(vptbHandle uintptr, name string, dtype int, dim1 int32, dim2 int32) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohVariableProfileTableBuilderAddInputProfileDims2(vptbHandle, _p0, dtype, dim1, dim2)
}

func _MenohVariableProfileTableBuilderAddInputProfileDims2(vptbHandle uintptr, name *byte, dtype int, dim1 int32, dim2 int32) (err error) {
	r1, _, _ := syscall.Syscall6(procmenoh_variable_profile_table_builder_add_input_profile_dims_2.Addr(), 5, uintptr(vptbHandle), uintptr(unsafe.Pointer(name)), uintptr(dtype), uintptr(dim1), uintptr(dim2), 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohVariableProfileTableBuilderAddInputProfileDims4(vptbHandle uintptr, name string, dtype int, dim1 int32, dim2 int32, dim3 int32, dim4 int32) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohVariableProfileTableBuilderAddInputProfileDims4(vptbHandle, _p0, dtype, dim1, dim2, dim3, dim4)
}

func _MenohVariableProfileTableBuilderAddInputProfileDims4(vptbHandle uintptr, name *byte, dtype int, dim1 int32, dim2 int32, dim3 int32, dim4 int32) (err error) {
	r1, _, _ := syscall.Syscall9(procmenoh_variable_profile_table_builder_add_input_profile_dims_4.Addr(), 7, uintptr(vptbHandle), uintptr(unsafe.Pointer(name)), uintptr(dtype), uintptr(dim1), uintptr(dim2), uintptr(dim3), uintptr(dim4), 0, 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohVariableProfileTableBuilderAddoutputProfile(vptbHandle uintptr, name string, dtype int) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohVariableProfileTableBuilderAddoutputProfile(vptbHandle, _p0, dtype)
}

func _MenohVariableProfileTableBuilderAddoutputProfile(vptbHandle uintptr, name *byte, dtype int) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_variable_profile_table_builder_add_output_profile.Addr(), 3, uintptr(vptbHandle), uintptr(unsafe.Pointer(name)), uintptr(dtype))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohBuildVariableProfileTable(vptbHandle uintptr, mdHandle uintptr, vptHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_build_variable_profile_table.Addr(), 3, uintptr(vptbHandle), uintptr(mdHandle), uintptr(vptHandle))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohDeleteVariableProfileTable(vptHandle uintptr) {
	syscall.Syscall(procmenoh_delete_variable_profile_table.Addr(), 1, uintptr(vptHandle), 0, 0)
	return
}

func MenohVariableProfileTableGetDtype(vptHandle uintptr, name string, dtypeHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohVariableProfileTableGetDtype(vptHandle, _p0, dtypeHandle)
}

func _MenohVariableProfileTableGetDtype(vptHandle uintptr, name *byte, dtypeHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_variable_profile_table_get_dtype.Addr(), 3, uintptr(vptHandle), uintptr(unsafe.Pointer(name)), uintptr(dtypeHandle))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohVariableProfileTableGetDimsSize(vptHandle uintptr, name string, sizeHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohVariableProfileTableGetDimsSize(vptHandle, _p0, sizeHandle)
}

func _MenohVariableProfileTableGetDimsSize(vptHandle uintptr, name *byte, sizeHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_variable_profile_table_get_dims_size.Addr(), 3, uintptr(vptHandle), uintptr(unsafe.Pointer(name)), uintptr(sizeHandle))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohVariableProfileTableGetDimsAt(vptHandle uintptr, name string, pos int, dimHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohVariableProfileTableGetDimsAt(vptHandle, _p0, pos, dimHandle)
}

func _MenohVariableProfileTableGetDimsAt(vptHandle uintptr, name *byte, pos int, dimHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall6(procmenoh_variable_profile_table_get_dims_at.Addr(), 4, uintptr(vptHandle), uintptr(unsafe.Pointer(name)), uintptr(pos), uintptr(dimHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohDeleteModelBuilder(mbHandle uintptr) {
	syscall.Syscall(procmenoh_delete_model_builder.Addr(), 1, uintptr(mbHandle), 0, 0)
	return
}

func MenohMakeModelBuilder(vptHandle uintptr, mbHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_make_model_builder.Addr(), 2, uintptr(vptHandle), uintptr(mbHandle), 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohModelBuilderAttachExternalBuffer(mbHandle uintptr, name string, buffer unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohModelBuilderAttachExternalBuffer(mbHandle, _p0, buffer)
}

func _MenohModelBuilderAttachExternalBuffer(mbHandle uintptr, name *byte, buffer unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_model_builder_attach_external_buffer.Addr(), 3, uintptr(mbHandle), uintptr(unsafe.Pointer(name)), uintptr(buffer))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohBuildModel(mbHandle uintptr, mdHandle uintptr, backend string, backendConfig string, modelHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(backend)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = syscall.BytePtrFromString(backendConfig)
	if err != nil {
		return
	}
	return _MenohBuildModel(mbHandle, mdHandle, _p0, _p1, modelHandle)
}

func _MenohBuildModel(mbHandle uintptr, mdHandle uintptr, backend *byte, backendConfig *byte, modelHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall6(procmenoh_build_model.Addr(), 5, uintptr(mbHandle), uintptr(mdHandle), uintptr(unsafe.Pointer(backend)), uintptr(unsafe.Pointer(backendConfig)), uintptr(modelHandle), 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohDeleteModel(mHandle uintptr) {
	syscall.Syscall(procmenoh_delete_model.Addr(), 1, uintptr(mHandle), 0, 0)
	return
}

func MenohModelGetVariableDtype(mHandle uintptr, name string, dtypeHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohModelGetVariableDtype(mHandle, _p0, dtypeHandle)
}

func _MenohModelGetVariableDtype(mHandle uintptr, name *byte, dtypeHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_model_get_variable_dtype.Addr(), 3, uintptr(mHandle), uintptr(unsafe.Pointer(name)), uintptr(dtypeHandle))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohModelGetVariableDimsSize(mHandle uintptr, name string, sizeHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohModelGetVariableDimsSize(mHandle, _p0, sizeHandle)
}

func _MenohModelGetVariableDimsSize(mHandle uintptr, name *byte, sizeHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_model_get_variable_dims_size.Addr(), 3, uintptr(mHandle), uintptr(unsafe.Pointer(name)), uintptr(sizeHandle))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohModelGetVariableDimsAt(mHandle uintptr, name string, pos int, dimHandle unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohModelGetVariableDimsAt(mHandle, _p0, pos, dimHandle)
}

func _MenohModelGetVariableDimsAt(mHandle uintptr, name *byte, pos int, dimHandle unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall6(procmenoh_model_get_variable_dims_at.Addr(), 4, uintptr(mHandle), uintptr(unsafe.Pointer(name)), uintptr(pos), uintptr(dimHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohModelgetVariableBufferHandle(mHandle uintptr, name string, buffer *unsafe.Pointer) (err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _MenohModelgetVariableBufferHandle(mHandle, _p0, buffer)
}

func _MenohModelgetVariableBufferHandle(mHandle uintptr, name *byte, buffer *unsafe.Pointer) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_model_get_variable_buffer_handle.Addr(), 3, uintptr(mHandle), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(buffer)))
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohModelRun(mHandle uintptr) (err error) {
	r1, _, _ := syscall.Syscall(procmenoh_model_run.Addr(), 1, uintptr(mHandle), 0, 0)
	if r1 != 0 {
		err = errnoErr(syscall.Errno(r1))
	}
	return
}

func MenohGetLastErrorMessage() (msg string) {
	r0, _, _ := syscall.Syscall(procmenoh_get_last_error_message.Addr(), 0, 0, 0, 0)
	msg = string(r0)
	return
}
