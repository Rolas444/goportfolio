package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateCollections(dbc *dynamodb.DynamoDB) {

}

func CheckTableEsists(dbc *dynamodb.DynamoDB, tableName string) error {
	_, err := dbc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	return err
}

// func createTableEsxists(dbc *dynamodb.DynamoDB, tablename string) error {
// 	_, err := dbc.CreateTable(&dynamodb.CreateTableInput{})
// }
