package external

import "C"
import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

// ModelData bind. Required to delete after making, call Delete function.
type ModelData struct {
	h uintptr
}

// Delete object.
func (m *ModelData) Delete() {
	MenohDeleteModelData(m.h)
}

// MakeModelDataFromONNX returns ModelData using ONNX file path.
func MakeModelDataFromONNX(path string) (*ModelData, error) {
	var h uintptr
	if err := checkError(MenohMakeModelDataFromONNX(path, unsafe.Pointer(&h))); err != nil {
		return nil, err
	}
	return &ModelData{h: h}, nil
}

// Optimize ModelData with profiling table.
func (m *ModelData) Optimize(table VariableProfileTable) error {
	return checkError(MenohModelDataOptimize(m.h, table.h))
}

// VariableProfileTableBuilder bind. Required to delete after making, call Delete function.
type VariableProfileTableBuilder struct {
	h uintptr
}

// Delete object.
func (b *VariableProfileTableBuilder) Delete() {
	MenohDeleteVariableProfileTableBuilder(b.h)
}

// MakeVariableProfileTableBuilder returns VariableProfileTableBuilder.
func MakeVariableProfileTableBuilder() (*VariableProfileTableBuilder, error) {
	var h uintptr
	if err := checkError(MenohMakeVariableProfileTableBuilder(unsafe.Pointer(&h))); err != nil {
		return nil, err
	}
	return &VariableProfileTableBuilder{h: h}, nil
}

// AddInputProfile adds input profile with layer name, data type and dimension size.
func (b *VariableProfileTableBuilder) AddInputProfile(name string, dtype TypeMenohDtype, dims ...int32) error {
	switch dimLen := len(dims); dimLen {
	case 2:
		return checkError(
			MenohVariableProfileTableBuilderAddInputProfileDims2(
				b.h, name, int(dtype), dims[0], dims[1]))
	case 4:
		return checkError(
			MenohVariableProfileTableBuilderAddInputProfileDims4(
				b.h, name, int(dtype), dims[0], dims[1], dims[2], dims[3]))
	default:
		return fmt.Errorf("dimension size %d is not supported", dimLen)
	}
}

// AddOutputProfile adds output profile with layer name and data type.
func (b *VariableProfileTableBuilder) AddOutputProfile(name string, dtype TypeMenohDtype) error {
	return checkError(
		MenohVariableProfileTableBuilderAddoutputProfile(b.h, name, int(dtype)))
}

// BuildVariableProfileTable returns VariableProfileTable.
func (b *VariableProfileTableBuilder) BuildVariableProfileTable(md ModelData) (
	*VariableProfileTable, error) {

	var h uintptr
	if err := checkError(MenohBuildVariableProfileTable(b.h, md.h, unsafe.Pointer(&h))); err != nil {
		return nil, err
	}
	return &VariableProfileTable{h: h}, nil
}

// VariableProfile represents profile information to make real variable.
type VariableProfile struct {
	Dtype TypeMenohDtype
	Dims  []int32
}

// VariableProfileTable bind. Required to delete after making, call Delete function.
type VariableProfileTable struct {
	h uintptr
}

// Delete object.
func (t *VariableProfileTable) Delete() {
	MenohDeleteVariableProfileTable(t.h)
}

// GetVariableProfile returns profile which setup variable information includes.
func (t *VariableProfileTable) GetVariableProfile(name string) (*VariableProfile, error) {
	var dtype int
	if err := checkError(MenohVariableProfileTableGetDtype(t.h, name, unsafe.Pointer(&dtype))); err != nil {
		return nil, err
	}
	var size int
	if err := checkError(MenohVariableProfileTableGetDimsSize(t.h, name, unsafe.Pointer(&size))); err != nil {
		return nil, err
	}
	dims := make([]int32, size)
	for i := 0; i < int(size); i++ {
		var dim int32
		if err := checkError(MenohVariableProfileTableGetDimsAt(t.h, name, i, unsafe.Pointer(&dim))); err != nil {
			return nil, err
		}
		dims[i] = dim
	}
	return &VariableProfile{
		Dtype: toDtype(dtype),
		Dims:  dims,
	}, nil
}

// ModelBuilder bind. Required to delete after making, call Delete function.
type ModelBuilder struct {
	h uintptr
}

// Delete object.
func (b *ModelBuilder) Delete() {
	MenohDeleteModelBuilder(b.h)
}

// MakeModelBuilder returns ModelBuilder.
func MakeModelBuilder(vpt VariableProfileTable) (*ModelBuilder, error) {
	var h uintptr
	if err := checkError(MenohMakeModelBuilder(vpt.h, unsafe.Pointer(&h))); err != nil {
		return nil, err
	}
	return &ModelBuilder{h: h}, nil
}

