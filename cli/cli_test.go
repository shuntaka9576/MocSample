package cli_test

import (
	"bytes"
	"testing"

	"strings"

	"github.com/shuntaka9576/MocSample/cli"
)

func TestCli_Run(t *testing.T) {
	var tests = []struct {
		pattern  string
		name     string
		args     []string
		expected string
	}{
		//{"normal", []string{"imageConverter", "-version"}, fmt.Sprintf("imageConverter version %s", cli.Version)},
		{"normal", "case", []string{"imageConverter", "-f", "png", "-t", "jpg", "."}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpg", "-t", "png", "."}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpg", "-t", "png", "./testdata"}, ""},
	}

	for _, tt := range tests {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		c := cli.NewApp(outStream, errStream)
		if tt.pattern == "normal" {
			if tt.pattern == "normal" {
				t.Run(tt.name, func(t *testing.T) {
					c.Run(tt.args)
					if !strings.Contains(outStream.String(), tt.expected) {
						t.Errorf("Output=%q, want %q", outStream.String(), tt.expected)
					}
				})
			}
		}
		if tt.pattern == "non-normal" {
			t.Run(tt.name, func(t *testing.T) {
				c.Run(tt.args)
			})
		}
	}
}
