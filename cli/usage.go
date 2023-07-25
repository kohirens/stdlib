package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var UsageTmpl = `
{{- define "option"}}
options:

{{printf "  -%-11s %v" .OptionName .OptionInfo}}{{with .DefaultVal }} (default = {{.}}){{end}}
{{end -}}
{{- define "subcommand"}}
{{printf "%-14s %v" .Command .Summary}}

usage: {{ .AppName }} [global options] {{ .Command }} [options] <args>
{{end}}

Usage: {{ .AppName }} [command] [options] <args>
`

var UsageTmplVars = map[string]string{}

// Usage Print the usage documentation.
func Usage(appName string, um map[string]string, subcommands map[string]*flag.FlagSet) error {
	tmpl, err1 := template.New("Usage").Parse(UsageTmpl)
	if err1 != nil {
		return fmt.Errorf(stderr.UsageTmplParse, err1.Error())
	}
	UsageTmplVars["appName"] = appName
	UsageTmplVars["AppName"] = appName

	if e := tmpl.Execute(os.Stdout, UsageTmplVars); e != nil {
		return fmt.Errorf(stderr.UsageTmplExecute, e.Error())
	}

	var err2 error
	flag.VisitAll(func(f *flag.Flag) { // global flags
		m, ok := um[f.Name]
		if ok {
			UsageTmplVars["option"] = f.Name
			UsageTmplVars["OptionName"] = f.Name
			UsageTmplVars["info"] = m
			UsageTmplVars["OptionInfo"] = m
			UsageTmplVars["dv"] = f.Value.String()
			UsageTmplVars["DefaultVal"] = f.Value.String()

			if e := tmpl.ExecuteTemplate(os.Stdout, "option", UsageTmplVars); e != nil {
				err2 = e
				return
			}
		}
	})

	for command, flagSet := range subcommands {
		UsageTmplVars["command"] = command
		UsageTmplVars["Command"] = command
		UsageTmplVars["summary"] = um[command]
		UsageTmplVars["Summary"] = um[command]

		if e := tmpl.ExecuteTemplate(os.Stdout, "subcommand", UsageTmplVars); e != nil {
			err2 = e
			break
		}

		flagSet.VisitAll(func(f *flag.Flag) { // global flags
			m, ok := um[command+"_"+f.Name]
			if ok {
				UsageTmplVars["option"] = f.Name
				UsageTmplVars["OptionName"] = f.Name
				UsageTmplVars["info"] = m
				UsageTmplVars["OptionInfo"] = m
				UsageTmplVars["dv"] = f.Value.String()
				UsageTmplVars["DefaultVal"] = f.Value.String()

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
