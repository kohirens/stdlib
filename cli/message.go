package cli

var stderr = struct {
	UsageTmplParse   string
	UsageTmplExecute string
}{
	UsageTmplParse:   "error parsing the Usage template: %v",
	UsageTmplExecute: "error executing the Usage template %v",
}
