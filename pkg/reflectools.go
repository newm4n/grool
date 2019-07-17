package pkg

import (
	"fmt"
	"reflect"
	"time"
)

func GetFunctionList(obj interface{}) ([]string, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("param is not a struct")
	}
	ret := make([]string, 0)

	objType := reflect.TypeOf(obj)
	for i := 0; i < objType.NumMethod(); i++ {
		ret = append(ret, objType.Method(i).Name)
	}
	return ret, nil
}

func GetFunctionParameterTypes(obj interface{}, methodName string) ([]reflect.Type, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("param is not a struct")
	}
	ret := make([]reflect.Type, 0)
	objType := reflect.TypeOf(obj)

	meth, found := objType.MethodByName(methodName)
	if found {
		x := meth.Type
		for i := 1; i < x.NumIn(); i++ {
			ret = append(ret, x.In(i))
		}
	} else {
		return nil, fmt.Errorf("function %s not found", methodName)
	}
	return ret, nil
}

func GetFunctionReturnTypes(obj interface{}, methodName string) ([]reflect.Type, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("param is not a struct")
	}
	ret := make([]reflect.Type, 0)
	objType := reflect.TypeOf(obj)

	meth, found := objType.MethodByName(methodName)
	if found {
		x := meth.Type
		for i := 0; i < x.NumOut(); i++ {
			ret = append(ret, x.Out(i))
		}
	} else {
		return nil, fmt.Errorf("function %s not found", methodName)
	}
	return ret, nil
}

func InvokeFunction(obj interface{}, methodName string, param []interface{}) ([]interface{}, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("param is not a struct")
	}
	objVal := reflect.ValueOf(obj)
	funcVal := objVal.MethodByName(methodName)
	argVals := make([]reflect.Value, len(param))
	for idx, val := range param {
		argVals[idx] = reflect.ValueOf(val)
	}
	retVals := funcVal.Call(argVals)
	ret := make([]interface{}, len(retVals))
	for idx, r := range retVals {
		ret[idx] = ValueToInterface(r)
	}
	return ret, nil
}

func IsValidField(obj interface{}, fieldName string) bool {
	if !IsStruct(obj) {
		return false
	}
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	if objType.Kind() == reflect.Struct {
		fieldVal := objVal.FieldByName(fieldName)
		return fieldVal.IsValid()
	} else {
		fieldVal := objVal.Elem().FieldByName(fieldName)
		return fieldVal.IsValid()
	}
}

func IsStruct(obj interface{}) bool {
	if !reflect.ValueOf(obj).IsValid() {
		return false
	}
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return objType.Kind() == reflect.Struct
	} else {
		return objType.Elem().Kind() == reflect.Struct
	}
}

func ValueToInterface(v reflect.Value) interface{} {
	if v.Type().Kind().String() == reflect.String.String() {
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
			newPtr := reflect.New(v.Elem().Type())
			newPtr.Elem().Set(v.Elem())

			return newPtr.Interface()
		default:
			return nil
		}
	}
}

func GetAttributeList(obj interface{}) ([]string, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("param is not a struct")
	}
	strRet := make([]string, 0)
	v := reflect.ValueOf(obj)
	e := v.Elem()
	for i := 0; i < e.Type().NumField(); i++ {
		strRet = append(strRet, e.Type().Field(i).Name)
	}
	return strRet, nil
}

func GetAttributeValue(obj interface{}, fieldName string) (interface{}, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return nil, fmt.Errorf("function named %s not exist in struct", fieldName)
	}
	structval := reflect.ValueOf(obj)
	attrVal := structval.FieldByName(fieldName)
	return ValueToInterface(attrVal), nil
}

func SetAttributeValue(obj interface{}, fieldName string, value interface{}) error {
	if !IsStruct(obj) {
		return fmt.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return fmt.Errorf("function named %s not exist in struct", fieldName)
	}

	var fieldVal reflect.Value

	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	// If Obj param is a pointer
	if objType.Kind() == reflect.Ptr {
		// And it points to a struct
		if objType.Elem().Kind() == reflect.Struct {
			fieldVal = objVal.Elem().FieldByName(fieldName)
		} else {
			// If its not point to struct ... return error
			return fmt.Errorf("object is pointing a non struct. %s", objType.Elem().Kind().String())
		}
	} else {
		// If Obj param is not a pointer.
		// And its a struct
		if objType.Kind() == reflect.Struct {
			fieldVal = objVal.FieldByName(fieldName)
		} else {
			// If its not a struct ... return error
			return fmt.Errorf("object is not a struct. %s", objType.Kind().String())
		}
	}

	// Check source data type compatibility with the field type
	if fieldVal.Type() != reflect.TypeOf(value) { // pointer check
		return fmt.Errorf("can not assign type %s to %s", reflect.TypeOf(value).String(), fieldVal.Type().String())
	}
	if fieldVal.CanSet() {
		switch fieldVal.Type().Kind() {
		case reflect.String:
			fieldVal.SetString(value.(string))
			break
		case reflect.Int:
			fieldVal.SetInt(int64(value.(int)))
			break
		case reflect.Int8:
			fieldVal.SetInt(int64(value.(int8)))
			break
		case reflect.Int16:
			fieldVal.SetInt(int64(value.(int16)))
			break
		case reflect.Int32:
			fieldVal.SetInt(int64(value.(int32)))
			break
		case reflect.Int64:
			fieldVal.SetInt(value.(int64))
			break
		case reflect.Uint:
			fieldVal.SetUint(uint64(value.(uint)))
			break
		case reflect.Uint8:
			fieldVal.SetUint(uint64(value.(uint8)))
			break
		case reflect.Uint16:
			fieldVal.SetUint(uint64(value.(uint16)))
			break
		case reflect.Uint32:
			fieldVal.SetUint(uint64(value.(uint32)))
			break
		case reflect.Uint64:
			fieldVal.SetUint(uint64(value.(uint64)))
			break
		case reflect.Float32:
			fieldVal.SetFloat(float64(value.(float32)))
			break
		case reflect.Float64:
			fieldVal.SetFloat(value.(float64))
			break
		case reflect.Bool:
			fieldVal.SetBool(value.(bool))
			break
		case reflect.Ptr:
			fieldVal.Set(reflect.ValueOf(value))
			break
		case reflect.Slice:
			// todo Add setter for slice type field
			return fmt.Errorf("unsupported operation")
		case reflect.Array:
			// todo Add setter for array type field
			return fmt.Errorf("unsupported operation")
		case reflect.Map:
			// todo Add setter for map type field
			return fmt.Errorf("unsupported operation")
		case reflect.Struct:
			// todo Add setter for slice type field
			return fmt.Errorf("unsupported operation")
		default:
			return nil
		}
	} else {
		return fmt.Errorf("can not set field")
	}
	return nil
}

