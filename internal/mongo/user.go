package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID             string `bson:"_id"`
	Name           string `bson:"name"`
	UserName       string `bson:"user_name"`
	Email          string `bson:"email"`
	EmailConfirmed bool   `bson:"email_confirmed"`
	PasswordHash   string `bson:"password_hash"`
}

func (n *User) toBSOND() bson.D {
	return bson.D{
		{Key: "name", Value: n.Name},
		{Key: "user_name", Value: n.UserName},
		{Key: "email", Value: n.Email},
		{Key: "email_confirmed", Value: n.EmailConfirmed},
		{Key: "password_hash", Value: n.PasswordHash},
	}
}

func (d *DBImpl) GetUser(ctx context.Context, id string) (*User, bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, true, fmt.Errorf("bad object id: %s", id)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	res := d.usersCollection.FindOne(ctx, filter)
	user := &User{}

	err = res.Decode(user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, true, err
		}

		return nil, false, err
	}

	return user, true, nil
}

func (d *DBImpl) IsValidUserCredentials(ctx context.Context, email string, passwordHash string) (bool, error) {
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password_hash", Value: passwordHash},
	}

	res := d.usersCollection.FindOne(ctx, filter)

	err := res.Decode(&User{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (d *DBImpl) GetUserByEmail(ctx context.Context, email string) (*User, bool, error) {
	filter := bson.D{{Key: "email", Value: email}}

	res := d.usersCollection.FindOne(ctx, filter)
	user := &User{}

	err := res.Decode(user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, true, err
		}

		return nil, false, err
	}

	return user, true, nil
}

func (d *DBImpl) AddUser(ctx context.Context, user *User) error {
	_, err := d.usersCollection.InsertOne(ctx, user.toBSOND())

	return err
}

func (d *DBImpl) UpdateUser(ctx context.Context, user *User) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return true, fmt.Errorf("bad object id: %s", user.ID)
	}

	update := bson.D{{Key: "$set", Value: user.toBSOND()}}

	_, err = d.usersCollection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
