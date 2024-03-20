package placefile

import (
	"fmt"
	"testing"
)

func TestPlacefileFont_String(t *testing.T) {
	type fields struct {
		FontNumber int
		Pixels     int
		Flags      FontAttributeFlag
		Face       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Font validation",
			fields{
				FontNumber: 1,
				Pixels:     12,
				Flags:      1,
				Face:       "example",
			},
			fmt.Sprintf("%d, %d, %d, \"example\"", 1, 12, 1),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pf := PlacefileFont{
				FontNumber: tt.fields.FontNumber,
				Pixels:     tt.fields.Pixels,
				Flags:      tt.fields.Flags,
				Face:       tt.fields.Face,
			}
			g := pf.String()
			if g != tt.want {
				t.Errorf("String() = %s want %v", g, tt.want)
			}
		})
	}
}
