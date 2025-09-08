package main

import "go.mongodb.org/mongo-driver/mongo"

type testDao struct {
	collection *mongo.Collection
}

func NewTestDao(db *mongo.Database) *testDao {
	return &testDao{
		collection: db.Collection("test"),
	}
}

func main() {

}