func IsAttributeArray(obj interface{}, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, fmt.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, fmt.Errorf("function named %s not exist in struct", fieldName)
	}
	objVal := reflect.ValueOf(obj)
	fieldVal := objVal.Elem().FieldByName(fieldName)
	return fieldVal.Type().Kind() == reflect.Array, nil
}

func IsAttributeMap(obj interface{}, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, fmt.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, fmt.Errorf("function named %s not exist in struct", fieldName)
	}
	objVal := reflect.ValueOf(obj)
	fieldVal := objVal.Elem().FieldByName(fieldName)
	return fieldVal.Type().Kind() == reflect.Map, nil
}

func IsAttributeNil(obj interface{}, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, fmt.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, fmt.Errorf("function named %s not exist in struct", fieldName)
	}
	objVal := reflect.ValueOf(obj)
	fieldVal := objVal.Elem().FieldByName(fieldName)
	return fieldVal.IsNil(), nil
}

func GetAttributeStringValue(obj interface{}, fieldName string) (string, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return "", err
	}
	return val.(string), err
}

func SetAttributeStringValue(obj interface{}, fieldName string, newValue string) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeIntValue(obj interface{}, fieldName string) (int, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int), err

}

func SetAttributeIntValue(obj interface{}, fieldName string, newValue int) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeInt8Value(obj interface{}, fieldName string) (int8, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int8), err
}

func SetAttributeInt8Value(obj interface{}, fieldName string, newValue int8) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeInt16Value(obj interface{}, fieldName string) (int16, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int16), err
}

func SetAttributeInt16Value(obj interface{}, fieldName string, newValue int16) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeInt32Value(obj interface{}, fieldName string) (int32, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int32), err
}

func SetAttributeInt32Value(obj interface{}, fieldName string, newValue int32) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeInt64Value(obj interface{}, fieldName string) (int64, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int64), err
}

func SetAttributeInt64Value(obj interface{}, fieldName string, newValue int64) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeUIntValue(obj interface{}, fieldName string) (uint, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint), err
}

func SetAttributeUIntValue(obj interface{}, fieldName string, newValue uint) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeUInt8Value(obj interface{}, fieldName string) (uint8, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint8), err
}

func SetAttributeUInt8Value(obj interface{}, fieldName string, newValue uint8) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeUInt16Value(obj interface{}, fieldName string) (uint16, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint16), err
}

func SetAttributeUInt16Value(obj interface{}, fieldName string, newValue uint16) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeUInt32Value(obj interface{}, fieldName string) (uint32, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint32), err
}

func SetAttributeUInt32Value(obj interface{}, fieldName string, newValue uint32) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeUInt64Value(obj interface{}, fieldName string) (uint64, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint64), err
}

func SetAttributeUInt64Value(obj interface{}, fieldName string, newValue uint64) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeBoolValue(obj interface{}, fieldName string) (bool, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return false, err
	}
	return val.(bool), err
}

func SetAttributeBoolValue(obj interface{}, fieldName string, newValue bool) error {
	return SetAttributeValue(obj, fieldName, newValue)
}
func GetAttributeFloat32Value(obj interface{}, fieldName string) (float32, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(float32), err
}

func SetAttributeFloat32Value(obj interface{}, fieldName string, newValue float32) error {
	return SetAttributeValue(obj, fieldName, newValue)
}
func GetAttributeFloat64Value(obj interface{}, fieldName string) (float64, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(float64), err
}

func SetAttributeFloat64Value(obj interface{}, fieldName string, newValue float64) error {
	return SetAttributeValue(obj, fieldName, newValue)
}

func GetAttributeTimeValue(obj interface{}, fieldName string) (time.Time, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return time.Now(), err
	}
	return val.(time.Time), err
}

func SetAttributeTimeValue(obj interface{}, fieldName string, newValue time.Time) error {
	return SetAttributeValue(obj, fieldName, newValue)
}
