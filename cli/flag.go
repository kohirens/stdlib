package cli

type BoolFlag struct {
	Name string
	Def  string
	Um   string
	Bool string
}

type IntFlag struct {
	name string
	def  string
	um   string
	int  string
}

type StrFlag struct {
	Name string
	Def  string
	Um   string
	Val  string
}

const UsageTmpl = `
{{define "option"}}
{{printf "  -%-11s %v" .option .info}}{{with .dv }} (default = {{.}}){{end}}
{{end}}
Usage: {{ .appName }} [subcommand] -[options] <args>

Options:
`
