package context

import "reflect"

type DataContext struct {
	ObjectStore map[string]interface{}
}

func (ctx *DataContext) GetType(variable string) (reflect.Type, error) {
	return nil, nil
}

func (ctx *DataContext) GetValue(variable string) (interface{}, error) {
	return nil, nil
}

func (ctx *DataContext) SetValue(variable string, newValue interface{}) error {
	return nil
}
