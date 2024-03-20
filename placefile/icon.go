package placefile

import "fmt"

type IconFile struct {
	//fileNumber:
	//	This is a number from 1 to 8 used to reference the file in subsequent Icon statements. The number
	//	cannot be used in other IconFile statements.
	//
	//iconWidth and iconHeight:
	//	width and height of each icon in the file (pixels)
	//
	//hotX, hotY:
	//	zero-based center of the icon, used for placement and as the center of rotation. Must be less than
	//	iconWidth, iconHeight
	//
	//fileName:
	//	local, network, or URL filename (PNG, JPG, TIF supported) if no alpha channel is in the file, solid
	//	black (0,0,0) is made transparent
	FileNumber       int // TODO: 1-8
	IconWidthPixels  int
	IconHeightPixels int
	HotX             int
	HotY             int
	Filename         string
}

type Icon struct {
	//lat, lon:
	//coordinates where the icon hotspot is placed
	//
	//angle:
	//clockwise angle of rotation of the icon (degrees)
	//
	//fileNumber:
	//fileNumber from earlier IconFile statement
	//
	//iconNumber:
	//number of the icon to draw from 1 to the number of icons
	//in the icon file. Icons in an icon file are assigned by
	//row from left to right, top to bottom.
	//
	//quotedStrPtr (optional):
	//text displayed when the mouse cursor hovers over the
	//lat, lon of the icon
	Lat        float32
	Lon        float32
	Angle      int
	FileNumber int
	IconNumber int
	HoverText  string
}

// lat, lon, angle, fileNumber, iconNumber, quotedStrPtr
func (i Icon) String() string {
	// lat, lon, angle, fileNumber, iconNumber, quotedStrPtr
	x := fmt.Sprintf(
		"Icon: %.2f, %.2f, %d, %d, %d",
		i.Lat,
		i.Lon,
		i.Angle,
		i.FileNumber,
		i.IconNumber,
	)
	if i.HoverText != "" {
		x += "," + i.HoverText
	}
	return x + "\n"
}

func (ii IconFile) String() string {
	//	FileNumber       int // TODO: 1-8
	//	IconWidthPixels  int
	//	IconHeightPixels int
	//	HotX             int
	//	HotY             int
	//	Filename         string
	return fmt.Sprintf(
		"IconFile: %d, %d, %d, %d, %d, %s",
		ii.FileNumber,
		ii.IconWidthPixels,
		ii.IconHeightPixels,
		ii.HotX,
		ii.HotY,
		quotedStrPtr(&ii.Filename),
	)
}
