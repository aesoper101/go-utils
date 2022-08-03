package str

import "strconv"

// ToInt converts a string to an int.
func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}

// ToIntE converts a string to an int.
func ToIntE(str string) (int, error) {
	return strconv.Atoi(str)
}

// ToInt32 converts a string to an int32.
func ToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}

	return int32(i)
}

// ToInt32E converts a string to an int32. If the string is not a valid int32, it returns an error.
func ToInt32E(str string) (int32, error) {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(i), nil
}

// ToInt64 converts a string to int64
func ToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return i
}

// ToInt64E converts a string to int64. If the string is not a valid int64, it returns an error.
func ToInt64E(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// ToFloat64 converts a string to a float64.
func ToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return f
}

// ToFloat64E converts a string to a float64. If the string is not a valid float64, it returns an error.
func ToFloat64E(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

// ToFloat32 converts a string to float32
func ToFloat32(str string) float32 {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}

	return float32(f)
}

// ToFloat32E converts a string to float32. If the string is not a valid float32, it returns an error.
func ToFloat32E(str string) (float32, error) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return float32(f), nil
}

// ToInt8 converts a string to int8
func ToInt8(str string) int8 {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// ToInt8E converts a string to int8. If the string is not a valid int8, it returns an error.
func ToInt8E(str string) (int8, error) {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}

	return int8(i), nil
}

// ToUint converts a string to uint,  if the string is not a valid uint, it returns an error.
// this functionx is based on
func ToUint(str string) uint {
	i, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0
	}

	return uint(i)
}

// ToUint8 converts a string to uint8
func ToUint8(str string) uint8 {
	i, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(i)
}

// ToBool converts a string to a bool.
func ToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}

// ToBoolE converts a string to a bool. If the string is not a valid bool, it returns an error.
func ToBoolE(str string) (bool, error) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}

	return b, nil
}

// BoolToString converts a bool to a string.
func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
