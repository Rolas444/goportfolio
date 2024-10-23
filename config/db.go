package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ClientDynamoDB() *dynamodb.DynamoDB {
	//connect to aws dynamoDB
	region := os.Getenv("AWS_REGION")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("Error al crear la sesión: %v", err)
	}

	return dynamodb.New(sess)
}

func ClientMongoDB() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client, err

}
