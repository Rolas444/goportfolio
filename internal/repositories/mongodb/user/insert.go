package user

import (
	"context"
	"goportfolio/internal/domain"
)

func (r *Repository) Insert(user domain.User) (string, error) {
	collection := r.dbclient.Database(r.dbname).Collection(r.collname)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
