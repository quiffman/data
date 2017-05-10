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

// parseFloat32 returns a float32 value from a raw item.
func parseFloat32(str string) (*float32, error) {
	if str != "" {
		return func() (*float32, error) {
			f1, e := strconv.ParseFloat(parseRaw(str), 32)
			if e != nil {
				return nil, e
			}
			f := float32(f1)
			return &f, nil
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

// Encoding locality information, is via convention. i.e.
// <locality>
// <locality>:<tag>
// <locality>/<gateway>
// <locality>/<gateway>:<tag>
// This is mainly used for runnet and linknet coding where either a
// gateway is defined at a remote locality (<gateway>) or multiple links
// need to be represented (<tag>).

// parseLocality returns a locality string from the raw hostname item.
func parseLocality(str string) string {
	raw := parseRaw(str)
	if strings.Contains(raw, "/") {
		raw = strings.Split(raw, "/")[0]
	}

	if strings.Contains(raw, ":") {
		return strings.Split(raw, ":")[0]
	} else {
		return raw
	}
}

// parseGateway returns a deployed gateway locality string from the raw hostname item.
func parseGateway(str string) string {
	if strings.Contains(parseRaw(str), "/") {
		return strings.Split(parseRaw(str), "/")[1]
	} else if strings.Contains(parseRaw(str), ":") {
		return strings.Split(parseRaw(str), ":")[0]
	} else {
		return parseRaw(str)
	}
}

// parseTag returns a host tag string from the raw hostname item.
func parseTag(str string) string {
	raw := parseRaw(str)
	if strings.Contains(raw, "/") {
		raw = strings.Split(raw, "/")[0]
	}
	if strings.Contains(raw, ":") {
		return strings.Split(raw, ":")[1]
	} else {
		return ""
	}
}
