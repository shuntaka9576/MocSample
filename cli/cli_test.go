package cli_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"fmt"
	"github.com/shuntaka9576/MocSample/cli"
)

func TestCli_Run(t *testing.T) {
	t.Helper()
	var tests = []struct {
		pattern  string
		name     string
		args     []string
		expected string
	}{
		{"normal", "case", []string{"imageConverter", "-f", "png", "-t", "png", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "png", "-t", "jpg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "png", "-t", "jpeg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "png", "-t", "gif", "../testdata"}, ""},

		{"normal", "case", []string{"imageConverter", "-f", "jpg", "-t", "jpg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpg", "-t", "png", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpg", "-t", "jpeg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpg", "-t", "gif", "../testdata"}, ""},

		{"normal", "case", []string{"imageConverter", "-f", "jpeg", "-t", "jpeg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpeg", "-t", "png", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpeg", "-t", "jpg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "jpeg", "-t", "gif", "../testdata"}, ""},

		{"normal", "case", []string{"imageConverter", "-f", "gif", "-t", "gif", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "gif", "-t", "png", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "gif", "-t", "jpg", "../testdata"}, ""},
		{"normal", "case", []string{"imageConverter", "-f", "gif", "-t", "jpeg", "../testdata"}, ""},
	}

	for _, tt := range tests {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		c := &cli.Cli{OutStream: outStream, ErrStream: errStream}
		if tt.pattern == "normal" {
			if tt.pattern == "normal" {
				t.Run(tt.name, func(t *testing.T) {
					err := c.Run(tt.args)
					fmt.Println(c.OutStream)
					if err != 0 {
						t.Error("Faild case")
					}
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
		time.Sleep(1 * time.Second)
	}
}

