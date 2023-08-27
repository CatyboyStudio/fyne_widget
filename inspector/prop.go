package inspector

import (
	V "cbsutil/valconv"
	"fmt"
)

type Property struct {
	Title     string
	GetValue  func() (any, error)
	SetValue  func(any) error
	GetString func() (string, error)
	SetString func(v string) error
	OnUpdate  func()
}

func none() {}

func NewProperty(t string) *Property {
	p := &Property{Title: t}
	p.GetString = p.DefaultGetString
	p.SetValue = p.DefaultSetValue
	p.OnUpdate = none
	return p
}

func (p *Property) WithGetValue(f func() (any, error)) *Property {
	p.GetValue = f
	return p
}

func (p *Property) WithSetValue(f func(any) error) *Property {
	p.SetValue = f
	return p
}

func (p *Property) WithOnUpdate(f func()) *Property {
	p.OnUpdate = f
	return p
}

func (p *Property) WithGetString(f func() (string, error)) *Property {
	p.GetString = f
	return p
}

func (p *Property) WithSetString(f func(string) error) *Property {
	p.SetString = f
	return p
}

func (p *Property) DefaultGetString() (string, error) {
	if p.GetValue == nil {
		return "", fmt.Errorf("%s miss GetValue", p.Title)
	}
	v, err := p.GetValue()
	if err != nil {
		return "", err
	}
	return V.AnyToString(v), nil
}

func (p *Property) DefaultSetValue(v any) error {
	if p.SetString == nil {
		return fmt.Errorf("%s miss SetString", p.Title)
	}
	s := V.AnyToString(v)
	return p.SetString(s)
}

// String
func String[T ~string](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return string(v), nil
		}).
		WithSetString(func(s string) error {
			return setter(T(s))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v := V.AnyToString(a)
			return setter(T(v))
		})
}

func StringRef[T ~string](label string, r *T) *Property {
	return String(label, func() (T, error) {
		return *r, nil
	}, func(s T) error {
		*r = s
		return nil
	})
}

// Bool
func Bool[T ~bool](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Bool(bool(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToBool().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToBool(a)
			return setter(T(v))
		})
}

func BoolRef[T ~bool](label string, r *T) *Property {
	return Bool(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// Int
func Int[T ~int](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Int(int(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToInt().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToInt(a)
			return setter(T(v))
		})
}

func IntRef[T ~int](label string, r *T) *Property {
	return Int(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// Int8
func Int8[T ~int8](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Int8(int8(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToInt8().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToInt8(a)
			return setter(T(v))
		})
}

func Int8Ref[T ~int8](label string, r *T) *Property {
	return Int8(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// Int16
func Int16[T ~int16](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Int16(int16(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToInt16().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToInt16(a)
			return setter(T(v))
		})
}

func Int16Ref[T ~int16](label string, r *T) *Property {
	return Int16(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// Int32
func Int32[T ~int32](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Int32(int32(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToInt32().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToInt32(a)
			return setter(T(v))
		})
}

func Int32Ref[T ~int32](label string, r *T) *Property {
	return Int32(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// UInt
func UInt[T ~uint](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.UInt(uint(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToUInt().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToUInt(a)
			return setter(T(v))
		})
}

func UIntRef[T ~uint](label string, r *T) *Property {
	return UInt(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// UInt8
func UInt8[T ~uint8](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.UInt8(uint8(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToUInt8().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToUInt8(a)
			return setter(T(v))
		})
}

func UInt8Ref[T ~uint8](label string, r *T) *Property {
	return UInt8(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// UInt16
func UInt16[T ~uint16](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.UInt16(uint16(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToUInt16().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToUInt16(a)
			return setter(T(v))
		})
}

func UInt16Ref[T ~uint16](label string, r *T) *Property {
	return UInt16(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// UInt32
func UInt32[T ~uint32](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.UInt32(uint32(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToUInt32().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToUInt32(a)
			return setter(T(v))
		})
}

func UInt32Ref[T ~uint32](label string, r *T) *Property {
	return UInt32(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// UInt64
func UInt64[T ~uint64](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.UInt64(uint64(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToUInt64().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToUInt64(a)
			return setter(T(v))
		})
}

func UInt64Ref[T ~uint64](label string, r *T) *Property {
	return UInt64(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// Float32
func Float32[T ~float32](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Float32(float32(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToFloat32().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToFloat32(a)
			return setter(T(v))
		})
}

func Float32Ref[T ~float32](label string, r *T) *Property {
	return Float32(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}

// Float64
func Float64[T ~float64](label string, getter func() (T, error), setter func(T) error) *Property {
	return NewProperty(label).
		WithGetString(func() (string, error) {
			v, err := getter()
			if err != nil {
				return "", err
			}
			return V.Float64(float64(v)).ToString().Value, nil
		}).
		WithSetString(func(s string) error {
			v := V.String(s).ToFloat64().Value
			return setter(T(v))
		}).
		WithGetValue(func() (any, error) {
			return getter()
		}).
		WithSetValue(func(a any) error {
			v, _ := V.AnyToFloat64(a)
			return setter(T(v))
		})
}

func Float64Ref[T ~float64](label string, r *T) *Property {
	return Float64(label, func() (T, error) {
		return *r, nil
	}, func(v T) error {
		*r = v
		return nil
	})
}
