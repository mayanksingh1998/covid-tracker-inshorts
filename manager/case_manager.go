package manager

import (
	"context"
	"covid-tracker/cache"
	"covid-tracker/contract"
	mongo2 "covid-tracker/mongo"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reflect"
)
var client *mongo.Client
var ctx = context.Background()

func GetCaseForCity(city string) [] contract.Cases{
	var response []contract.Cases
	rdb:= cache.CacheInitializer()
	val,_ := rdb.Get(ctx, "Delhi").Result()
	_ = json.Unmarshal([]byte(val), response)
	if reflect.ValueOf(response).IsNil(){
		client := mongo2.GetMongoClient()
		col := client.Database("covid-tracker").Collection("state")
		filterCursor, err := col.Find(ctx, bson.M{"name": city})
		if err != nil {
			log.Fatal(err)
		}
		var stateInfo []bson.M
		if err = filterCursor.All(ctx, &stateInfo); err != nil {
			log.Fatal(err)
		}
		stateID := (stateInfo[0]["_id"]).(primitive.ObjectID).Hex()
		col2 := client.Database("covid-tracker").Collection("cases")
		filterCursor2, err := col2.Find(ctx, bson.M{"stateid": stateID})
		if err != nil {
			log.Fatal(err)
		}
		if err = filterCursor2.All(ctx, &response); err != nil {
			log.Fatal(err)
		}
		p, err := json.Marshal(response)
		err = rdb.Set(ctx, city, p, 0).Err()
	}
	return response
}