// AttachExternalBuffer attaches data buffer to get data. This process must be done
// before building Model.
func (b *ModelBuilder) AttachExternalBuffer(variableName string, bufferPtr unsafe.Pointer) error {
	return checkError(
		MenohModelBuilderAttachExternalBuffer(b.h, variableName, bufferPtr))
}

// BuildModel returns Model.
func (b *ModelBuilder) BuildModel(md ModelData, backend, backendConfig string) (*Model, error) {
	var h uintptr
	if err := checkError(MenohBuildModel(b.h, md.h, backend, backendConfig, unsafe.Pointer(&h))); err != nil {
		return nil, err
	}
	return &Model{h: h}, nil
}

// Variable represents data include data attribution and pointer.
type Variable struct {
	Dtype        TypeMenohDtype
	Dims         []int32
	BufferHandle unsafe.Pointer
}

// Model bind. Required to delete after making, call Delete function.
type Model struct {
	h uintptr
}

// Delete object.
func (m *Model) Delete() {
	MenohDeleteModel(m.h)
}

// GetVariable returns Variable, which set the target data information.
func (m *Model) GetVariable(name string) (*Variable, error) {
	var dtype int
	if err := checkError(MenohModelGetVariableDtype(m.h, name, unsafe.Pointer(&dtype))); err != nil {
		return nil, err
	}
	var size int
	if err := checkError(MenohModelGetVariableDimsSize(m.h, name, unsafe.Pointer(&size))); err != nil {
		return nil, err
	}
	dims := make([]int32, size)
	for i := 0; i < int(size); i++ {
		var dim int32
		if err := checkError(MenohModelGetVariableDimsAt(m.h, name, i, unsafe.Pointer(&dim))); err != nil {
			return nil, err
		}
		dims[i] = int32(dim)
	}
	var ptr unsafe.Pointer
	if err := checkError(MenohModelgetVariableBufferHandle(m.h, name, &ptr)); err != nil {
		return nil, err
	}

	return &Variable{
		Dtype:        toDtype(dtype),
		Dims:         dims,
		BufferHandle: ptr,
	}, nil
}

// Run calculation.
func (m *Model) Run() error {
	return checkError(MenohModelRun(m.h))
}

// TypeMenohDtype binds 'menoh_dtype_constant' enum
type TypeMenohDtype int

// Dtype
const (
	TypeFloat TypeMenohDtype = iota
)

func toDtype(typeCode int) TypeMenohDtype {
	return TypeMenohDtype(int(typeCode))
}

type typeMenohError int

// Menoh Error type
const (
	typeSuccess typeMenohError = iota
	typeSTDError
	typeUnknownError
	typeInvalidFilename
	typeUnsupportedONNXOpsetVersion
	typeONNXParseError
	typeInvalidDtype
	typeInvalidAttributeType
	typeUnsupportedOperatorAttribute
	typeDimensionMismatch
	typeVariableNotFound
	typeIndexOutOfRange
	typeJSONParseError
	typeInvalidBackendName
	typeUnsupportedOperator
	typeFailedToConfigureOperator
	typeBackendError
	typeSameNameVariableAlreadyExist
)

func (e typeMenohError) String() string {
	switch e {
	case typeSuccess:
		return "success"
	case typeSTDError:
		return "std error"
	case typeUnknownError:
		return "unknown error"
	case typeInvalidFilename:
		return "invalid filename"
	case typeUnsupportedONNXOpsetVersion:
		return "unsupported ONNX opset version"
	case typeONNXParseError:
		return "ONNX parse error"
	case typeInvalidDtype:
		return "invalid dtype"
	case typeInvalidAttributeType:
		return "invalid attribute type"
	case typeUnsupportedOperatorAttribute:
		return "unsupported operator attribute"
	case typeDimensionMismatch:
		return "dimension mismatch"
	case typeVariableNotFound:
		return "variable not found"
	case typeIndexOutOfRange:
		return "index out of range"
	case typeJSONParseError:
		return "JSON parse error"
	case typeInvalidBackendName:
		return "invalid backend name"
	case typeUnsupportedOperator:
		return "unsupported operator"
	case typeFailedToConfigureOperator:
		return "failed to configure operator"
	case typeBackendError:
		return "backend error"
	case typeSameNameVariableAlreadyExist:
		return "save name variable already exist"
	default:
		return "unknown type error"
	}
}

func checkError(err error) error {
	if err == nil {
		return nil
	}
	if err == syscall.EINVAL {
		return err
	}
	errType := typeMenohError(uintptr(err.(syscall.Errno)))
	if errType == typeSuccess {
		return nil
	}
	// MenohGetLastErrorMessage returns uintptr, not convert to string correctly
	// So not get last message and only return error type
	return errors.New(errType.String())
}
