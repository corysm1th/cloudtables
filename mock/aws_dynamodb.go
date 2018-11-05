package mock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	cloudtables "github.com/corysm1th/cloudtables/pkg"
)

// DynamoDBClient mocks the AWS API DynamoDB endpoint.
type DynamoDBClient struct{ dynamodbiface.DynamoDBAPI }

// ListTables returns an array of mock DynamoDB tables.
func (m *DynamoDBClient) ListTables(i *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	table1, table2 := "TestTable1", "TestTable2"
	t1, t2 := &table1, &table2
	data := dynamodb.ListTablesOutput{TableNames: []*string{t1, t2}}
	return &data, nil
}

// CreateDynamoDBs returns an array of pointers to DynamoDB objects.
func CreateDynamoDBs() []cloudtables.DynamoDBObj {
	names := []string{"users", "secrets", "contacts", "products"}
	dbs := []cloudtables.DynamoDBObj{}
	for _, n := range names {
		db := cloudtables.DynamoDBObj{
			Account: "test",
			Region:  "us-west-2",
			Name:    n,
		}
		dbs = append(dbs, db)
	}
	return dbs
}
