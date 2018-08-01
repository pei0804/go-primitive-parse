package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Investment struct {
	Price  float64
	Str    *A
	Map    B
	Time   time.Time
	Rating int64
}

type A string
type B map[string]string

func main() {
	a := A("A")
	inv := Investment{Price: 534.432, Str: &a, Rating: 4, Map: map[string]string{}, Time: time.Now()}
	fmt.Println(parseType(inv.Price))
	fmt.Println(parseType(inv.Str))
	fmt.Println(parseType(inv.Rating))
	fmt.Println(parseType(inv.Map))
	fmt.Println(parseType(inv.Time))
}

func parseType(value interface{}) (string, error) {
	rv := reflect.ValueOf(value)
	typeName, err := parseKind(rv.Kind())
	if err != nil {
		return "", err
	}

	if typeName == "Ptr" {
		return getPointerPrimitiveType(value)
	}
	return getPrimitiveType(value)
}

func getPrimitiveType(v interface{}) (string, error) {
	return parseKind(reflect.ValueOf(v).Kind())
}

func getPointerPrimitiveType(v interface{}) (string, error) {
	return parseKind(reflect.Indirect(reflect.ValueOf(v)).Kind())
}

var ErrUnrecognizedType = errors.New("unrecognized type")

func parseKind(kind reflect.Kind) (string, error) {
	switch kind {
	case reflect.Bool:
		return "Bool", nil
	case reflect.Int:
		return "Int", nil
	case reflect.Int8:
		return "Int8", nil
	case reflect.Int16:
		return "Int16", nil
	case reflect.Int32:
		return "Int32", nil
	case reflect.Int64:
		return "Int64", nil
	case reflect.Uint:
		return "Uint", nil
	case reflect.Uint8:
		return "Uint8", nil
	case reflect.Uint16:
		return "Uint16", nil
	case reflect.Uint32:
		return "Uint32", nil
	case reflect.Uint64:
		return "Uint64", nil
	case reflect.Uintptr:
		return "Uintptr", nil
	case reflect.Float32:
		return "Float32", nil
	case reflect.Float64:
		return "Float64", nil
	case reflect.Complex64:
		return "Complex64", nil
	case reflect.Complex128:
		return "Complex128", nil
	case reflect.Array:
		return "Array", nil
	case reflect.Chan:
		return "Chan", nil
	case reflect.Func:
		return "Func", nil
	case reflect.Interface:
		return "Interface", nil
	case reflect.Map:
		return "Map", nil
	case reflect.Ptr:
		return "Ptr", nil
	case reflect.Slice:
		return "Slice", nil
	case reflect.String:
		return "String", nil
	case reflect.Struct:
		return "Struct", nil
	case reflect.UnsafePointer:
		return "UnsafePointer", nil
	default:
		return "Unrecoginized", ErrUnrecognizedType
	}
}
