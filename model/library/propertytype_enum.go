// Code generated by "go-enum -type=PropertyType"; DO NOT EDIT.

// Install go-enum by `go get install github.com/searKing/golang/tools/go-enum`
package library

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UNKNOWN-0]
	_ = x[INT-1]
	_ = x[DOUBLE-2]
	_ = x[STRING-3]
	_ = x[STRUCT-4]
	_ = x[PROTO-5]
	_ = x[BOOLEAN-6]
}

const _PropertyType_name = "UNKNOWNINTDOUBLESTRINGSTRUCTPROTOBOOLEAN"

var _PropertyType_index = [...]uint8{0, 7, 10, 16, 22, 28, 33, 40}

func _() {
	var _nil_PropertyType_value = func() (val PropertyType) { return }()

	// An "cannot convert PropertyType literal (type PropertyType) to type fmt.Stringer" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ fmt.Stringer = _nil_PropertyType_value
}

func (i PropertyType) String() string {
	if i < 0 || i >= PropertyType(len(_PropertyType_index)-1) {
		return "PropertyType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PropertyType_name[_PropertyType_index[i]:_PropertyType_index[i+1]]
}

// New returns a pointer to a new addr filled with the PropertyType value passed in.
func (i PropertyType) New() *PropertyType {
	clone := i
	return &clone
}

var _PropertyType_values = []PropertyType{0, 1, 2, 3, 4, 5, 6}

var _PropertyType_name_to_values = map[string]PropertyType{
	_PropertyType_name[0:7]:   0,
	_PropertyType_name[7:10]:  1,
	_PropertyType_name[10:16]: 2,
	_PropertyType_name[16:22]: 3,
	_PropertyType_name[22:28]: 4,
	_PropertyType_name[28:33]: 5,
	_PropertyType_name[33:40]: 6,
}

// ParsePropertyTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ParsePropertyTypeString(s string) (PropertyType, error) {
	if val, ok := _PropertyType_name_to_values[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PropertyType values", s)
}

// PropertyTypeValues returns all values of the enum
func PropertyTypeValues() []PropertyType {
	return _PropertyType_values
}

// IsAPropertyType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PropertyType) Registered() bool {
	for _, v := range _PropertyType_values {
		if i == v {
			return true
		}
	}
	return false
}

func _() {
	var _nil_PropertyType_value = func() (val PropertyType) { return }()

	// An "cannot convert PropertyType literal (type PropertyType) to type encoding.BinaryMarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.BinaryMarshaler = &_nil_PropertyType_value

	// An "cannot convert PropertyType literal (type PropertyType) to type encoding.BinaryUnmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.BinaryUnmarshaler = &_nil_PropertyType_value
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for PropertyType
func (i PropertyType) MarshalBinary() (data []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for PropertyType
func (i *PropertyType) UnmarshalBinary(data []byte) error {
	var err error
	*i, err = ParsePropertyTypeString(string(data))
	return err
}

func _() {
	var _nil_PropertyType_value = func() (val PropertyType) { return }()

	// An "cannot convert PropertyType literal (type PropertyType) to type json.Marshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ json.Marshaler = _nil_PropertyType_value

	// An "cannot convert PropertyType literal (type PropertyType) to type encoding.Unmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ json.Unmarshaler = &_nil_PropertyType_value
}

// MarshalJSON implements the json.Marshaler interface for PropertyType
func (i PropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PropertyType
func (i *PropertyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PropertyType should be a string, got %s", data)
	}

	var err error
	*i, err = ParsePropertyTypeString(s)
	return err
}

func _() {
	var _nil_PropertyType_value = func() (val PropertyType) { return }()

	// An "cannot convert PropertyType literal (type PropertyType) to type encoding.TextMarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.TextMarshaler = _nil_PropertyType_value

	// An "cannot convert PropertyType literal (type PropertyType) to type encoding.TextUnmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.TextUnmarshaler = &_nil_PropertyType_value
}

// MarshalText implements the encoding.TextMarshaler interface for PropertyType
func (i PropertyType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for PropertyType
func (i *PropertyType) UnmarshalText(text []byte) error {
	var err error
	*i, err = ParsePropertyTypeString(string(text))
	return err
}

//func _() {
//	var _nil_PropertyType_value = func() (val PropertyType) { return }()
//
//	// An "cannot convert PropertyType literal (type PropertyType) to type yaml.Marshaler" compiler error signifies that the base type have changed.
//	// Re-run the go-enum command to generate them again.
//	var _ yaml.Marshaler = _nil_PropertyType_value
//
//	// An "cannot convert PropertyType literal (type PropertyType) to type yaml.Unmarshaler" compiler error signifies that the base type have changed.
//	// Re-run the go-enum command to generate them again.
//	var _ yaml.Unmarshaler = &_nil_PropertyType_value
//}

// MarshalYAML implements a YAML Marshaler for PropertyType
func (i PropertyType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for PropertyType
func (i *PropertyType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ParsePropertyTypeString(s)
	return err
}

func _() {
	var _nil_PropertyType_value = func() (val PropertyType) { return }()

	// An "cannot convert PropertyType literal (type PropertyType) to type driver.Valuer" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ driver.Valuer = _nil_PropertyType_value

	// An "cannot convert PropertyType literal (type PropertyType) to type sql.Scanner" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ sql.Scanner = &_nil_PropertyType_value
}

func (i PropertyType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *PropertyType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := ParsePropertyTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

// PropertyTypeSliceContains reports whether sunEnums is within enums.
func PropertyTypeSliceContains(enums []PropertyType, sunEnums ...PropertyType) bool {
	var seenEnums = map[PropertyType]bool{}
	for _, e := range sunEnums {
		seenEnums[e] = false
	}

	for _, v := range enums {
		if _, has := seenEnums[v]; has {
			seenEnums[v] = true
		}
	}

	for _, seen := range seenEnums {
		if !seen {
			return false
		}
	}

	return true
}

// PropertyTypeSliceContainsAny reports whether any sunEnum is within enums.
func PropertyTypeSliceContainsAny(enums []PropertyType, sunEnums ...PropertyType) bool {
	var seenEnums = map[PropertyType]struct{}{}
	for _, e := range sunEnums {
		seenEnums[e] = struct{}{}
	}

	for _, v := range enums {
		if _, has := seenEnums[v]; has {
			return true
		}
	}

	return false
}
