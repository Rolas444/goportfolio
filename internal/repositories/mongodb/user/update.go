package user

import (
	"context"
	"goportfolio/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Update(user domain.User) (string, error) {

	collection := r.dbclient.Database(r.dbname).Collection(r.collname)
	// idColl := user.ID
	objID, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{"$set": user}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", err
	}

	return string(result.ModifiedCount), nil
}
