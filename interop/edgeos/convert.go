package edgeos

import (
	"fmt"
	"sort"
	"strings"
)

// ForceConsistentMapOrdering is used in tests to ensure consistent output
var ForceConsistentMapOrdering = false

type InvalidMapValueTypeError struct {
	valueType string
}

func (e InvalidMapValueTypeError) Error() string {
	return fmt.Sprintf("Invalid map value type: %s", e.valueType)
}

type StringBuilder interface {
	WriteByte(byte) error
	WriteString(string) (int, error)
}

const indent = "  "

func ConfigFromMap(out StringBuilder, in map[string]interface{}, depth int) error {
	keys := mapKeys(in)
	if ForceConsistentMapOrdering {
		sort.Strings(keys)
	}

	indentString := strings.Repeat(indent, depth)
	for _, k := range keys {
		v := in[k]
		switch t := v.(type) {
		case string:
			if err := writeKV(out, k, t, indentString); err != nil {
				return err
			}
		case []interface{}:
			for _, item := range t {
				s, ok := item.(string)
				if !ok {
					return &InvalidMapValueTypeError{
						valueType: fmt.Sprintf("%T", item),
					}
				}
				if err := writeKV(out, k, s, indentString); err != nil {
					return err
				}
			}
		case []string:
			for _, item := range t {
				if err := writeKV(out, k, item, indentString); err != nil {
					return err
				}
			}
		case map[string]interface{}:
			if _, err := out.WriteString(indentString); err != nil {
				return err
			}
			if _, err := out.WriteString(k); err != nil {
				return err
			}
			if _, err := out.WriteString(" {\n"); err != nil {
				return err
			}

			if err := ConfigFromMap(out, t, depth+1); err != nil {
				return err
			}

			if _, err := out.WriteString(indentString); err != nil {
				return err
			}
			if _, err := out.WriteString("}\n"); err != nil {
				return err
			}
		default:
			return InvalidMapValueTypeError{
				valueType: fmt.Sprintf("%T", v),
			}
		}
	}
	return nil
}

func writeKV(out StringBuilder, k, v string, indent string) error {
	quoted := strings.Contains(v, " ")
	if _, err := out.WriteString(indent); err != nil {
		return err
	}
	if _, err := out.WriteString(k); err != nil {
		return err
	}
	if err := out.WriteByte(' '); err != nil {
		return err
	}
	if quoted {
		if err := out.WriteByte('"'); err != nil {
			return err
		}
	}
	if _, err := out.WriteString(v); err != nil {
		return err
	}
	if quoted {
		if err := out.WriteByte('"'); err != nil {
			return err
		}
	}
	return out.WriteByte('\n')
}

func mapKeys[T any](in map[string]T) []string {
	out := make([]string, 0, len(in))
	for key := range in {
		out = append(out, key)
	}
	return out
}
