package pkg

import (
	"fmt"
	"github.com/juju/errors"
	"reflect"
)

func ValueAdd(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() + b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() + int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) + b.Float()), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%d%s", a.Int(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) + b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() + b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) + b.Float()), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%d%s", a.Uint(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() + float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() + float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() + b.Float()), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%f%s", a.Float(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.String:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(fmt.Sprintf("%s%d", a.String(), b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(fmt.Sprintf("%s%d", a.String(), b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(fmt.Sprintf("%s%f", a.String(), b.Float())), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%s%s", a.String(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}

func ValueSub(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() - b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() - int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) - b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) - b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() - b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) - b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() - float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() - float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() - b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}

func ValueMul(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() * b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() * int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) * b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) * b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() * b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) * b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() * float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() * float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() * b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}

func ValueDiv(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() / b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() / int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) / b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) / b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() / b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) / b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() / float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() / float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() / b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}
