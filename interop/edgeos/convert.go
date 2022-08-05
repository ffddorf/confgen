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

	indentDepth := strings.Repeat(indent, depth)
	for _, k := range keys {
		v := in[k]
		if _, err := out.WriteString(indentDepth); err != nil {
			return err
		}
		if _, err := out.WriteString(k); err != nil {
			return err
		}
		if err := out.WriteByte(' '); err != nil {
			return err
		}
		switch t := v.(type) {
		case string:
			if strings.Contains(t, " ") {
				if err := out.WriteByte('"'); err != nil {
					return err
				}
				if _, err := out.WriteString(t); err != nil {
					return err
				}
				if err := out.WriteByte('"'); err != nil {
					return err
				}
			} else {
				if _, err := out.WriteString(t); err != nil {
					return err
				}
			}
			if err := out.WriteByte('\n'); err != nil {
				return err
			}
		case map[string]interface{}:
			if _, err := out.WriteString("{\n"); err != nil {
				return err
			}

			if err := ConfigFromMap(out, t, depth+1); err != nil {
				return err
			}

			if _, err := out.WriteString(indentDepth); err != nil {
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

func mapKeys[T any](in map[string]T) []string {
	out := make([]string, 0, len(in))
	for key := range in {
		out = append(out, key)
	}
	return out
}
