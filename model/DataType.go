package model

var (
	StringDataType  = DataType(0)
	DecimalDataType = DataType(1)
	BooleanDataType = DataType(2)
	FloatDataType   = DataType(3)
	TimeDataType    = DataType(4)
)

type DataType int
