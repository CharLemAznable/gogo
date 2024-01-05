package internal

import (
	. "reflect"
	"unsafe"
)

func Equal(x, y any) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := ValueOf(x)
	v2 := ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	return deepValueEqual(v1, v2, make(map[visit]bool))
}

// During deepValueEqual, must keep track of checks that are
// in progress. The comparison algorithm assumes that all
// checks in progress are true when it reencounters them.
// Visited comparisons are stored in a map indexed by visit.
type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ Type
}

// Tests for deep equality using reflected types. The map argument tracks
// comparisons that have already been seen, which allows short circuiting on
// recursive types.
func deepValueEqual(v1, v2 Value, visited map[visit]bool) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}

	if hard(v1, v2) {
		addr1 := ptrVal(v1)
		addr2 := ptrVal(v2)
		if uintptr(addr1) > uintptr(addr2) {
			// Canonicalize order to reduce number of entries in visited.
			// Assumes non-moving garbage collector.
			addr1, addr2 = addr2, addr1
		}

		// Short circuit if references are already seen.
		typ := v1.Type()
		v := visit{addr1, addr2, typ}
		if visited[v] {
			return true
		}

		// Remember for later.
		visited[v] = true
	}

	switch v1.Kind() {
	case Array:
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i), visited) {
				return false
			}
		}
		return true
	case Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}
		// Special case for []byte, which is common.
		if v1.Type().Elem().Kind() == Uint8 {
			return string(v1.Bytes()) == string(v2.Bytes())
		}
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i), visited) {
				return false
			}
		}
		return true
	case Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return deepValueEqual(v1.Elem(), v2.Elem(), visited)
	case Pointer:
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}
		return deepValueEqual(v1.Elem(), v2.Elem(), visited)
	case Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !deepValueEqual(v1.Field(i), v2.Field(i), visited) {
				return false
			}
		}
		return true
	case Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(val1, val2, visited) {
				return false
			}
		}
		return true
	case Func:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		// compare func pointer
		return v1.Pointer() == v2.Pointer()
	case Int, Int8, Int16, Int32, Int64:
		return v1.Int() == v2.Int()
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		return v1.Uint() == v2.Uint()
	case String:
		return v1.String() == v2.String()
	case Bool:
		return v1.Bool() == v2.Bool()
	case Float32, Float64:
		return v1.Float() == v2.Float()
	case Complex64, Complex128:
		return v1.Complex() == v2.Complex()
	default:
		// Normal equality suffices, just delegate to reflect.DeepEqual
		return DeepEqual(v1.Interface(), v2.Interface())
	}
}

// We want to avoid putting more in the visited map than we need to.
// For any possible reference cycle that might be encountered,
// hard(v1, v2) needs to return true for at least one of the types in the cycle,
// and it's safe and valid to get Value's internal pointer.
func hard(v1, v2 Value) bool {
	switch v1.Kind() {
	case Pointer:
		if reflectValueTypPtrData(v1) == 0 {
			// not-in-heap pointers can't be cyclic.
			// At least, all of our current uses of runtime/internal/sys.NotInHeap
			// have that property. The runtime ones aren't cyclic (and we don't use
			// DeepEqual on them anyway), and the cgo-generated ones are
			// all empty structs.
			return false
		}
		fallthrough
	case Map, Slice, Interface:
		// Nil pointers cannot be cyclic. Avoid putting them in the visited map.
		return !v1.IsNil() && !v2.IsNil()
	}
	return false
}

func reflectValueTypPtrData(v Value) uintptr {
	return uintptr(ValueOf(v).FieldByName("typ").Elem().FieldByName("ptrdata").Uint())
}

// For a Pointer or Map value, we need to check flagIndir,
// which we do by calling the pointer method.
// For Slice or Interface, flagIndir is always set,
// and using v.ptr suffices.
func ptrVal(v Value) unsafe.Pointer {
	switch v.Kind() {
	case Pointer, Map:
		return reflectValuePointer(v)
	default:
		return reflectValuePtr(v)
	}
}

const flagIndir uintptr = 1 << 7

func reflectValuePointer(v Value) unsafe.Pointer {
	if reflectValueFlag(v)&flagIndir != 0 {
		return *(*unsafe.Pointer)(reflectValuePtr(v))
	}
	return reflectValuePtr(v)
}

func reflectValuePtr(v Value) unsafe.Pointer {
	return ValueOf(v).FieldByName("ptr").UnsafePointer()
}

func reflectValueFlag(v Value) uintptr {
	return uintptr(ValueOf(v).FieldByName("flag").Uint())
}
