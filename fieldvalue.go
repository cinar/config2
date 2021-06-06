package config2

import (
	"fmt"
	"math/bits"
	"reflect"
	"strconv"
)

type fieldValue struct {
	Kind  reflect.Kind
	Value reflect.Value
}

func (fv fieldValue) String() string {
	if fv.Value.IsValid() {
		return fmt.Sprint(fv.Value.Interface())
	}

	return ""
}

func (fv fieldValue) Set(s string) error {
	return setFieldValue(&fv, s)
}

func (fv fieldValue) IsBoolFlag() bool {
	return fv.Kind == reflect.Bool
}

func setBoolFieldValue(fv *fieldValue, s string) error {
	value, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}

	fv.Value.SetBool(value)

	return nil
}

func setIntFieldValue(fv *fieldValue, s string, bitSize int) error {
	value, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return err
	}

	fv.Value.SetInt(value)

	return nil
}

func setFloatFieldValue(fv *fieldValue, s string, bitSize int) error {
	value, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return err
	}

	fv.Value.SetFloat(value)

	return nil
}

func setFieldValue(fv *fieldValue, s string) error {
	var err error

	switch fv.Kind {
	case reflect.String:
		fv.Value.SetString(s)

	case reflect.Bool:
		err = setBoolFieldValue(fv, s)

	case reflect.Int:
		err = setIntFieldValue(fv, s, bits.UintSize)

	case reflect.Int8:
		err = setIntFieldValue(fv, s, 8)

	case reflect.Int16:
		err = setIntFieldValue(fv, s, 16)

	case reflect.Int32:
		err = setIntFieldValue(fv, s, 32)

	case reflect.Int64:
		err = setIntFieldValue(fv, s, 64)

	case reflect.Uint:
		err = setIntFieldValue(fv, s, bits.UintSize)

	case reflect.Uint8:
		err = setIntFieldValue(fv, s, 8)

	case reflect.Uint16:
		err = setIntFieldValue(fv, s, 16)

	case reflect.Uint32:
		err = setIntFieldValue(fv, s, 32)

	case reflect.Uint64:
		err = setIntFieldValue(fv, s, 64)

	case reflect.Float32:
		err = setFloatFieldValue(fv, s, 32)

	case reflect.Float64:
		err = setFloatFieldValue(fv, s, 64)

	default:
		err = fmt.Errorf("Unsupported field value %s", fv.Kind.String())
	}

	return err
}
