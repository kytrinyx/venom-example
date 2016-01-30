package venom

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		desc     string
		args     []string
		expected string
	}{
		{"default file", []string{}, "root.example.com"},
		{"override with file", []string{"--file=fixtures/config.yaml"}, "fixture.example.com"},
		{"override with flag", []string{"--url=flag.example.com"}, "flag.example.com"},
	}

	args := os.Args
	defer func() {
		os.Args = args
	}()

	for _, test := range tests {
		os.Args = append(args, test.args...)

		c, err := NewConfig()
		if err != nil {
			t.Fatal(err)
		}

		if c.URL != test.expected {
			t.Errorf("case %s: got: %s, want: %s", test.desc, c.URL, test.expected)
		}
	}
}
