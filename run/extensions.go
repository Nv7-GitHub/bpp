package run

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/parser"
)

type ExtensionGroup []*Extension

func (e ExtensionGroup) BuildForParser() map[string]parser.ExternalFunction {
	out := make(map[string]parser.ExternalFunction, len(e))
	for _, ext := range e {
		out[ext.Name] = parser.ExternalFunction{
			ParTypes: ext.ParTypes,
			RetType:  ext.RetType,
		}
	}
	return out
}

type Extension struct {
	Name     string
	Fn       reflect.Value // reflect.FuncOf
	ParTypes []parser.Type

	ReturnsVal bool
	RetType    parser.Type
}

func (e *Extension) Call(vals []interface{}) (interface{}, error) {
	parVals := make([]reflect.Value, len(vals))
	for i, val := range vals {
		parVals[i] = reflect.ValueOf(val)
	}

	retVals := e.Fn.Call(parVals)
	if len(retVals) == 0 {
		return nil, nil
	}

	if e.ReturnsVal {
		return retVals[0].Interface(), nil
	}

	return nil, nil
}

func NewExtension(fnObj interface{}, name string) (*Extension, error) {
	fn := reflect.ValueOf(fnObj)

	// Get par types
	fnTyp := fn.Type()
	parCnt := fnTyp.NumIn()
	parTypes := make([]parser.Type, parCnt)
	var err error

	for i := 0; i < parCnt; i++ {
		parTypes[i], err = getParserType(fnTyp.In(i))
		if err != nil {
			return nil, err
		}
	}

	// Get ret type
	numOut := fnTyp.NumOut()
	hasRet := numOut > 0
	var retType parser.Type = nil

	if hasRet && numOut > 1 {
		return nil, fmt.Errorf("extension cannot return multiple values")
	} else if hasRet {
		retType, err = getParserType(fnTyp.Out(0))
		if err != nil {
			return nil, err
		}
	}

	return &Extension{
		Fn:         fn,
		Name:       name,
		ParTypes:   parTypes,
		ReturnsVal: hasRet,
		RetType:    retType,
	}, nil
}

var typMap = map[reflect.Kind]parser.BasicType{
	reflect.Int:     parser.INT,
	reflect.Float64: parser.FLOAT,
	reflect.String:  parser.STRING,
}

func getParserType(typ reflect.Type) (parser.Type, error) {
	basicTyp, exists := typMap[typ.Kind()]
	if exists {
		return basicTyp, nil
	}

	if typ.Kind() == reflect.Slice {
		valTyp, err := getParserType(typ.Elem())
		if err != nil {
			return nil, err
		}
		return parser.NewArrayType(valTyp), nil
	}

	return nil, fmt.Errorf("unknown type \"%s\"", typ.String())
}
