package cloudtables

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"
)

// GetDynamoDB fetches DynamoDB resources
func GetDynamoDB(svc dynamodbiface.DynamoDBAPI, account, region string) ([]*DynamoDBObj, int, error) {
	var DynamoDBObjs []*DynamoDBObj
	var count int
	input := dynamodb.ListTablesInput{}
	obj := &DynamoDBObj{}
	obj.Account = account
	obj.Region = region
	tables, err := svc.ListTables(&input)
	if err != nil {
		return DynamoDBObjs, count, errors.Wrap(err, "ListTables Request Failed.")
	}
	count = len(tables.TableNames)
	for _, n := range tables.TableNames {
		obj.Name = *n
		DynamoDBObjs = append(DynamoDBObjs, obj)
	}
	return DynamoDBObjs, count, nil
}
