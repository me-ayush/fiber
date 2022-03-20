package main

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const dbname = "personDb"
const collectionName = "Person"

func indexRoute(res *fiber.Ctx) error {
	// return res.SendString("Hello World")
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some Problem! Please try again")
	}
	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res.Status(400).SendString("There is some Problem! Please try again")
	}
	defer curr.Close(context.Background())
	var result []bson.M
	curr.All(context.Background(), &result)
	json, _ := json.Marshal(result)
	return res.Status(200).Send(json)

}
func addPerson(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some Problem! Please try again")
	}
	var newPerson Person
	json.Unmarshal([]byte(res.Body()), &newPerson)
	curr, err := collection.InsertOne(context.Background(), newPerson)
	if err != nil {
		return res.Status(400).SendString("There is some Problem! Please try again")
	}
	response, _ := json.Marshal(curr)
	return res.Status(200).Send(response)
}

func updatePerson(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some Problem! Please try again")
	}
	id := res.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var updatePerson Person
	json.Unmarshal([]byte(res.Body()), &updatePerson)
	var filter bson.M = bson.M{
		"_id": objID,
	}
	update := bson.M{
		"$set": updatePerson,
	}
	curr, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return res.Status(400).SendString("There is some Problem! Please try again")
	}
	response, _ := json.Marshal(curr)
	return res.Status(200).Send(response)

}
func main() {
	app := fiber.New()
	app.Get("/", indexRoute)
	app.Post("/create", addPerson)
	app.Put("/update/:id", updatePerson)

	app.Listen(":3000")
}
