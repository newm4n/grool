package pkg

import (
	"reflect"
	"testing"
	"time"
)

type TestObject struct {
	A string
	B int
	C float64
	D bool
	E time.Time
	F *TestSubObject
}

type TestSubObject struct {
	A string
	B int
	C float64
}

type TestObjectNoPtr struct {
	A string
	B int
	C float64
	D bool
	E time.Time
	F TestSubObjectNoPtr
}

type TestSubObjectNoPtr struct {
	A string
	B int
	C float64
}

func TestGetMemberVariables(t *testing.T) {
	testObject := &TestObject{
		A: "string data",
		B: 123,
		C: 456.789,
		D: true,
		E: time.Date(2019, time.January, 1, 1, 1, 1, 0, time.Local),
		F: &TestSubObject{
			A: "string",
			B: 123,
			C: 456.789,
		},
	}

	vars, err := GetMemberVariables(testObject)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(vars) != 6 {
		t.Error("Invalid member variable count")
		t.Fail()
	}

	if vars[0].Name() != "A" {
		t.Errorf("Expect vars[0] = a but %s", vars[0].Name())
		t.Fail()
	}
	if vars[0].Type() != "string" {
		t.Errorf("Expect vars[0] type = string but %s", vars[0].Type())
		t.Fail()
	}
	if vars[0].GetValue().(string) != "string data" {
		t.Errorf("Expect vars[0] type = string but %s", vars[0].Type())
		t.Fail()
	}

	if vars[1].Name() != "B" {
		t.Errorf("Expect vars[1] = b but %s", vars[1].Name())
		t.Fail()
	}
	if vars[1].Type() != "int" {
		t.Errorf("Expect vars[1] type = int but %s", vars[1].Type())
		t.Fail()
	}
	if vars[1].GetValue().(int) != 123 {
		t.Errorf("Expect vars[1] type = string but %s", vars[1].Type())
		t.Fail()
	}

	if vars[2].Name() != "C" {
		t.Errorf("Expect vars[2] = c but %s", vars[2].Name())
		t.Fail()
	}
	if vars[2].Type() != "float64" {
		t.Errorf("Expect vars[2] type = float64 but %s", vars[2].Type())
		t.Fail()
	}
	if vars[3].Name() != "D" {
		t.Errorf("Expect vars[3] = d but %s", vars[3].Name())
		t.Fail()
	}
	if vars[3].Type() != "bool" {
		t.Errorf("Expect vars[3] type = bool but %s", vars[3].Type())
		t.Fail()
	}
	if vars[4].Name() != "E" {
		t.Errorf("Expect vars[4] = e but %s", vars[4].Name())
		t.Fail()
	}
	if vars[4].Type() != "time.Time" {
		t.Errorf("Expect vars[4] type = time.Time but %s", vars[4].Type())
		t.Fail()
	}
	if vars[5].Name() != "F" {
		t.Errorf("Expect vars[5] = f but %s", vars[5].Name())
		t.Fail()
	}
	if vars[5].Type() != "*pkg.TestSubObject" {
		t.Errorf("Expect vars[5] type = *pkg.TestSubObject but %s", vars[5].Type())
		t.Fail()
	}

	tso := vars[5].GetValue().(*TestSubObject)
	if reflect.ValueOf(tso).Type().String() != reflect.ValueOf(testObject.F).Type().String() {
		t.Errorf(reflect.ValueOf(tso).Type().String(), "!=", reflect.ValueOf(testObject.F).Type().String())
		t.Fail()
	}

	if tso.A != testObject.F.A {
		t.Errorf("%v != %v", tso.A, testObject.F.A)
		t.Fail()
	}
	if tso.B != testObject.F.B {
		t.Errorf("%v != %v", tso.B, testObject.F.B)
		t.Fail()
	}
	if tso.C != testObject.F.C {
		t.Errorf("%v != %v", tso.C, testObject.F.C)
		t.Fail()
	}

	if tso != testObject.F {
		t.Errorf("Expect vars[5] object equals TestSubObject but its not")
		t.Fail()
	}

}

func TestGetMemberPtrVariables(t *testing.T) {
	testObject := TestObjectNoPtr{
		A: "string data",
		B: 123,
		C: 456.789,
		D: true,
		E: time.Date(2019, time.January, 1, 1, 1, 1, 0, time.Local),
		F: TestSubObjectNoPtr{
			A: "string",
			B: 123,
			C: 456.789,
		},
	}

	vars, err := GetMemberVariables(testObject)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(vars) != 6 {
		t.Error("Invalid member variable count")
		t.Fail()
	}

	if vars[0].Name() != "A" {
		t.Errorf("Expect vars[0] = a but %s", vars[0].Name())
		t.Fail()
	}
	if vars[0].Type() != "string" {
		t.Errorf("Expect vars[0] type = string but %s", vars[0].Type())
		t.Fail()
	}
	if vars[0].GetValue().(string) != "string data" {
		t.Errorf("Expect vars[0] type = string but %s", vars[0].Type())
		t.Fail()
	}

	if vars[1].Name() != "B" {
		t.Errorf("Expect vars[1] = b but %s", vars[1].Name())
		t.Fail()
	}
	if vars[1].Type() != "int" {
		t.Errorf("Expect vars[1] type = int but %s", vars[1].Type())
		t.Fail()
	}
	if vars[1].GetValue().(int) != 123 {
		t.Errorf("Expect vars[1] type = string but %s", vars[1].Type())
		t.Fail()
	}

	if vars[2].Name() != "C" {
		t.Errorf("Expect vars[2] = c but %s", vars[2].Name())
		t.Fail()
	}
	if vars[2].Type() != "float64" {
		t.Errorf("Expect vars[2] type = float64 but %s", vars[2].Type())
		t.Fail()
	}
	if vars[3].Name() != "D" {
		t.Errorf("Expect vars[3] = d but %s", vars[3].Name())
		t.Fail()
	}
	if vars[3].Type() != "bool" {
		t.Errorf("Expect vars[3] type = bool but %s", vars[3].Type())
		t.Fail()
	}
	if vars[4].Name() != "E" {
		t.Errorf("Expect vars[4] = e but %s", vars[4].Name())
		t.Fail()
	}
	if vars[4].Type() != "time.Time" {
		t.Errorf("Expect vars[4] type = time.Time but %s", vars[4].Type())
		t.Fail()
	}
	if vars[5].Name() != "F" {
		t.Errorf("Expect vars[5] = f but %s", vars[5].Name())
		t.Fail()
	}
	if vars[5].Type() != "*pkg.TestSubObject" {
		t.Errorf("Expect vars[5] type = *pkg.TestSubObject but %s", vars[5].Type())
		t.Fail()
	}

	tso := vars[5].GetValue().(TestSubObjectNoPtr)
	if reflect.ValueOf(tso).Type().String() != reflect.ValueOf(testObject.F).Type().String() {
		t.Errorf(reflect.ValueOf(tso).Type().String(), "!=", reflect.ValueOf(testObject.F).Type().String())
		t.Fail()
	}

	if tso.A != testObject.F.A {
		t.Errorf("%v != %v", tso.A, testObject.F.A)
		t.Fail()
	}
	if tso.B != testObject.F.B {
		t.Errorf("%v != %v", tso.B, testObject.F.B)
		t.Fail()
	}
	if tso.C != testObject.F.C {
		t.Errorf("%v != %v", tso.C, testObject.F.C)
		t.Fail()
	}

	if tso != testObject.F {
		t.Errorf("Expect vars[5] object equals TestSubObject but its not")
		t.Fail()
	}

}
