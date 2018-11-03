package cloudtables

// QueryFilter maps struct properties to query parameters.
type QueryFilter struct {
	Key   string
	Value string
}

// Storage is the parent interface for the storage backend.
type Storage interface {
	DynamoDBObjHandler
	EC2EIPObjHandler
	EC2InstObjHandler
	S3BucketObjHandler
}

// DynamoDBObjHandler bridges DynamoDBObjs to a storage backend.
type DynamoDBObjHandler interface {
	// StoreDynamoDBObj stores an array of DynamoDB Tables in the database.
	StoreDynamoDBObj([]DynamoDBObj) error
	// SelectDynamoDBObj retrieves an array of DynamoDB Tables from the database.
	SelectDynamoDBObj() ([]*DynamoDBObj, error)
	// DeleteDynamoDBObjs removes an array of DynamoDB Tables from the database.
	DeleteDynamoDBObjs() error
}

// EC2EIPObjHandler bridges EC2EIPObjs to a storage backend.
type EC2EIPObjHandler interface {
	// StoreEC2EIPObj stores an array of Elastic IPs in the database.
	StoreEC2EIPObj([]EC2EIPObj) error
	// SelectEC2EIPObj retrieves an array of Elastic IPs from the database.
	SelectEC2EIPObj() ([]*EC2EIPObj, error)
	// DeleteEC2EIPObjs removes an array of Elastic IPs from the database.
	DeleteEC2EIPObjs() error
}

// EC2InstObjHandler bridges EC2InstObjs to a storage backend.
type EC2InstObjHandler interface {
	// StoreEC2InstObj stores an array of EC2 Instances in the database.
	StoreEC2InstObj([]EC2InstObj) error
	// SelectEC2InstObj retrieves an array of EC2 Instances from the database.
	SelectEC2InstObj() ([]*EC2InstObj, error)
	// DeleteEC2InstObjs removes an array of EC2 Instances from the database.
	DeleteEC2InstObjs() error
}

// S3BucketObjHandler bridges S3BucketObjs to a storage backend.
type S3BucketObjHandler interface {
	// StoreS3BucketObj stores an array of S3 Buckets in the database.
	StoreS3BucketObj([]S3BucketObj) error
	// SelectS3BucketObj retrieves an array of S3 Buckets from the database.
	SelectS3BucketObj() ([]*S3BucketObj, error)
	// DeleteS3BucketObjs removes an array of S3 Buckets from the database.
	DeleteS3BucketObjs() error
}
