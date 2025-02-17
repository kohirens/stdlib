package fsio

var stderr = struct {
	HashFile,
	OpenFile string
}{
	HashFile: "problem generating a hash for file: %v",
	OpenFile: "problem while opening file: %v",
}

var Stdout = struct {
	BytesRead            string
	ConnectTo            string
	DomainOnRedirectList string
	EnvVarEmpty          string
	LoadPage             string
	RunCli               string
}{
	BytesRead:            "number of bytes read from %v is %d",
	ConnectTo:            "connecting to %v",
	DomainOnRedirectList: "domain %v in in the list of domains to redirect to %v",
	EnvVarEmpty:          "environment variable %v is empty",
	LoadPage:             "loading the %v page",
	RunCli:               "Running CLI",
}

var Stderr = struct {
	DecodeJSON string
	EncodeJSON string
	OpenFile   string
	ReadFile   string
	WriteFile  string
}{
	DecodeJSON: "could not decode JSON from file %v: %w",
	EncodeJSON: "could not encode JSON: %w",
	OpenFile:   "could not open file %v: %w",
	ReadFile:   "could not read file %v: %w",
	WriteFile:  "could not write content to file %v: %w",
}
