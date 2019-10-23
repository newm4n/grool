package context

import (
	"github.com/newm4n/grool/pkg"
	"reflect"
	"testing"
)

type TestAStruct struct {
	BStruct *TestBStruct
}

type TestBStruct struct {
	CStruct *TestCStruct
}

type TestCStruct struct {
	Str string
	It  int
}

func TestDataContext_GetType(t *testing.T) {
	TA := &TestAStruct{BStruct: &TestBStruct{CStruct: &TestCStruct{
		Str: "TestValue",
		It:  100,
	}}}

	ctx := NewDataContext()
	err := ctx.Add("ta", TA)
	if err != nil {
		t.Fatal(err)
	}

	typ, err := ctx.GetType("ta.BStruct.CStruct.Str")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	} else if typ.Kind() != reflect.String {
		t.Errorf("Not string, but  %s", typ.Kind().String())
		t.FailNow()
	}
}

func TestDataContext_GetValue(t *testing.T) {
	TA := &TestAStruct{BStruct: &TestBStruct{CStruct: &TestCStruct{
		Str: "TestValue",
		It:  100,
	}}}

	ctx := NewDataContext()
	err := ctx.Add("ta", TA)
	if err != nil {
		t.Fatal(err)
	}

	val, err := ctx.GetValue("ta.BStruct.CStruct.Str")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	} else if pkg.ValueToInterface(val).(string) != "TestValue" {
		t.Errorf("Value is not correct")
		t.FailNow()
	}

}
