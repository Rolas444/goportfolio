package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ConnectClient(region string, idKey string, secretKey string) (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(idKey, secretKey, ""),
	})
	if err != nil {
		// log.Fatalf("Error al crear la sesión: %v", err)
		return nil, err
	}

	dbClient := dynamodb.New(sess)
	return dbClient, nil
}
