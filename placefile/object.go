package placefile

import "fmt"

type Object struct {
	Lat       float32
	Lon       float32
	Threshold int
	icon      []Icon
	text      []Text
	colortext []ColorText
}

func (o *Object) AddIcon(angle, fileNumber, iconNumber int, hT string) {
	i := Icon{
		Lat:        0,
		Lon:        0,
		Angle:      angle,
		FileNumber: fileNumber,
		IconNumber: iconNumber,
		HoverText:  hT,
	}
	o.icon = append(o.icon, i)
}

func (o *Object) AddText(lat, lon int64, fontNumber int, textString string, hT *string) {
	t := Text{
		Lat:        lat,
		Lon:        lon,
		FontNumber: fontNumber,
		TextString: textString,
		HoverText:  quotedStrPtr(hT),
	}
	o.text = append(o.text, t)
	return
}

func (o *Object) AddColorText(r, g, b int, lat, lon int64, fontNumber int, textString string, hT *string) {
	c := ColorText{
		Color: Color{
			R: r,
			G: g,
			B: b,
		},
		Text: Text{
			Lat:        lat,
			Lon:        lon,
			FontNumber: fontNumber,
			TextString: textString,
			HoverText:  quotedStrPtr(hT),
		},
	}
	o.colortext = append(o.colortext, c)
}

func (o *Object) parseIconDefs() string {
	var res string
	for _, icon := range o.icon {
		res += icon.String()
	}
	return res
}

func (o *Object) parseColorText() string {
	var res string
	for _, ct := range o.colortext {
		res += fmt.Sprintf("%s\n%s\n", ct.Color.String(), ct.Text.String())
	}
	return res
}

func (o *Object) String() string {
	var res string
	res = fmt.Sprintf("Object: %.2f, %.2f\n", o.Lat, o.Lon)
	res += o.parseIconDefs()
	res += o.parseColorText()
	for _, t := range o.text {
		res += t.String()
	}
	res += "End:\n"
	return res
}

func (o *Object) MarshalYAML() (interface{}, error) {
	return o.String(), nil
}
