package mongodb

import (
	"context"
	"net/url"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDB struct {
	conn *mongo.Client
}

func NewMongoDB(client *mongo.Client) *MongoDB {
	return &MongoDB{client}
}

func (m *MongoDB) Create(ctx context.Context, u *url.URL) error {
	coll := m.conn.Database("tidy-url").Collection("urls")
	_, err := coll.InsertOne(ctx, u)
	return err
}

func (m *MongoDB) Get(ctx context.Context, hash string) (*url.URL, error) {
	coll := m.conn.Database("tidy-url").Collection("urls")
	f := bson.D{{Key: "hash", Value: hash}}

	var result url.URL
	err := coll.FindOne(ctx, f).Decode(&result)
	return &result, err
}
