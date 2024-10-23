package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// type User struct {
// 	ID       string `dynamo:"id" json:"id"`
// 	UserName string `dynamo:"username" json:"username"`
// 	Name     string `dynamo:"name" json:"name"`
// 	Password string `dynamo:"password" json:"-"`
// }

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Name     string             `bson:"name"`
	LastName string             `bson:"lastname"`
	Password string             `bson:"password"`
	Projects []Projects         `bson:"projects"`
}
