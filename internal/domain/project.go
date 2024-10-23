package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Projects struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Tools       []string           `bson:"tools"`
	icon        string             `bson:"icon"`
	images      []string           `bson:"images"`
	UsuarioID   primitive.ObjectID `bson:"user_id"`
}
