package placefile

import (
	"fmt"
)

type Polygon struct {
	R         int
	G         int
	B         int
	A         int
	Threshold int
	latlon    string
}

func (p *Polygon) String() string {
	var res string
	res = fmt.Sprintf("Polygon: %d %d %d %d\n", p.R, p.G, p.B, p.A)
	res += p.latlon
	res += "End:\n"
	return res
}

func (p *Polygon) AddRawLatLon(ll string) {
	p.latlon = ll
}

func (p *Polygon) MarshalYAML() (interface{}, error) {
	return p.String(), nil
}
