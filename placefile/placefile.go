package placefile

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type IconFiles []IconFile
type Objects []Object

func (iconFiles IconFiles) String() string {
	var res string
	for _, i := range iconFiles {
		res += i.String() + "\n"
	}
	return res
}

func (objects Objects) String() string {
	var res string
	for _, o := range objects {
		res += "\n" + o.String()
	}
	return res
}

type Placefile struct {
	internalPlacefile
	iconFiles   IconFiles `yaml:"-"`
	objects     Objects   `yaml:"-"`
	topComments []string  `yaml:"-"`
}

type PlacefileInput struct {
	Refresh   int
	Title     string
	Font      PlacefileFont
	Threshold int
}

func NewPlacefile(input *PlacefileInput) *Placefile {
	return &Placefile{
		internalPlacefile: internalPlacefile{
			Refresh:   input.Refresh,
			Title:     input.Title,
			Font:      input.Font,
			Threshold: input.Threshold,
		},
		iconFiles: nil,
		objects:   nil,
	}
}

type internalPlacefile struct {
	Refresh   int           `yaml:"Refresh"`
	Title     string        `yaml:"Title"`
	Font      PlacefileFont `yaml:"Font"`
	Threshold int           `yaml:"Threshold"`

	//IconFile  IconFile      `yaml:"IconFile"`
}

func (p *Placefile) AddIconFile(fileNumber, iconWidthPixels, iconHeightPixels, hotX, hotY int, filename string) {
	i := IconFile{
		FileNumber:       fileNumber,
		IconWidthPixels:  iconWidthPixels,
		IconHeightPixels: iconHeightPixels,
		HotX:             hotX,
		HotY:             hotY,
		Filename:         filename,
	}
	p.iconFiles = append(p.iconFiles, i)
}

func (p *Placefile) AddObject(o Object) {
	p.objects = append(p.objects, o)
}

func (p *Placefile) AddTopOfFileComment(s string) {
	p.topComments = append(p.topComments, s)
}

func (p *Placefile) MarshalYAML() (interface{}, error) {
	//spaceRegex, err := regexp.Compile("\x20\x20\x20\x20")
	//if err != nil {
	//	panic(err)
	//}
	var b []byte
	for _, c := range p.topComments {
		b = append(b, []byte(fmt.Sprintf("; %s\n", c))...)
	}
	bb, err := yaml.Marshal(p.internalPlacefile)
	b = append(b, bb...)

	if err != nil {
		return nil, err
	}
	//b = bytes.TrimSpace(b[0:])
	s := string(b[0:])
	if len(p.iconFiles) > 0 {
		s += p.iconFiles.String()
	}
	if len(p.objects) > 0 {
		s += p.objects.String()
	}
	return s, nil
}

func (p *Placefile) ToString() (string, error) {
	s, err := p.MarshalYAML()
	if err != nil {
		return "", err
	}
	return strings.Replace(s.(string), "\x20\x20\x20\x20", "", -1), nil
}
