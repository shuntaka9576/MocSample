package converter_test

import (
	"testing"

	"github.com/shuntaka9576/MocSample/converter"
	_ "github.com/shuntaka9576/MocSample/imagetypes/gif"
	_ "github.com/shuntaka9576/MocSample/imagetypes/jpg"
	_ "github.com/shuntaka9576/MocSample/imagetypes/png"
)

func TestGetConverter(t *testing.T) {
	type input struct {
		from, to string
	}
	tests := []struct {
		pattern string
		name    string
		input   input
	}{
		{"normal", "png to jpg", input{"png", "jpg"}},
		{"normal", "png to jpeg", input{"png", "jpeg"}},
		{"normal", "png to gif", input{"png", "gif"}},
		{"normal", "gif to jpg", input{"gif", "jpg"}},
		{"normal", "gif to jpeg", input{"gif", "jpeg"}},
		{"normal", "gif to png", input{"gif", "png"}},
		{"normal", "jpg to png", input{"jpg", "png"}},
		{"normal", "jpeg to png", input{"jpeg", "png"}},
		{"normal", "jpg to gif", input{"jpg", "gif"}},
		{"normal", "jpeg to gif", input{"jpeg", "gif"}},
		{"non-normal", "fail from argument", input{"jp", "gif"}},
		{"non-normal", "fail to argument", input{"jpg", "gi"}},
		{"non-normal", "fail all argument", input{"jp", "gi"}},
	}
	for _, tt := range tests {
		if tt.pattern == "normal" {
			t.Run(tt.name, func(t *testing.T) {
				_, err := converter.GetConverter(tt.input.from, tt.input.to)
				if err != nil {
					t.Errorf("but got %v", err.Error())
				}
			})
		}
		if tt.pattern == "non-normal" {
			t.Run(tt.name, func(t *testing.T) {
				_, err := converter.GetConverter(tt.input.from, tt.input.to)
				if err == nil {
					t.Error("Test Fail")
				}
			})
		}
	}
}
