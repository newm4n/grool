package pkg

import (
	"fmt"
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
	G int8
	H int16
	I int32
	J int64
	K uint
	L uint8
	M uint16
	N uint32
	O uint64
	P float32
	Q []int
	R map[int]string
	S TestSubObjectNoPtr
}

func (to *TestObject) FunctionA(arg1 string, arg2 string) (string, error) {
	return fmt.Sprintf("A call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to *TestObject) FunctionB(arg1, arg2 string) (string, error) {
	return fmt.Sprintf("B call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to *TestObject) FunctionC(arg1 int, arg2 string) (string, error) {
	return fmt.Sprintf("C call Arg1 : %d and Arg2 : %s", arg1, arg2), nil
}

func (to TestObject) FunctionD(arg1 string, arg2 string) (string, error) {
	return fmt.Sprintf("A call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to TestObject) FunctionE(arg1, arg2 string) (string, error) {
	return fmt.Sprintf("B call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to TestObject) FunctionF(arg1 int, arg2 string) (string, error) {
	return fmt.Sprintf("C call Arg1 : %d and Arg2 : %s", arg1, arg2), nil
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

func TestGetFunctionList(t *testing.T) {
	to := &TestObject{}
	functions, err := GetFunctionList(to)
	if err != nil {
		t.Error("Got error")
		t.FailNow()
	}
	if len(functions) != 6 {
		t.Errorf("Got %d functions", len(functions))
		t.FailNow()
	}
	if functions[0] != "FunctionA" {
		t.Errorf("1st function name %s", functions[0])
		t.FailNow()
	}
	if functions[1] != "FunctionB" {
		t.Errorf("2nd function name %s", functions[1])
		t.FailNow()
	}

	to2 := TestObject{}
	functions, err = GetFunctionList(to2)
	if err != nil {
		t.Error("Got error")
		t.FailNow()
	}
	if len(functions) != 3 {
		t.Errorf("Got %d functions", len(functions))
		t.FailNow()
	}
	if functions[0] != "FunctionD" {
		t.Errorf("1st function name %s", functions[0])
		t.FailNow()
	}
	if functions[1] != "FunctionE" {
		t.Errorf("2nd function name %s", functions[1])
		t.FailNow()
	}
}

func TestGetFunctionParameterTypes(t *testing.T) {
	to := &TestObject{}
	types, err := GetFunctionParameterTypes(to, "FunctionC")
	if err != nil {
		t.Errorf("Error : %v", err)
		t.FailNow()
	}
	if len(types) != 2 {
		t.Errorf("Invalid argument count : %d", len(types))
		for idx, typ := range types {
			if typ.Kind().String() == "ptr" {
				t.Errorf("#%d : %s", idx, typ.Elem().Name())
			} else {
				t.Errorf("#%d : %s", idx, typ.Kind().String())
			}
		}
		t.FailNow()
	}
	if types[0].Kind().String() != "int" || types[1].Kind().String() != "string" {
		t.Errorf("Invalid argument type kind")
		t.FailNow()
	}
}

func TestGetFunctionReturnTypes(t *testing.T) {
	to := &TestObject{}
	types, err := GetFunctionReturnTypes(to, "FunctionC")
	if err != nil {
		t.Errorf("Error : %v", err)
		t.FailNow()
	}
	if len(types) != 2 {
		t.Errorf("Invalid return count : %d", len(types))
		for idx, typ := range types {
			if typ.Kind().String() == "ptr" {
				t.Errorf("#%d : %s", idx, typ.Elem().Name())
			} else {
				t.Errorf("#%d : %s", idx, typ.Kind().String())
			}
		}
		t.FailNow()
	}
	if types[0].Kind().String() != "string" || types[1].Kind().String() != "interface" {
		t.Errorf("Invalid argument type kind")
		for idx, typ := range types {
			if typ.Kind().String() == "ptr" {
				t.Errorf("#%d : %s", idx, typ.Elem().Name())
			} else {
				t.Errorf("#%d : %s", idx, typ.Kind().String())
			}
		}
		t.FailNow()
	}
}

func TestInvokeFunction(t *testing.T) {
	to := &TestObject{}
	param := make([]interface{}, 2)
	param[0] = 10
	param[1] = "Ten"
	rets, err := InvokeFunction(to, "FunctionC", param)
	if err != nil {
		t.Errorf("Got error : %v", err)
	} else {
		if len(rets) != 2 {
			t.Errorf("Invalid ret outs : %d", len(rets))
		}
		if rets[0].(string) != "C call Arg1 : 10 and Arg2 : Ten" {
			t.Errorf("Invalid turn : %s", rets[0].(string))
		}
		if rets[1] != nil {
			t.Errorf("2nd return should be nil")
		}
	}
}

func TestGetAttributeList(t *testing.T) {
	to := &TestObject{}
	names, err := GetAttributeList(to)
	if err != nil {
		t.Errorf("Got error : %v", err)
		t.FailNow()
	}
	if len(names) != 19 {
		t.Errorf("Invalid field count : %d", len(names))
		t.FailNow()
	}
	check := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	}
	for i := 0; i < len(check); i++ {
		if names[i] != check[i] {
			t.Errorf("Attribute #%d, expect %s, but actual %s", i, check[i], names[i])
			t.FailNow()
		}
	}
}

func TestGetAttributeValue(t *testing.T) {
	to := TestObjectNoPtr{
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
	itv, err := GetAttributeValue(to, "A")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	if itv.(string) != "string data" {
		t.FailNow()
	}
}

func TestSetAttributeValue(t *testing.T) {
	testObject := &TestObject{
		A: "string data",
		B: 123,
		C: 456.789,
		D: true,
	}
	err := SetAttributeValue(testObject, "A", "strong data")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	err = SetAttributeValue(testObject, "B", 456)
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	err = SetAttributeValue(testObject, "B", 456.123)
	if err == nil {
		t.Errorf("Should not be able to set with different type")
		t.FailNow()
	}
	if testObject.A != "strong data" && testObject.B != 456 {
		t.Errorf("Setting string fail")
		t.FailNow()
	}

	tso := &TestSubObject{
		A: "TSO",
		B: 2019,
		C: 2019.6,
	}
	err = SetAttributeValue(testObject, "F", tso)
	if err != nil {
		t.Errorf("Should not be able to set with different type : %v", err)
		t.FailNow()
	}
	if testObject.F.A != "TSO" && testObject.F.B != 2019 {
		t.Errorf("Setting object fail")
		t.FailNow()
	}
}
