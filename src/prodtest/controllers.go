package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection = db().Database("prodTest").Collection("users")

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person User

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		fmt.Print(err)
	}

	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(insertResult.InsertedID)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body User

	e := json.NewDecoder(r.Body).Decode(&body)

	if e != nil {
		fmt.Print(e)
	}

	var result primitive.M

	err := userCollection.FindOne(context.TODO(), bson.D{{"name",
		body.Name}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json: "name"`
		Age  int    `json: "age"`
		City string `json: "city"`
	}

	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}

	filter := bson.D{{"name", body.Name}}
	after := options.After

	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{{"$set", bson.D{{"city", body.City},
		{"age", body.Age}}}}

	updateResult := userCollection.FindOneAndUpdate(context.TODO(),
		filter, update, &returnOpt)
	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		fmt.Println(err.Error())
	}

	opts := options.Delete().SetCollation(&options.Collation{})

	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(res.DeletedCount)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := userCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(context.TODO()) {

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(results)
}
