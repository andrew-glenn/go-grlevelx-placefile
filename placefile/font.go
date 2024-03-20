package placefile

import "fmt"

type FontAttributeFlag int

const (
	FontBold   FontAttributeFlag = 1
	FontItalic FontAttributeFlag = 2
)

type PlacefileFont struct {
	FontNumber int //FontNumberEnum
	Pixels     int
	Flags      FontAttributeFlag
	Face       string
}

func (pf PlacefileFont) String() string {
	return fmt.Sprintf(
		"%d, %d, %d, %s", pf.FontNumber, pf.Pixels, pf.Flags, quotedStrPtr(&pf.Face))
	//quotedStrPtr(&pf.Face),
}

func (pf PlacefileFont) MarshalYAML() (interface{}, error) {
	return pf.String(), nil
}
