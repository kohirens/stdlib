package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type Command struct {
	Flags    *flag.FlagSet
	Name     string
	Summary  string
	Messages StringMap
}

// Message
// Deprecated: use StringMap
type Message = map[string]string

// StringMap use for usage messages, templates, and template vars.
type StringMap = map[string]string

const defaultUsageTmpl = `
{{- define "optionHeader"}}
Options:
{{end -}}

{{- define "option"}}
{{printf "  -%-11s %v" .OptionName .OptionInfo}}{{with .DefaultVal }} (default = {{.}}){{end}}
{{end -}}

{{- define "subcommand"}}
{{printf "%-14s %v" .Command .Summary}}

Usage: {{ .AppName }} [global options] {{ .Command }} [options] <args>
{{end}}

Usage: {{ .AppName }} [command] [options] <args>
`

var (
	command     *Command
	subcommands = []*Command{}

	// UsageTmpl Go template to used for display usage information.
	UsageTmpl = defaultUsageTmpl

	usageTmpls = StringMap{}

	// UsageTmplVars Variables used to fill in the actions in a template.
	// Deprecated: Use
	UsageTmplVars = StringMap{}
	usageVars     = StringMap{}
)

func AddCommand(name, summary string, um StringMap, flags *flag.FlagSet) *Command {
	c := &Command{
		flags,
		name,
		summary,
		um,
	}

	subcommands = append(subcommands, c)

	return c
}

// AddGlobalCommand Set up the application usage information.
func AddGlobalCommand(name, summary, tmplStr string, msgs, vars StringMap) *Command {
	command = &Command{
		flag.CommandLine,
		name,
		summary,
		msgs,
	}

	usageVars["AppName"] = name

	AddTmpl(name, tmplStr, vars)

	return command
}

// AddTmpl Set a custom template for each command.
func AddTmpl(command, tmplStr string, vars StringMap) {
	usageTmpls[command] = tmplStr

	for k, v := range vars {
		usageVars[k] = v
	}
}

// Usage Print the usage documentation.
//
//	NOTE: Flags that do not have an entry in the um (usage message) list are
//	hidden.
func Usage(appName string, um StringMap, subcommands map[string]*flag.FlagSet) error {
	tmpl, err1 := template.New("Usage").Parse(UsageTmpl)
	if err1 != nil {
		return fmt.Errorf(stderr.UsageTmplParse, err1.Error())
	}

	UsageTmplVars["AppName"] = appName
	oht := tmpl.Lookup("optionHeader")
	ot := tmpl.Lookup("option")

	if e := tmpl.Execute(os.Stdout, UsageTmplVars); e != nil {
		return fmt.Errorf(stderr.UsageTmplExecute, e.Error())
	}

	var err2 error
	if ot != nil {
		if e := printOptionUsage(ot, oht, flag.CommandLine, um, ""); e != nil {
			return e
		}
	}

	if sct := tmpl.Lookup("subcommand"); sct != nil {
		for c, flagSet := range subcommands {
			UsageTmplVars["Command"] = c
			UsageTmplVars["Summary"] = um[c]

			if e := tmpl.ExecuteTemplate(os.Stdout, "subcommand", UsageTmplVars); e != nil {
				err2 = e
				break
			}
			if ot != nil {
				if e := printOptionUsage(ot, oht, flagSet, um, c); e != nil {
					err2 = e
					break
				}
			}
		}
	}

	return err2
}

// UsageV2 Print the usage documentation.
//
//	NOTE: Flags that do not have an entry in the um (usage message) list are
//	hidden.
func UsageV2(name string) error {
	if len(subcommands) < 1 {
		return nil
	}

	tmplStr, ok := usageTmpls[name]

	if !ok || tmplStr == "" { // use the default when non provided
		tmplStr = defaultUsageTmpl
	}

	if e := printUsage(name, tmplStr, usageVars); e != nil {
		return e
	}

	return nil
}

func printUsage(tmplName, tmplStr string, vars StringMap) error {
	tmpl, err1 := template.New(tmplName + "_usage").Parse(tmplStr)
	if err1 != nil {
		return fmt.Errorf(stderr.UsageTmplParse, err1.Error())
	}

	if e := tmpl.Execute(os.Stdout, vars); e != nil {
		return e
	}

	if len(subcommands) > 0 { // display an option command header
		cht := tmpl.Lookup("commandHeader")
		if e := cht.Execute(os.Stdout, vars); e != nil {
			return e
		}
	}

	var err2 error
	for _, c := range subcommands {
		if e := printSubcommandUsage(c, tmpl, vars); e != nil {
			err2 = e
			break
		}
	}

	return err2
}

func printSubcommandUsage(c *Command, tmpl *template.Template, vars StringMap) error {
	vars["Command"] = c.Name
	vars["Summary"] = c.Summary

	if ct := tmpl.Lookup("subcommand"); ct != nil {
		if e := ct.Execute(os.Stdout, vars); e != nil { // command
			return e
		}
	}

	if ot := tmpl.Lookup("option"); ot != nil {
		oht := tmpl.Lookup("optionHeader")
		if e := printOptionUsage(ot, oht, command.Flags, command.Messages, command.Name); e != nil {
			return e
		}
	}

	return nil
}

func printOptionUsage(tmpl, oht *template.Template, flags *flag.FlagSet, um StringMap, command string) error {
	// We need to know there is at least 1 flag defined.
	// flag.Nflag only counts flags that are set and there seems no other way.
	// So we wastefully visit each flag to find out.
	numberOfFlags := 0
	flags.VisitAll(func(f *flag.Flag) {
		_, ok := um[f.Name]
		if command != "" {
			_, ok = um[command+"_"+f.Name]
		}

		if ok {
			numberOfFlags++
		}

		return
	})

	if numberOfFlags < 1 {
		return nil
	}

	if oht != nil {
		if e := oht.ExecuteTemplate(os.Stdout, "optionHeader", usageVars); e != nil {
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
			usageVars["OptionName"] = f.Name
			usageVars["OptionInfo"] = m
			usageVars["DefaultVal"] = f.Value.String()

			if e := tmpl.ExecuteTemplate(os.Stdout, "option", usageVars); e != nil {
				err1 = e
				return
			}
		}
	})

	return err1
}
