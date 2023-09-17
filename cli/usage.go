package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type Messages = map[string]string

var UsageTmpl = `
{{- define "optionHeader"}}
options:
{{end -}}

{{- define "option"}}
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
//
//	NOTE: Flags that do not have an entry in the um (usage message) list are
//	hidden.
func Usage(appName string, um map[string]string, subcommands map[string]*flag.FlagSet) error {
	tmpl, err1 := template.New("Usage").Parse(UsageTmpl)
	if err1 != nil {
		return fmt.Errorf(stderr.UsageTmplParse, err1.Error())
	}

	UsageTmplVars["AppName"] = appName
	oht := tmpl.Lookup("optionHeader")

	if e := tmpl.Execute(os.Stdout, UsageTmplVars); e != nil {
		return fmt.Errorf(stderr.UsageTmplExecute, e.Error())
	}

	var err2 error
	if e := executeOptionTmpl(tmpl, oht, flag.CommandLine, um, ""); e != nil {
		return e
	}

	if sct := tmpl.Lookup("subcommand"); sct != nil {
		for command, flagSet := range subcommands {
			UsageTmplVars["Command"] = command
			UsageTmplVars["Summary"] = um[command]

			if e := tmpl.ExecuteTemplate(os.Stdout, "subcommand", UsageTmplVars); e != nil {
				err2 = e
				break
			}

			if e := executeOptionTmpl(tmpl, oht, flagSet, um, command); e != nil {
				err2 = e
				break
			}
		}
	}

	return err2
}

func executeOptionTmpl(tmpl, oht *template.Template, flags *flag.FlagSet, um Messages, command string) error {
	// We need to know there is at least 1 flag defined.
	// flag.Nflag only counts flags that are set and there seems no other way.
	// So we wastefully visit each flag to find out.
	numberOfFlags := 0
	flags.VisitAll(func(f *flag.Flag) {
		numberOfFlags++
		return
	})

	if numberOfFlags < 1 {
		return nil
	}

	if oht != nil {
		if e := oht.ExecuteTemplate(os.Stdout, "optionHeader", UsageTmplVars); e != nil {
			return e
		}
	}

	var err1 error
	var m string
	var ok bool
	flags.VisitAll(func(f *flag.Flag) { // global flags
		if command != "" {
			m, ok = um[command+"_"+f.Name]
		} else {
			m, ok = um[f.Name]
		}

		if ok {
			UsageTmplVars["OptionName"] = f.Name
			UsageTmplVars["OptionInfo"] = m
			UsageTmplVars["DefaultVal"] = f.Value.String()

			if e := tmpl.ExecuteTemplate(os.Stdout, "option", UsageTmplVars); e != nil {
				err1 = e
				return
			}
		}
	})

	return err1
}
