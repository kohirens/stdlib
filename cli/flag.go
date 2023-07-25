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
