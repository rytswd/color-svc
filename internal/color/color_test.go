package color

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnabledColors(t *testing.T) {
	cases := map[string]struct {
		red    bool
		green  bool
		blue   bool
		yellow bool
		want   []Color
	}{
		"red": {
			red:    true,
			green:  false,
			blue:   false,
			yellow: false,
			want: []Color{
				{Name: "Red", HexCode: "#ff0000"},
				{Name: "Maroon", HexCode: "#800000"},
				{Name: "Purple", HexCode: "#800080"},
				{Name: "Fuchsia", HexCode: "#ff00ff"},
			},
		},
		"green": {
			red:    false,
			green:  true,
			blue:   false,
			yellow: false,
			want: []Color{
				{Name: "Green", HexCode: "#008000"},
				{Name: "Lime", HexCode: "#00ff00"},
				{Name: "Aquamarine", HexCode: "#7fffd4"},
			},
		},
		"blue": {
			red:    false,
			green:  false,
			blue:   true,
			yellow: false,
			want: []Color{
				{Name: "Blue", HexCode: "#0000ff"},
				{Name: "Navy", HexCode: "#000080"},
				{Name: "Teal", HexCode: "#008080"},
				{Name: "Aqua", HexCode: "#00ffff"},
				{Name: "Aliceblue", HexCode: "#f0f8ff"},
				{Name: "Azure", HexCode: "#f0ffff"},
				{Name: "BlueViolet", HexCode: "#8a2be2"},
			},
		},
		"yellow": {
			red:    false,
			green:  false,
			blue:   false,
			yellow: true,
			want: []Color{
				{Name: "Yellow", HexCode: "#ffff00"},
				{Name: "Olive", HexCode: "#808000"},
				{Name: "Orange", HexCode: "#ffa500"},
				{Name: "Beige", HexCode: "#f5f5dc"},
				{Name: "Bisque", HexCode: "#ffe4c4"},
				{Name: "BlanchedAlmond", HexCode: "#ffebcd"},
			},
		},
		"all": {
			red:    true,
			green:  true,
			blue:   true,
			yellow: true,
			want: []Color{
				{Name: "Red", HexCode: "#ff0000"},
				{Name: "Maroon", HexCode: "#800000"},
				{Name: "Purple", HexCode: "#800080"},
				{Name: "Fuchsia", HexCode: "#ff00ff"},
				{Name: "Green", HexCode: "#008000"},
				{Name: "Lime", HexCode: "#00ff00"},
				{Name: "Aquamarine", HexCode: "#7fffd4"},
				{Name: "Blue", HexCode: "#0000ff"},
				{Name: "Navy", HexCode: "#000080"},
				{Name: "Teal", HexCode: "#008080"},
				{Name: "Aqua", HexCode: "#00ffff"},
				{Name: "Aliceblue", HexCode: "#f0f8ff"},
				{Name: "Azure", HexCode: "#f0ffff"},
				{Name: "BlueViolet", HexCode: "#8a2be2"},
				{Name: "Yellow", HexCode: "#ffff00"},
				{Name: "Olive", HexCode: "#808000"},
				{Name: "Orange", HexCode: "#ffa500"},
				{Name: "Beige", HexCode: "#f5f5dc"},
				{Name: "Bisque", HexCode: "#ffe4c4"},
				{Name: "BlanchedAlmond", HexCode: "#ffebcd"},
			},
		},
		"none": {
			red:    false,
			green:  false,
			blue:   false,
			yellow: false,
			want:   []Color{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := EnabledColors(tc.red, tc.green, tc.blue, tc.yellow)
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("unexpected result: (-want / +got)\n%s", diff)
			}
		})
	}
}
