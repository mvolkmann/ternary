package ternary

// Any implements a ternary for any type.
// Typically a type assertion will be required.
// For example, color := ter.Any(temperature >= 100, "red", "blue").(string)
func Any(cond bool, t interface{}, f interface{}) interface{} {
	if cond {
		return t
	}
	return f
}

// Int implements a ternary for int values.
func Int(cond bool, t int, f int) int {
	if cond {
		return t
	}
	return f
}

// String implements a ternary for string values.
func String(cond bool, t string, f string) string {
	if cond {
		return t
	}
	return f
}
