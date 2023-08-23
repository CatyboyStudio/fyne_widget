package inspector

import (
	"fmt"
	"strconv"
)

func V2S_TypeError(v any, t string) (string, error) {
	return "", fmt.Errorf("Expected %s, pass in %T", v, t)
}

func S2V_Bool(v string) (any, error) {
	return strconv.ParseBool(v)
}

func S2V_Int(v string) (any, error) {
	rv, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(rv), nil
}

func S2V_Int32(v string) (any, error) {
	rv, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(rv), nil
}

func S2V_Int64(v string) (any, error) {
	return strconv.ParseInt(v, 10, 64)
}

func S2V_Float32(v string) (any, error) {
	return strconv.ParseFloat(v, 32)
}

func S2V_Float64(v string) (any, error) {
	return strconv.ParseFloat(v, 64)
}

func S2V_String(v string) (any, error) {
	return v, nil
}

func V2S_Any(v any) (string, error) {
	if v == nil {
		return "", nil
	}
	if s, ok := v.(string); ok {
		return s, nil
	}
	if s, ok := v.(fmt.Stringer); ok {
		return s.String(), nil
	}
	if s, ok := v.(fmt.GoStringer); ok {
		return s.GoString(), nil
	}
	return fmt.Sprintf("%v", v), nil
}
