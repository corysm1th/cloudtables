package cloudtables

// DynamoDBObj represents a DynamoDB Table
type DynamoDBObj struct {
	Account string
	Region  string
	Name    string
}

// EC2InstObj represents an EC2 instance.
type EC2InstObj struct {
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

// EC2EIPObj represents an AWS Elastic IP address.
type EC2EIPObj struct {
	PublicIP   string
	Domain     string
	InstanceID string
	Account    string
	Region     string
}

// S3BucketObj represents an AWS S3 bucket
type S3BucketObj struct {
	Name    string
	Account string
}
