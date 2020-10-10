package equalish

import (
	"reflect"
	"unsafe"
)

// considers numbers (of any type) equal
// if they differ by less than one part in a billion
const multiplier = 1e9

func numbersApproximate(x, y float64) bool {
	if x == y { // float64 本身就是近似值，但 0 特殊，0 时精确相等
		return true
	}
	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}
	df := diff * multiplier
	if df < x && df < y {
		return true
	}
	return false
}

func approximateEqual(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return numbersApproximate(float64(x.Int()), float64(y.Int()))

	case reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return numbersApproximate(float64(x.Uint()), float64(y.Uint()))

	case reflect.Float32, reflect.Float64:
		return numbersApproximate(float64(x.Float()), float64(y.Float()))

	case reflect.Complex64, reflect.Complex128:
		realEqualish := numbersApproximate(float64(real(x.Complex())), float64(real(y.Complex())))
		imagEqualish := numbersApproximate(float64(imag(x.Complex())), float64(imag(y.Complex())))
		return realEqualish && imagEqualish
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return approximateEqual(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !approximateEqual(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !approximateEqual(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !approximateEqual(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// Equal reports whether x and y are deeply equal, with numeric values
// differing by less than one part in a billion.
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return approximateEqual(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
