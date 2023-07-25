package cli

import (
	"bytes"
	"flag"
	"io"
	"os"
	"testing"
)

func TestUsage(t *testing.T) {
	var help bool
	var opt1 string
	flag.BoolVar(&help, "help", false, "")
	fixedSubCmd := flag.NewFlagSet("test-cmd", flag.ExitOnError)
	fixedSubCmd.StringVar(&opt1, "opt1", "", "opt1 info")
	flag.Parse()
	fixedSubCmd.Parse(flag.Args())

	tests := []struct {
		name       string
		um         map[string]string
		subcommand map[string]*flag.FlagSet
		want       string
	}{
		{
			"display usage message",
			map[string]string{
				"help":          "display this help",
				"test-cmd":      "test-cmd summary",
				"test-cmd_opt1": "opt1 summary",
			},
			map[string]*flag.FlagSet{
				"test-cmd": fixedSubCmd,
			},
			`

Usage: tester [command] [options] <args>

options:

  -help        display this help (default = false)

test-cmd       test-cmd summary

usage: tester [global options] test-cmd [options] <args>

options:

  -opt1        opt1 summary
`,
		},
	}

	oldStdout := os.Stdout

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w
			gotErr := Usage("tester", tt.um, tt.subcommand)

			if gotErr != nil {
				t.Errorf("Usage() error %v, want nil", gotErr.Error())
			}
			outC := make(chan string)
			// copy the output in a separate goroutine so printing can't block indefinitely
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()
			w.Close()
			os.Stdout = oldStdout // restoring the real stdout
			got := <-outC
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
	defer func() { os.Stdout = oldStdout }()
}