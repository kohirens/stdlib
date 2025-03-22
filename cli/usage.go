package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type Usage struct {
	Command *Command
}
type Map map[string]string

type Command struct {
	Flags       *flag.FlagSet
	Name        string
	Summary     string
	Messages    Map
	Template    string
	Vars        Map
	Subcommands map[string]*Command
}

const defaultUsageTmpl = `
{{- define "optionHeader"}}
Options
{{end -}}

{{- define "option"}}
{{printf "  -%-11s %v" .OptionName .OptionInfo}}{{with .DefaultValue }} (default = {{.}}){{end}}
{{end -}}

{{- define "commandHeader"}}
Commands
{{end -}}

{{- define "subcommand"}}
  {{.Command}} - {{.Summary}}

    usage: {{.AppName}} -[global options] {{.Command}} -[options] <args>

    See {{.AppName}} {{.Command}} -help
{{end}}
Usage: {{.AppName}} -[options] <args>
`

var (
	appName              string
	printedCommandHeader = false
	printedOptionHeader  = false
)

// AddCommand Add additional application command usage information.
//
//	*flag.FlagSet.Usage function will be replaced with cli.Print.
func (u *Usage) AddCommand(flags *flag.FlagSet, name string, msgs, vars Map, summary, tmplStr string) *Usage {
	u.Command.AddCommand(flags, name, msgs, vars, summary, tmplStr)

	return u
}

// AddCommand Add additional application command usage information.
//
//	*flag.FlagSet.Usage function will be replaced with cli.Print.
func (c *Command) AddCommand(flags *flag.FlagSet, name string, msgs, vars Map, summary, tmplStr string) *Command {
	if flags == nil {
		panic("need non-nil *flag.FlagSet " + name)
	}

	if name == "" {
		panic("a command name cannot be an empty string")
	}

	v := Map{}
	if vars != nil {
		v = vars
	}

	c.Subcommands[name] = &Command{
		Flags:    flags,
		Name:     name,
		Summary:  summary,
		Messages: msgs,
		Template: tmplStr,
		Vars:     v,
	}

	flags.Usage = func() {
		if e := PrintUsage(c.Subcommands[name]); e != nil {
			panic(e)
		}
	}

	return c
}

// NewUsage Set up the application usage information.
//
//	flag.Usage function will be replaced with this cli.Print.
func NewUsage(name string, msgs, vars Map, summary, tmplStr string) *Usage {
	v := Map{}
	if vars != nil {
		v = vars
	}

	appName = name

	u := &Usage{
		Command: &Command{
			Name:        name,
			Flags:       flag.CommandLine,
			Summary:     summary,
			Messages:    msgs,
			Template:    tmplStr,
			Vars:        v,
			Subcommands: map[string]*Command{},
		},
	}

	flag.Usage = func() {
		if e := PrintUsage(u.Command); e != nil {
			panic(e)
		}
	}

	return u
}

// Print Display application usage information.
//
//	NOTE: Flags that do not have an entry in the messages list are hidden.
func (c *Command) Print() error {
	return PrintUsage(c)
}

// PrintUsage Display application usage information.
//
//	NOTE: Flags that do not have an entry in the messages list are hidden.
func PrintUsage(c *Command) error {
	tmplStr := c.Template
	if tmplStr == "" { // use the default when non provided
		tmplStr = defaultUsageTmpl
	}

	// setting these here gives the user the opportunity to change them
	// anytime before printing.
	c.Vars["AppName"] = appName
	if appName != c.Name {
		c.Vars["Command"] = c.Name
		c.Vars["Summary"] = c.Summary
	}

	tmpl, err1 := template.New(c.Name + "_usage").Parse(tmplStr)
	if err1 != nil {
		return fmt.Errorf(stderr.UsageTmplParse, err1.Error())
	}

	if e := tmpl.Execute(os.Stdout, c.Vars); e != nil {
		return e
	}

	if e := printOptionUsage(c, tmpl); e != nil {
		return e
	}

	var err2 error
	for _, sc := range c.Subcommands {
		if e := printSubCommandSummary(sc, tmpl); e != nil {
			err2 = e
			break
		}
	}

	return err2
}

// printCommandHeader Print the command header once per usage run.
func printCommandHeader(tmpl *template.Template, vars Map) error {
	if printedCommandHeader {
		return nil
	}

	cht := tmpl.Lookup("commandHeader")
	if cht == nil {
		return nil
	}

	if e := cht.Execute(os.Stdout, vars); e != nil {
		return e
	}

	printedCommandHeader = true

	return nil
}

func printSubCommandSummary(c *Command, tmpl *template.Template) error {
	// setting these here gives the user the opportunity to change them
	// anytime before printing.
	c.Vars["AppName"] = appName
	c.Vars["Command"] = c.Name
	c.Vars["Summary"] = c.Summary

	if e := printCommandHeader(tmpl, c.Vars); e != nil {
		return e
	}

	if ct := tmpl.Lookup("subcommand"); ct != nil {
		if e := ct.Execute(os.Stdout, c.Vars); e != nil { // command
			return e
		}
	}

	return nil
}

// printOptionHeader Print the option header once per usage run.
func printOptionHeader(tmpl *template.Template, vars Map) error {
	if printedOptionHeader {
		return nil
	}

	cht := tmpl.Lookup("optionHeader")
	if cht == nil {
		return nil
	}

	if e := cht.Execute(os.Stdout, vars); e != nil {
		return e
	}

	printedOptionHeader = true

	return nil
}

// printOptionUsage
//
//	NOTE: Flags that do not have an entry in the messages list are hidden.
func printOptionUsage(c *Command, parentTmpl *template.Template) error {
	ot := parentTmpl.Lookup("option")
	if ot == nil { // when there is no option template do nothing.
		return nil
	}

	numberOfFlags := 0
	msgs := c.Messages
	// We need to know there is at least 1 flag defined.
	// flag.Nflag only counts flags that are set and there seems no other way.
	// So we wastefully visit each flag to find out.

	c.Flags.VisitAll(func(f *flag.Flag) {
		if _, ok := msgs[f.Name]; ok {
			numberOfFlags++
		}

		return
	})

	if numberOfFlags < 1 {
		return nil
	}

	if e := printOptionHeader(parentTmpl, c.Vars); e != nil {
		return e
	}

	var err1 error
	var m string
	var ok bool
	c.Flags.VisitAll(func(f *flag.Flag) { // global flags
		m, ok = msgs[f.Name]

		if ok {
			c.Vars["OptionName"] = f.Name
			c.Vars["OptionInfo"] = m
			c.Vars["DefaultValue"] = f.Value.String()

			if e := ot.Execute(os.Stdout, c.Vars); e != nil {
				err1 = e
				return
			}
		}
	})

	return err1
}
