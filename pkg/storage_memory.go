package cloudtables

// StorageMem implements the Storage interface to persist cloud objects in-memory
type StorageMem struct {
	DynamoDBs    []DynamoDBObj
	EC2IPs       []EC2EIPObj
	EC2Instances []EC2InstObj
	S3Buckets    []S3BucketObj
}

// NewStorageMem returns a new instance of in-memory storage
func NewStorageMem() *StorageMem {
	s := StorageMem{
		DynamoDBs:    []DynamoDBObj{},
		EC2IPs:       []EC2EIPObj{},
		EC2Instances: []EC2InstObj{},
		S3Buckets:    []S3BucketObj{},
	}
	return &s
}

// StoreDynamoDBObj stores DynamoDB objects
func (s *StorageMem) StoreDynamoDBObj(objs []DynamoDBObj) error {
	s.DynamoDBs = []DynamoDBObj{}
	for _, o := range objs {
		s.DynamoDBs = append(s.DynamoDBs, o)
	}
	return nil
}

// SelectDynamoDBObj retrieves DynamoDB objects
func (s *StorageMem) SelectDynamoDBObj() ([]*DynamoDBObj, error) {
	DynamoDBs := []*DynamoDBObj{}
	if len(s.DynamoDBs) == 0 {
		return DynamoDBs, nil
	}
	for _, o := range s.DynamoDBs {
		DynamoDBs = append(DynamoDBs, &o)
	}
	return DynamoDBs, nil
}

// DeleteDynamoDBObjs deletes all DynamoDB objects
func (s *StorageMem) DeleteDynamoDBObjs() error {
	s.DynamoDBs = []DynamoDBObj{}
	return nil
}

// StoreEC2EIPObj stores an array of Elastic IPs in the database.
func (s *StorageMem) StoreEC2EIPObj(objs []EC2EIPObj) error {
	s.EC2IPs = []EC2EIPObj{}
	for _, o := range objs {
		s.EC2IPs = append(s.EC2IPs, o)
	}
	return nil
}

// SelectEC2EIPObj retrieves an array of Elastic IPs from the database.
func (s *StorageMem) SelectEC2EIPObj() ([]*EC2EIPObj, error) {
	EIPs := []*EC2EIPObj{}
	for _, o := range s.EC2IPs {
		EIPs = append(EIPs, &o)
	}
	return EIPs, nil
}

// DeleteEC2EIPObjs removes an array of Elastic IPs from the database.
func (s *StorageMem) DeleteEC2EIPObjs() error {
	s.EC2IPs = []EC2EIPObj{}
	return nil
}

// StoreEC2InstObj stores an array of EC2 Instances in the database.
func (s *StorageMem) StoreEC2InstObj(objs []EC2InstObj) error {
	s.EC2Instances = []EC2InstObj{}
	for _, o := range objs {
		s.EC2Instances = append(s.EC2Instances, o)
	}
	return nil
}

// SelectEC2InstObj retrieves an array of EC2 Instances from the database.
func (s *StorageMem) SelectEC2InstObj() ([]*EC2InstObj, error) {
	EC2Instances := []*EC2InstObj{}
	for _, o := range s.EC2Instances {
		EC2Instances = append(EC2Instances, &o)
	}
	return EC2Instances, nil
}

// DeleteEC2InstObjs removes an array of EC2 Instances from the database.
func (s *StorageMem) DeleteEC2InstObjs() error {
	s.EC2Instances = []EC2InstObj{}
	return nil
}

// StoreS3BucketObj stores an array of S3 Buckets in the database.
func (s *StorageMem) StoreS3BucketObj(objs []S3BucketObj) error {
	s.S3Buckets = []S3BucketObj{}
	for _, o := range objs {
		s.S3Buckets = append(s.S3Buckets, o)
	}
	return nil
}

// SelectS3BucketObj retrieves an array of S3 Buckets from the database.
func (s *StorageMem) SelectS3BucketObj() ([]*S3BucketObj, error) {
	S3Buckets := []*S3BucketObj{}
	for _, o := range s.S3Buckets {
		S3Buckets = append(S3Buckets, &o)
	}
	return S3Buckets, nil
}

// DeleteS3BucketObjs removes an array of S3 Buckets from the database.
func (s *StorageMem) DeleteS3BucketObjs() error {
	s.S3Buckets = []S3BucketObj{}
	return nil
}
