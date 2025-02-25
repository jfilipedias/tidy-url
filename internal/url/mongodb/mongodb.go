package mongodb

import (
	"context"
	"net/url"

	"github.com/jfilipedias/tidy-url/internal/constant"
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

	var result url.URL
	err := coll.FindOne(ctx, bson.D{{Key: "hash", Value: hash}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, constant.ErrEntityNotFound
		} else {
			return nil, err
		}
	}
	return &result, err
}
