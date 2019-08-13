package context

import (
	"github.com/newm4n/grool/pkg"
	"log"
	"reflect"
	"time"
)

type GroolFunctions struct {
}

func (gf *GroolFunctions) MakeTime(year, month, day, hour, minute, second int64) time.Time {
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Local)
}

func (gf *GroolFunctions) Now() time.Time {
	return time.Now()
}

func (gf *GroolFunctions) Log(text string) {
	log.Println(text)
}

func (gf *GroolFunctions) IsNil(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Struct {
		return false
	}
	return val.IsValid() == false || val.IsNil()
}

func (gf *GroolFunctions) IsZero(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Struct {
		if val.Type().String() == "time.Time" {
			return i.(time.Time).IsZero()
		} else {
			return false
		}
	} else if val.Kind() == reflect.Ptr {
		return val.IsNil()
	} else {
		if pkg.GetBaseKind(val) == reflect.String {
			return len(val.String()) == 0
		} else if pkg.GetBaseKind(val) == reflect.Int64 {
			return val.Int() == 0
		} else if pkg.GetBaseKind(val) == reflect.Uint64 {
			return val.Uint() == 0
		} else if pkg.GetBaseKind(val) == reflect.Float64 {
			return val.Float() == 0
		} else {
			return false
		}
	}
}
