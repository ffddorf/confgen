package edgeos

import (
	"fmt"
	"reflect"
	"strings"
)

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
	indentDepth := strings.Repeat(indent, depth)
	for k, v := range in {
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
				valueType: reflect.TypeOf(v).Name(),
			}
		}
	}
	return nil
}
