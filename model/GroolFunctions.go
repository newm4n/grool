package context

import (
	"github.com/newm4n/grool/model"
	"github.com/newm4n/grool/pkg"
	"log"
	"reflect"
	"time"
)

// GroolFunctions strucr hosts the built-in functions ready to invoke from the rule engine execution.
type GroolFunctions struct {
	knowledge *model.KnowledgeBase
}

func (gf *GroolFunctions) MakeTime(year, month, day, hour, minute, second int64) time.Time {
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Local)
}

// Now is an extension tn time.Now().
func (gf *GroolFunctions) Now() time.Time {
	return time.Now()
}

// Log extension to log.Print
func (gf *GroolFunctions) Log(text string) {
	log.Println(text)
}

// Enables nill checking on variables.
func (gf *GroolFunctions) IsNil(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Struct {
		return false
	}
	return !val.IsValid() || val.IsNil()
}

// Enable zero checking
func (gf *GroolFunctions) IsZero(i interface{}) bool {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Struct:
		if val.Type().String() == "time.Time" {
			return i.(time.Time).IsZero()
		}
		return false
	case reflect.Ptr:
		return val.IsNil()
	default:
		switch pkg.GetBaseKind(val) {
		case reflect.String:
			return len(val.String()) == 0
		case reflect.Int64:
			return val.Int() == 0
		case reflect.Uint64:
			return val.Uint() == 0
		case reflect.Float64:
			return val.Float() == 0
		default:
			return false
		}
	}
}

func (gf *GroolFunctions) Retract(ruleName string) {
	gf.knowledge.Retract(ruleName)
}
