package sss

var Stdout = struct {
	ReadingObject string
	S3Download    string
	S3Move        string
	S3Upload      string
}{
	ReadingObject: "reading object %v",
	S3Download:    "will download file %v to memory",
	S3Move:        "will move file from %v to %v",
	S3Upload:      "will upload file %v to %v",
}

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
