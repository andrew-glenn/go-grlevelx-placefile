package placefile

import "fmt"

type ColorLineWithText struct {
	LineWidth int
	R         int
	G         int
	B         int
	A         int
	Threshold int
	HoverText string
	latlon    string
}

func (c *ColorLineWithText) String() string {
	var res string
	res = fmt.Sprintf(
		"Color: %d %d %d %d\nLine: %d,,%s\n",
		c.R, c.G, c.B, c.A,
		c.LineWidth, c.HoverText,
	)
	res += c.latlon
	res += "End:\n"
	return res
}

func (c *ColorLineWithText) AddRawLatLon(ll string) {
	c.latlon = ll
}

func (c *ColorLineWithText) MarshalYAML() (interface{}, error) {
	return c.String(), nil
}
