package fsio

var stderr = struct {
	HashFile,
	OpenFile string
}{
	HashFile: "problem generating a hash for file: %v",
	OpenFile: "problem while opening file: %v",
}

var stdout = struct {
}{}
