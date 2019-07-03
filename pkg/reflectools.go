package pkg

import (
	"reflect"
)

type Variable interface {
	GetObject() interface{}
	Name() string
	Type() string
	GetValue() interface{}
}

type VariableImpl struct {
	obj interface{}
	idx int
}

func (impl *VariableImpl) GetObject() interface{} {
	return impl.obj
}
func (impl *VariableImpl) Name() string {
	return reflect.ValueOf(impl.obj).Elem().Type().Field(impl.idx).Name
}
func (impl *VariableImpl) Type() string {
	return reflect.ValueOf(impl.obj).Elem().Type().Field(impl.idx).Type.String()
}
func (impl *VariableImpl) GetValue() interface{} {
	v := reflect.ValueOf(impl.obj).Elem().Field(impl.idx)
	if impl.Type() == reflect.String.String() {
		return v.String()
	} else {
		switch v.Type().Kind() {
		case reflect.Int:
			return int(v.Int())
		case reflect.Int8:
			return int8(v.Int())
		case reflect.Int16:
			return int16(v.Int())
		case reflect.Int32:
			return int32(v.Int())
		case reflect.Int64:
			return v.Int()
		case reflect.Uint:
			return uint(v.Uint())
		case reflect.Uint8:
			return uint8(v.Uint())
		case reflect.Uint16:
			return uint16(v.Uint())
		case reflect.Uint32:
			return uint32(v.Uint())
		case reflect.Uint64:
			return v.Uint()
		case reflect.Float32:
			return float32(v.Float())
		case reflect.Float64:
			return v.Float()
		case reflect.Bool:
			return v.Bool()
		case reflect.Ptr:

			//fmt.Println(v.Type(), v.Kind(), reflect.Indirect(v).Type())
			//fmt.Println(v.Elem().Type(), v.Elem().Kind())

			newPtr := reflect.New(v.Elem().Type())
			newPtr.Elem().Set(v.Elem())

			return newPtr.Interface()
		default:
			return nil
		}
	}
}

func GetMemberVariables(obj interface{}) ([]Variable, error) {
	ret := make([]Variable, 0)
	v := reflect.ValueOf(obj)
	e := v.Elem()
	for i := 0; i < e.NumField(); i++ {
		ret = append(ret, &VariableImpl{
			obj: obj,
			idx: i,
		})
	}
	return ret, nil
}
