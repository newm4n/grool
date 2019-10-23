package model

import (
	"github.com/newm4n/grool/pkg"
	"log"
	"reflect"
	"strings"
	"time"
)

// GroolFunctions strucr hosts the built-in functions ready to invoke from the rule engine execution.
type GroolFunctions struct {
	Knowledge *KnowledgeBase
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

// Retract will retract a rule from next evaluation cycle.
func (gf *GroolFunctions) Retract(ruleName string) {
	gf.Knowledge.Retract(strings.ReplaceAll(ruleName, "\"", ""))
}

// GetTimeYear will get the year value of time
func (gf *GroolFunctions) GetTimeYear(time time.Time) int {
	return time.Year()
}

// GetTimeMonth will get the month value of time
func (gf *GroolFunctions) GetTimeMonth(time time.Time) int {
	return int(time.Month())
}

// GetTimeDay will get the day value of time
func (gf *GroolFunctions) GetTimeDay(time time.Time) int {
	return time.Day()
}

// GetTimeHour will get the hour value of time
func (gf *GroolFunctions) GetTimeHour(time time.Time) int {
	return time.Hour()
}

// GetTimeMinute will get the minute value of time
func (gf *GroolFunctions) GetTimeMinute(time time.Time) int {
	return time.Minute()
}

// GetTimeSecond will get the second value of time
func (gf *GroolFunctions) GetTimeSecond(time time.Time) int {
	return time.Second()
}

// IsTimeBefore will check if the 1st argument is before the 2nd argument.
func (gf *GroolFunctions) IsTimeBefore(time, before time.Time) bool {
	return time.Before(before)
}

// IsTimeAfter will check if the 1st argument is after the 2nd argument.
func (gf *GroolFunctions) IsTimeAfter(time, after time.Time) bool {
	return time.After(after)
}

// TimeFormat will format a time according to format layout.
func (gf *GroolFunctions) TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}
