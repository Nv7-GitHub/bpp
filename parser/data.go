package parser

import (
	"strconv"
	"strings"
)

// ParseData parses a literal and converts it to a Data statement with the corresponding type
func ParseData(src string, line int) *Data {
	src = strings.TrimSpace(src)
	if len(src) == 0 {
		return &Data{
			kind:           NULL,
			Data:           nil,
			BasicStatement: &BasicStatement{line: line},
		}
	}

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
