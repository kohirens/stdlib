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
	DownLoadKey        string
	EncodeJSON         string
	ReadObject         string
	InvalidObjectState string
	NoSuchKey          string
}{
	DownLoadKey:        "cannot download key %v from bucket %v: %v",
	EncodeJSON:         "could not encode JSON: %w",
	ReadObject:         "cannot read object key %v: %v",
	InvalidObjectState: "s3 invalid object state: %v",
	NoSuchKey:          "no such key %v",
}
