package sss

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kohirens/stdlib/logger"
	wSession "github.com/kohirens/stdlib/web/session"
	"io"
)

type StorageBucket struct {
	Context context.Context
	Name    string
	S3      *s3.S3
	prefix  string
}

var Log = logger.StdLogger{}

// Load Session data from S3, the ID serves as the object name. We recommend the
// site domain be used as a prefix to prevent collision in the bucket.
// The timeout on the Context will interrupt the request if it expires.
// See also https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#example_S3_GetObject_shared00
func (c *StorageBucket) Load(key string) (*wSession.Data, error) {
	fullKey := c.prefix + key

	Log.Infof("Loading session for key %v", fullKey)

	obj, e1 := c.S3.GetObject(&s3.GetObjectInput{
		Bucket: &c.Name,
		Key:    &fullKey,
	})

	if e1 != nil {
		e := decipherError(e1)
		return nil, fmt.Errorf(Stderr.DownLoadKey, key, c.Name, e.Error())
	}

	b, e2 := io.ReadAll(obj.Body)
	if e2 != nil {
		return nil, fmt.Errorf(Stderr.ReadObject, key)
	}

	data := &wSession.Data{}
	if e := json.Unmarshal(b, data); e != nil {
		return nil, fmt.Errorf(Stderr.DecodeJSON, key)
	}
	return data, nil
}

// Save Session data to S3.
func (c *StorageBucket) Save(data *wSession.Data) error {
	content, e1 := json.Marshal(data)
	if e1 != nil {
		return fmt.Errorf(Stderr.EncodeJSON, e1)
	}

	_, e := c.Upload(content, data.Id)
	if e != nil {
		return e
	}

	return nil
}

// Upload Uploads an object to S3, returning the eTag on success. The Context
// will interrupt the request if the timeout expires.
// For more info, see https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#example_S3_PutObject_shared00
func (c *StorageBucket) Upload(b []byte, key string) (string, error) {
	fullKey := c.prefix + key

	Log.Infof("Saving data for key %v", fullKey)

	put, e1 := c.S3.PutObjectWithContext(c.Context, &s3.PutObjectInput{
		Bucket:               &c.Name,
		Key:                  &fullKey,
		Body:                 aws.ReadSeekCloser(bytes.NewReader(b)), //bytes.NewReader(b),
		ServerSideEncryption: aws.String("AES256"),
	})

	if e1 != nil {
		return "", decipherError(e1)
	}

	return *put.ETag, nil
}

// DecipherError Put an S3 error into context or something more human comprehensible.
func decipherError(e error) error {
	var err awserr.Error

	ok := errors.As(e, &err)

	if ok {
		switch err.Code() {
		case s3.ErrCodeNoSuchKey:
			return fmt.Errorf(Stderr.NoSuchKey, err.Error())
		case s3.ErrCodeInvalidObjectState:
			return fmt.Errorf(Stderr.InvalidObjectState, err.Error())
		}
	}

	return e
}

// NewStorageClient Initializes an S3 client to use as session storage.
// Credentials are expected to be configured in the environment to be picked up
// by the AWS SDK. Panics on failure.
func NewStorageClient(bucket string, ctx context.Context) *StorageBucket {
	sess := session.Must(session.NewSession())
	return &StorageBucket{
		Name:    bucket,
		S3:      s3.New(sess),
		Context: ctx,
	}
}

// Prefix Set a prefix for the bucket to prepend to keys before downloaded or uploading objects.
func (c *StorageBucket) Prefix(prefix string) *StorageBucket {
	c.prefix = prefix
	return c
}
