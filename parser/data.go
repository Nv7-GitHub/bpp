package parser

import (
	"strconv"
)

type DataType int

func (a DataType) IsEqual(b DataType) bool {
	return (a&b) != 0 || (a&b) == a
}

const (
	STRING     DataType                       = 1 << iota // string
	INT                                                   // int
	FLOAT                                                 // float64
	ARRAY                                                 // []Data
	IDENTIFIER                                            // string
	NULL                                                  // nil
	VARIADIC                                              // Multiple args
	ANY        = STRING | INT | FLOAT | ARRAY             // interface{}
)

type Data struct {
	kind DataType
	Data interface{}

	line int
}

func (d *Data) Line() int {
	return d.line
}

func (d *Data) Type() DataType {
	return d.kind
}

func ParseData(src string, line int) *Data {
	if src[0] == '"' && src[len(src)-1] == '"' {
		return &Data{
			kind: STRING,
			Data: src[1 : len(src)-1],
			line: line,
		}
	}

	intDat, err := strconv.Atoi(src)
	if err == nil {
		return &Data{
			kind: INT,
			Data: intDat,
			line: line,
		}
	}

	floatDat, err := strconv.ParseFloat(src, 64)
	if err == nil {
		return &Data{
			kind: FLOAT,
			Data: floatDat,
			line: line,
		}
	}

	return &Data{
		kind: STRING | IDENTIFIER,
		Data: src,
		line: line,
	}
}
