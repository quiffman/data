package list

import (
	"strconv"
	"strings"
)

// parseRaw removes any initial comment # indictators.
func parseRaw(str string) string {
	return strings.TrimLeft(strings.TrimSpace(str), "#!%")
}

// parseBool returns a bool value from a raw item.
func parseBool(str string) (bool, error) {
	return strconv.ParseBool(parseRaw(str))
}

// parseInt16 returns a int16 value from a raw item.
func parseInt16(str string) (*int16, error) {
	if str != "" {
		return func() (*int16, error) {
			if str == "n/a" {
				return nil, nil
			}
			i1, e := strconv.ParseInt(parseRaw(str), 10, 16)
			if e != nil {
				return nil, e
			}
			i := int16(i1)
			return &i, nil
		}()
	}
	return nil, nil
}

// parseFloat64 returns a float64 value from a raw item.
func parseFloat64(str string) (*float64, error) {
	if str != "" {
		return func() (*float64, error) {
			if str == "n/a" {
				return nil, nil
			}
			f, e := strconv.ParseFloat(parseRaw(str), 64)
			if e != nil {
				return nil, e
			}
			return &f, nil
		}()
	}
	return nil, nil
}
