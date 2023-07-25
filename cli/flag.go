package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

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

var UsageTmpl = `
{{- define "option"}}
{{printf "  -%-11s %v" .option .info}}{{with .dv }} (default = {{.}}){{end}}
{{end}}
{{- define "subcommand"}}
{{printf "%-14s %v" .command .summary}}
{{end}}

Usage: {{ .appName }} [subcommand] [options] <args>

Options:
`

var UsageTmplVars = map[string]string{}

// Usage Print the usage documentation.
func Usage(appName string, um map[string]string, subcommands map[string]*flag.FlagSet) error {
	tmpl, err1 := template.New("Usage").Parse(UsageTmpl)
	if err1 != nil {
		return fmt.Errorf(stderr.UsageTmplParse, err1.Error())
	}

	uTmplData := map[string]string{
		"appName": appName,
	}

	if e := tmpl.Execute(os.Stdout, uTmplData); e != nil {
		return fmt.Errorf(stderr.UsageTmplExecute, e.Error())
	}

	var err2 error
	flag.VisitAll(func(f *flag.Flag) { // global flags
		m, ok := um[f.Name]
		if ok {
			UsageTmplVars["option"] = f.Name
			UsageTmplVars["info"] = m
			UsageTmplVars["dv"] = f.Value.String()

			if e := tmpl.ExecuteTemplate(os.Stdout, "option", UsageTmplVars); e != nil {
				err2 = e
				return
			}
		}
	})

	for command, flagSet := range subcommands {
		UsageTmplVars["command"] = command
		UsageTmplVars["summary"] = um[command]

		if e := tmpl.ExecuteTemplate(os.Stdout, "subcommand", UsageTmplVars); e != nil {
			err2 = e
			break
		}

		flagSet.VisitAll(func(f *flag.Flag) { // global flags
			m, ok := um[command+"_"+f.Name]
			if ok {
				UsageTmplVars["option"] = f.Name
				UsageTmplVars["info"] = m
				UsageTmplVars["dv"] = f.Value.String()

				if e := tmpl.ExecuteTemplate(os.Stdout, "option", UsageTmplVars); e != nil {
					err2 = e
					return
				}
			}
		})
		if err2 != nil {
			break
		}
	}

	return err2
}
