package main

import (
	"context"
	"covid-tracker/constants"
	"covid-tracker/contract"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)


var client *mongo.Client

func main() {
	resp, err := http.Get(constants.CovidCasesUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(constants.MongoDbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	var ans map[string]map[string]string
	json.Unmarshal([]byte(body), &ans)
	for i := 0; i <= 35; i++ {
		state := contract.State{Name: ans["Name of State / UT"][strconv.Itoa(i)], AbbreviationCode: ans["abbreviation_code"][strconv.Itoa(i)], StateCode: ans["state_code"][strconv.Itoa(i)]}
		fmt.Println(state)
		col := client.Database("covid-tracker").Collection("state")
		result, insertErr := col.InsertOne(ctx, state)
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
			os.Exit(1) // safely exit script on error
		} else {
			fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
			fmt.Println("InsertOne() API result:", result)
			newID := result.InsertedID
			stringObjectID := newID.(primitive.ObjectID).Hex()
			fmt.Println(stringObjectID)
			col := client.Database("covid-tracker").Collection("cases")

			activeCases, _ := strconv.Atoi(ans["Active"][strconv.Itoa(i)])
			casesActive := contract.Cases{StateID: stringObjectID, Status: "active", Count: activeCases}
			inactiveCases, _ := strconv.Atoi(ans["Cured/Discharged/Migrated"][strconv.Itoa(i)])
			casesInactive := contract.Cases{StateID: stringObjectID, Status: "inactive", Count: inactiveCases}
			deathCount, _ := strconv.Atoi(ans["Death"][strconv.Itoa(i)])
			deathCountCases := contract.Cases{StateID: stringObjectID, Status: "death", Count: deathCount}
			_, insertErr := col.InsertOne(ctx, casesActive)

			if insertErr != nil {
				fmt.Println("InsertOne ERROR in active cases:", insertErr)
			}
			_, insertErr = col.InsertOne(ctx, casesInactive)
			if insertErr != nil {
				fmt.Println("InsertOne ERROR in inactive cases:", insertErr)
			}
			_, insertErr = col.InsertOne(ctx, deathCountCases)
			if insertErr != nil {
				fmt.Println("InsertOne ERROR in death count:", insertErr)
			}
		}
	}
}

				//cases_active := contract.Cases{StateID: reflect.ValueOf(newID).String(), Status: ans["Active"][strconv.Itoa(i)]}
				//cases_inactive := contract.Cases{StateID: reflect.ValueOf(newID).String(), Status: ans["Cured/Discharged/Migrated"][strconv.Itoa(i)]}
				//cases_death := contract.Cases{StateID: reflect.ValueOf(newID).String(), Status: ans["Death"][strconv.Itoa(i)]}
				//col := client.Database("covid-tracker").Collection("cases")
				//_, insertErr := col.InsertOne(ctx, cases_active)
				//_, insertErr = col.InsertOne(ctx, cases_inactive)
				//_, insertErr = col.InsertOne(ctx, cases_death)
				//if insertErr != nil {
				//	fmt.Println("InsertOne ERROR:", insertErr)
				//	os.Exit(1) // safely exit script on error
				//
				//}
				//col := client.Database("covid-tracker").Collection("state")
				//filterCursor, err := col.Find(ctx, bson.M{"name": "Delhi"})
				//if err != nil {
				//	log.Fatal(err)
				//}
				//var episodesFiltered []bson.M
				//if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
				//	log.Fatal(err)
				//}
				//asd:=((episodesFiltered[0]["_id"]).(primitive.ObjectID).Hex())
				//col2 := client.Database("covid-tracker").Collection("cases")
				//filterCursor2, err := col2.Find(ctx, bson.M{"stateid": asd})
				//if err != nil {
				//	log.Fatal(err)
				//}
				//var episodesFiltered22 []bson.M
				//if err = filterCursor2.All(ctx, &episodesFiltered22); err != nil {
				//	log.Fatal(err)
				//}
				//fmt.Println(episodesFiltered22)
				//a := result["state_code"]["0"]
				////v :=a[]
				//fmt.Println(reflect.TypeOf(a))
				//
				////valq , _ := a["0"]
				//fmt.Println(a)

				//var data output2
				//json.Unmarshal(body, &data)
				//fmt.Println(data.as)
				//ctx , _ := context.WithTimeout(context.Background(), 10*time.Second)
				//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
				//client,err = mongo.Connect(ctx, clientOptions)
				//if err != nil {
				//	log.Fatal(err)
				//}
				//// Check the connection
				//err = client.Ping(ctx, nil)
				//if err != nil {
				//	log.Fatal(err)
				//}
				//fmt.Println("Connected to MongoDB!")
				//databases, err := client.ListDatabaseNames(ctx, bson.M{})
				//if err != nil {
				//	log.Fatal(err)
				//}
				//fmt.Println(databases)

				//a := contract.State{Name: "delhi", StateCode: "0", AbbreviationCode: "1"}
				//col := client.Database("covid-tracker").Collection("state")
				//result, insertErr := col.InsertOne(ctx, a)
				//if insertErr != nil {
				//	fmt.Println("InsertOne ERROR:", insertErr)
				//	os.Exit(1) // safely exit script on error
				//} else {
				//	fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
				//	fmt.Println("InsertOne() API result:", result)
				//
				//	// get the inserted ID string
				//	newID := result.InsertedID
				//	fmt.Println("InsertOne() newID:", newID)
				//	fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))
				//}




