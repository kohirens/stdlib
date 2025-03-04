package sss

var Stdout = struct {
}{}

var Stderr = struct {
	DecodeJSON         string
	DownLoadKey        string
	EncodeJSON         string
	ReadObject         string
	InvalidObjectState string
	NoSuchKey          string
}{
	DecodeJSON:         "could not decode JSON: %v",
	DownLoadKey:        "cannot download key %v from bucket %v: %v",
	EncodeJSON:         "could not encode JSON: %v",
	ReadObject:         "cannot read object key %v: %v",
	InvalidObjectState: "s3 invalid object state: %v",
	NoSuchKey:          "no such key %v",
}
