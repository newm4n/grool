package context

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/newm4n/grool/pkg"
	"reflect"
	"strings"
)

func NewDataContext() *DataContext {
	return &DataContext{
		ObjectStore: make(map[string]interface{}),
	}
}

type DataContext struct {
	ObjectStore         map[string]interface{}
	VariableChangeCount uint64
}

func (ctx *DataContext) Add(key string, obj interface{}) {
	ctx.ObjectStore[key] = obj
}

func (ctx *DataContext) ExecMethod(methodName string, args []reflect.Value) (reflect.Value, error) {
	varArray := strings.Split(methodName, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		return traceMethod(val, varArray[1:], args)
	} else {
		return reflect.ValueOf(nil), errors.Errorf("data context not found '%s'", varArray[0])
	}
}

func (ctx *DataContext) GetType(variable string) (reflect.Type, error) {
	varArray := strings.Split(variable, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		return traceType(val, varArray[1:])
	} else {
		return nil, errors.Errorf("data context not found '%s'", varArray[0])
	}
}

func (ctx *DataContext) GetValue(variable string) (reflect.Value, error) {
	varArray := strings.Split(variable, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		vval, err := traceValue(val, varArray[1:])
		if err != nil {
			fmt.Printf("blah %s = %v\n", variable, vval)
		}
		return vval, err
	} else {
		return reflect.ValueOf(nil), fmt.Errorf("data context not found '%s'", varArray[0])
	}
}

func (ctx *DataContext) SetValue(variable string, newValue reflect.Value) error {
	varArray := strings.Split(variable, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		err := traceSetValue(val, varArray[1:], newValue)
		if err == nil {
			ctx.VariableChangeCount++
		}
		return err
	} else {
		return errors.Errorf("data context not found '%s'", varArray[0])
	}
}

func traceType(obj interface{}, path []string) (reflect.Type, error) {
	if len(path) == 1 {
		return pkg.GetAttributeType(obj, path[0])
	} else if len(path) > 1 {
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return nil, errors.Trace(err)
		}
		return traceType(pkg.ValueToInterface(objVal), path[1:])
	} else {
		return reflect.TypeOf(obj), nil
	}
}

func traceValue(obj interface{}, path []string) (reflect.Value, error) {
	if len(path) == 1 {
		return pkg.GetAttributeValue(obj, path[0])
	} else if len(path) > 1 {
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return objVal, errors.Trace(err)
		}
		return traceValue(pkg.ValueToInterface(objVal), path[1:])
	} else {
		return reflect.ValueOf(obj), nil
	}
}

func traceSetValue(obj interface{}, path []string, newValue reflect.Value) error {
	if len(path) == 1 {
		return pkg.SetAttributeValue(obj, path[0], newValue)
	} else if len(path) > 1 {
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return errors.Trace(err)
		}
		return traceSetValue(objVal, path[1:], newValue)
	} else {
		return errors.Errorf("no attribute path specified")
	}
}

func traceMethod(obj interface{}, path []string, args []reflect.Value) (reflect.Value, error) {
	if len(path) == 1 {
		types, err := pkg.GetFunctionParameterTypes(obj, path[0])
		if err != nil {
			return reflect.ValueOf(nil), errors.Errorf("error while fetching function %s() parameter types.", path[0])
		} else {
			if len(types) != len(args) {
				return reflect.ValueOf(nil), errors.Errorf("invalid argument count for function %s(). need %d argument while there are %d", path[0], len(types), len(args))
			} else {
				iargs := make([]interface{}, 0)
				for i, t := range types {
					if t.Kind() != args[i].Kind() {
						if t.Kind() == reflect.Interface && args[i].Kind() == reflect.Struct {
							iargs = append(iargs, pkg.ValueToInterface(args[i]))
						} else {
							return reflect.ValueOf(nil), errors.Errorf("invalid argument types for function %s(). argument #%d, require %s but %s", path[0], i, t.Kind().String(), args[i].Kind().String())
						}
					} else {
						iargs = append(iargs, pkg.ValueToInterface(args[i]))
					}
				}
				rets, err := pkg.InvokeFunction(obj, path[0], iargs)
				if err != nil {
					return reflect.ValueOf(nil), err
				} else {
					if len(rets) > 1 {
						return reflect.ValueOf(rets[0]), errors.Errorf("multiple return value for function %s(). ", path[0])
					} else if len(rets) == 1 {
						return reflect.ValueOf(rets[0]), nil
					} else {
						return reflect.ValueOf(nil), nil
					}
				}
			}
		}
	} else if len(path) > 1 {
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		return traceMethod(objVal, path[1:], args)
	} else {
		return reflect.ValueOf(nil), errors.Errorf("no function path specified")
	}
}
