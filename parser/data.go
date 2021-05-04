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
	*BasicStatement
	kind DataType
	Data interface{}
}

func (d *Data) Type() DataType {
	return d.kind
}

func ParseData(src string, line int) *Data {
	if src[0] == '"' && src[len(src)-1] == '"' {
		return &Data{
			kind:           STRING,
			Data:           src[1 : len(src)-1],
			BasicStatement: &BasicStatement{line: line},
		}
	}

	intDat, err := strconv.Atoi(src)
	if err == nil {
		return &Data{
			kind:           INT,
			Data:           intDat,
			BasicStatement: &BasicStatement{line: line},
		}
	}

	floatDat, err := strconv.ParseFloat(src, 64)
	if err == nil {
		return &Data{
			kind:           FLOAT,
			Data:           floatDat,
			BasicStatement: &BasicStatement{line: line},
		}
	}

	return &Data{
		kind:           STRING | IDENTIFIER,
		Data:           src,
		BasicStatement: &BasicStatement{line: line},
	}
}
