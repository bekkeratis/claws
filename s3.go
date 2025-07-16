package claws

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Returns an S3 service, to be reused by the client
func S3Client(config aws.Config) *s3.Client {
	return s3.NewFromConfig(config)
}

// Writes a byte array to S3 in the given bucket and filename
func S3Write(client *s3.Client, bucket string, filename string, content []byte) error {
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(content),
	})
	return err
}

// Returns the list of buckets
// Testing purposes only, this function does not use pagination
func Buckets(client *s3.Client) ([]string, error) {
	list, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return []string{}, err
	}
	buckets := make([]string, len(list.Buckets))
	for i, b := range list.Buckets {
		buckets[i] = *b.Name
	}
	return buckets, nil
}
