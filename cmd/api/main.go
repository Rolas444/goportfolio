package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

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

	svc := dynamodb.New(sess)

	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalf("Error al listar tablas: %v", err)
	}

	names := []string{}

	fmt.Println("Tablas en DynamoDB:")
	for _, tableName := range result.TableNames {

		names = append(names, *tableName)

		fmt.Println(*tableName)
	}

	ginEngine := gin.Default()

	ginEngine.POST("/projects", func(c *gin.Context) {
		c.JSON(200, gin.H{"tables ": names[0]})
	})

	log.Fatalln(ginEngine.Run(":8001"))

}
