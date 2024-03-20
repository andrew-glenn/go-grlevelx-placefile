package placefile

import (
	"gopkg.in/yaml.v3"
)

func quotedStrPtr(s *string) string {
	var n yaml.Node
	if s != nil {
		n.Style = yaml.DoubleQuotedStyle
		n.SetString(*s)
		b, err := yaml.Marshal(&n)
		if err != nil {
			panic(err)
		}
		return string(b[:len(b)-1])
	}
	return ""
}

func QuotedString(s string) yaml.Node {
	x := yaml.Node{}
	x.SetString(s)
	x.Style = yaml.DoubleQuotedStyle
	return x
}

func StrPtr(s string) *string {
	return &s
}
