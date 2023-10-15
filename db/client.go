package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client structure holds the MongoDB client and context
type Client struct {
	Client *mongo.Client
	Ctx    context.Context
	dbName string
}

// NewClient creates a new MongoDB client
func NewClient(uri, dbName string) (*Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{client, ctx, dbName}, nil
}

// InsertDocument inserts a new document into the specified MongoDB collection
func (c *Client) InsertDocument(collectionName string, document interface{}) error {
	collection := c.Client.Database(c.dbName).Collection(collectionName)
	_, err := collection.InsertOne(c.Ctx, document)
	return err
}

// UpdateDocument updates a document in the specified MongoDB collection
func (c *Client) UpdateDocument(collectionName string, filter, update interface{}) error {
	collection := c.Client.Database(c.dbName).Collection(collectionName)
	_, err := collection.UpdateOne(c.Ctx, filter, update)
	return err
}

// ReplaceDocument replaces a document in the specified MongoDB collection
func (c *Client) ReplaceDocument(collectionName string, filter, replacement interface{}) error {
	collection := c.Client.Database(c.dbName).Collection(collectionName)
	_, err := collection.ReplaceOne(c.Ctx, filter, replacement)
	return err
}

// RemoveDocument removes a document from the specified MongoDB collection
func (c *Client) RemoveDocument(collectionName string, filter interface{}) error {
	collection := c.Client.Database(c.dbName).Collection(collectionName)
	_, err := collection.DeleteOne(c.Ctx, filter)
	return err
}

// CreateCollection creates a new collection in the MongoDB database
func (c *Client) CreateCollection(collectionName string) error {
	return c.Client.Database(c.dbName).CreateCollection(c.Ctx, collectionName)
}

// AddItemToArray adds an item to an array in a document
func (c *Client) AddItemToArray(collectionName string, filter interface{}, arrayName string, item interface{}) error {
	collection := c.Client.Database(c.dbName).Collection(collectionName)
	update := bson.M{"$push": bson.M{arrayName: item}}
	_, err := collection.UpdateOne(c.Ctx, filter, update)
	return err
}

// RemoveItemFromArray removes an item from an array in a document
func (c *Client) RemoveItemFromArray(collectionName string, filter interface{}, arrayName string, item interface{}) error {
	collection := c.Client.Database(c.dbName).Collection(collectionName)
	update := bson.M{"$pull": bson.M{arrayName: item}}
	_, err := collection.UpdateOne(c.Ctx, filter, update)
	return err
}
