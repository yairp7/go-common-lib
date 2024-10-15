package common

import "reflect"

func GetFieldValue(field reflect.Value) (any, reflect.Kind) {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int(), reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint(), reflect.Uint64
	case reflect.Float32, reflect.Float64:
		return field.Float(), reflect.Float64
	case reflect.String:
		return field.String(), reflect.String
	case reflect.Bool:
		return field.Bool(), reflect.Bool
	case reflect.Slice:
		return field.Slice(0, field.Len()), reflect.Slice
	case reflect.Struct:
		return field, reflect.Struct
	}
	return nil, reflect.Invalid
}

func SetFieldValue(field reflect.Value, newValue any) {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		newVal := newValue.(int64)
		field.SetInt(newVal)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		newVal := newValue.(uint64)
		field.SetUint(newVal)
		break
	case reflect.Float32, reflect.Float64:
		newVal := newValue.(float64)
		field.SetFloat(newVal)
		break
	case reflect.String:
		newVal := newValue.(string)
		field.SetString(newVal)
		break
	case reflect.Bool:
		newVal := newValue.(bool)
		field.SetBool(newVal)
		break
	case reflect.Slice:
		field.Set(newValue.(reflect.Value))
		break
	case reflect.Struct:
		field.Set(newValue.(reflect.Value))
		break
	}
}
