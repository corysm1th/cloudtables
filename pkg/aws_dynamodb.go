package cloudtables

import (
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"
)

// SyncDynamoDB fetches DynamoDB resources, and stores them in the database.
func SyncDynamoDB(svc dynamodbiface.DynamoDBAPI, account, region string) error {
	t := dynamodb.ListTablesInput{}
	o := DynamoDBTableObj{}
	o.Account = account
	o.Region = region
	tables, err := svc.ListTables(&t)
	if err != nil {
		return errors.Wrap(err, "ListTables Request Failed.")
	}
	log.Println("DynamoDB")
	for _, n := range tables.TableNames {
		o.Name = *n
		log.Printf("Account: %s  Region: %s  Table: %s", o.Account, o.Region, o.Name)
	}
	// TODO: Store table in dynamodb_table_obj
	// TODO: Store count in metric
	return nil
}
