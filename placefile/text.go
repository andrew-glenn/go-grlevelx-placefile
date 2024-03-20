package placefile

import (
	"fmt"
)

type Text struct {
	//Text:    lat, lon, fontNumber, "string", "hover"
	//lat, lon:
	//	coordinates for the center of the text string
	//
	//fontNumber:
	//	fontNumber from earlier Font statement
	//
	//string:
	//	text to display
	//
	//quotedStrPtr (optional):
	//	text displayed when the mouse cursor hovers over the
	//	lat, lon of the string
	Lat        int64
	Lon        int64
	FontNumber int
	TextString string
	HoverText  string
}

func (t Text) String() string {
	return fmt.Sprintf(
		"Text: %d, %d, %d, %s",
		t.Lat,
		t.Lon,
		t.FontNumber,
		quotedStrPtr(&t.TextString),
	)
}

type Color struct {
	R int
	G int
	B int
}

func (c Color) String() string {
	return fmt.Sprintf("Color: %d %d %d", c.R, c.G, c.B)
}

type ColorText struct {
	Color
	Text
}
