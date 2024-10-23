package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	// if err != nil {
	// 	log.Fatalf("Error al listar tablas: %v", err)
	// }

	// names := []string{}

	// fmt.Println("Tablas en DynamoDB:")
	// for _, tableName := range result.TableNames {

	// 	names = append(names, *tableName)

	// 	fmt.Println(*tableName)
	// }

	ginEngine := gin.Default()

	// ginEngine.POST("/projects", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"tables ": names[0]})
	// })

	log.Fatalln(ginEngine.Run(":8001"))

}
