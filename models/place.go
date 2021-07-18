package models

import (
	"clickbus/db"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Place struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Slug      string             `bson:"slug,omitempty"`
	City      string             `bson:"city,omitempty"`
	State     string             `bson:"state,omitempty"`
	CreatedAt string             `bson:"createdAt,omitempty"`
	UpdatedAt string             `bson:"updatedAt,omitempty"`
}

func getCollection() *mongo.Collection {
	return db.GetDb().Database("clickBus").Collection("places")
}

func GetAllPlaces(query map[string]string) ([]Place, error) {
	collection := getCollection()

	var places []Place
	queryMap := map[string]bson.M{}
	for k, v := range query {
		queryMap[k] = bson.M{"$regex": v, "$options": "i"}
	}
	cursor, err := collection.Find(context.TODO(), queryMap)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &places); err != nil {
		return nil, err
	}

	return places, nil
}

func GetPlace(id string) (*Place, error) {
	collection := getCollection()
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	var result Place

	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetPlaceNew(id interface{}) (*Place, error) {
	collection := getCollection()

	filter := bson.M{"_id": id}

	var result Place

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreatePlace(name, slug, city, state string) (*mongo.InsertOneResult, error) {
	collection := getCollection()

	book := Place{Name: name, Slug: slug, City: city, State: state, CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()}
	insertResult, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func UpdatePlace(update map[string]string, id string) (*mongo.UpdateResult, error) {
	collection := getCollection()

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})

	if err != nil {
		return nil, err
	}

	return updateResult, err
}

func DeletePlace(id string) error {
	collection := getCollection()
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	if err != nil {
		return err
	}

	return nil
}
