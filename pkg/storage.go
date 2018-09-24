package cloudtables

// QueryFilter maps struct properties to query parameters.
type QueryFilter struct {
	Key   string
	Value string
}

// DynamoDBObjHandler bridges DynamoDBObjs to a storage backend.
type DynamoDBObjHandler interface {
	// StoreDynamoDBObj stores an array of DynamoDB Tables in the database.
	StoreDynamoDBObj(*[]DynamoDBObj) error
	// SelectDynamoDBObj retrieves an array of DynamoDB Tables from the database.
	SelectDynamoDBObj(*[]QueryFilter) error
	// DeleteDynamoDBObj removes an array of DynamoDB Tables from the database.
	DeleteDynamoDBObj(*[]DynamoDBObj) error
}

// EC2EIPObjHandler bridges EC2EIPObjs to a storage backend.
type EC2EIPObjHandler interface {
	// StoreEC2EIPObj stores an array of Elastic IPs in the database.
	StoreEC2EIPObj(*[]EC2EIPObj) error
	// SelectEC2EIPObj retrieves an array of Elastic IPs from the database.
	SelectEC2EIPObj(*[]QueryFilter) error
	// DeleteEC2EIPObj removes an array of Elastic IPs from the database.
	DeleteEC2EIPObj(*[]EC2EIPObj) error
}

// EC2InstObjHandler bridges EC2InstObjs to a storage backend.
type EC2InstObjHandler interface {
	// StoreEC2InstObj stores an array of EC2 Instances in the database.
	StoreEC2InstObj(*[]EC2InstObj) error
	// SelectEC2InstObj retrieves an array of EC2 Instances from the database.
	SelectEC2InstObj(*[]QueryFilter) error
	// DeleteEC2InstObj removes an array of EC2 Instances from the database.
	DeleteEC2InstObj(*[]EC2InstObj) error
}

// S3BucketObjHandler bridges S3BucketObjs to a storage backend.
type S3BucketObjHandler interface {
	// StoreEC2 stores an array of S3 Buckets in the database.
	StoreS3BucketObj(*[]S3BucketObj) error
	// SelectObject retrieves an array of S3 Buckets from the database.
	SelectS3BucketObj(*[]QueryFilter) error
	// DeleteObject removes an array of S3 Buckets from the database.
	DeleteS3BucketObj(*[]S3BucketObj) error
}
