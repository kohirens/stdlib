package fsio

var stderr = struct {
	HashFile,
	OpenFile,
	Read string
}{
	HashFile: "problem generating a hash for file: %v",
	OpenFile: "problem while opening file: %v",
	Read:     "could not read content: %v",
}
