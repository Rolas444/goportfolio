package user

import (
	"goportfolio/internal/ports/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	dbclient *mongo.Client
	dbname   string
	collname string
}

func NewUserRepository(client *mongo.Client, dbname string) service.UserRepository {
	return &Repository{dbclient: client, dbname: dbname, collname: "users"}
}
