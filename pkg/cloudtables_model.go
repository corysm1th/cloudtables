package cloudtables

// DynamoDBTableObj represents a DynamoDB Table
type DynamoDBTableObj struct {
	Account string
	Region  string
	Name    string
}

// EC2Obj represents an EC2 instance.
type EC2Obj struct {
	ID               string
	Name             string
	Type             string
	Key              string
	PrivateIP        string
	PublicIP         string
	Account          string
	AvailabilityZone string
	Region           string
}

// EIP represents an AWS Elastic IP address.
type EIP struct {
	pubIP      string
	domain     string
	instanceID string
	awsAcct    string
	region     string
}

// S3Bucket represents an AWS S3 bucket
type S3Bucket struct {
	bucket  string
	awsAcct string
}
