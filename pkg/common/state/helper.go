package state_helper

import "strings"

type StateParser struct {
	prefix string
	names  map[int32]string
	values map[string]int32
}

func (p StateParser) ToString(x int32) string {
	s, ok := p.names[x]
	if !ok {
		return "unknown"
	}
	return strings.ToLower(strings.Trim(s, p.prefix))
}

func (p StateParser) ToValue(x string) int32 {
	i, ok := p.values[strings.ToUpper(strings.Join([]string{p.prefix, x}, "_"))]
	if !ok {
		return 0
	}
	return i
}

func NewStateParser(prefix string, names map[int32]string, values map[string]int32) StateParser {
	return StateParser{
		prefix: prefix,
		names:  names,
		values: values,
	}
}
