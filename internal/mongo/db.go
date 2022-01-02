package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ DB = (*DBImpl)(nil)

type DB interface {
	// AddNote - add usere to database.
	AddUser(context.Context, *User) error
	// GetNote - get user by id.
	GetUser(context.Context, string) (*User, bool, error)
	// UpdateUser - update user by id.
	UpdateUser(context.Context, *User) (bool, error)
}

type DBImpl struct {
	mongoDB         *mongo.Database
	usersCollection *mongo.Collection
}

func NewDBImpl(db *mongo.Database, usersCollection *mongo.Collection) *DBImpl {
	return &DBImpl{
		mongoDB:         db,
		usersCollection: usersCollection,
	}
}

func Connect(ctx context.Context, addr string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(addr))
}
