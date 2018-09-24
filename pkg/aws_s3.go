package cloudtables

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"
)

// GetAWSBuckets fetches AWS S3 buckets.
func GetAWSBuckets(svc s3iface.S3API, account string) (*[]S3BucketObj, int, error) {
	var buckets []S3BucketObj
	input := s3.ListBucketsInput{}
	obj := S3BucketObj{}
	obj.Account = account
	var count int
	resp, err := svc.ListBuckets(&input)
	if err != nil {
		return &buckets, count, errors.Wrap(err, "ListBuckets request failed.")
	}
	count += len(resp.Buckets)
	for _, b := range resp.Buckets {
		obj.Name = *b.Name
		buckets = append(buckets, obj)
	}
	return &buckets, count, nil
}
