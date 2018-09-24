package mock

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3Client mocks the AWS API S3 endpoint.
type S3Client struct{ s3iface.S3API }

// ListBuckets returns an arry of mock AWS S3 buckets.
func (s *S3Client) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	n1, n2, n3 := "test_catvideos", "test_dank_memes", "test_go_executables"
	b1, b2, b3 := s3.Bucket{Name: &n1}, s3.Bucket{Name: &n2}, s3.Bucket{Name: &n3}
	output := s3.ListBucketsOutput{Buckets: []*s3.Bucket{&b1, &b2, &b3}}
	return &output, nil
}
