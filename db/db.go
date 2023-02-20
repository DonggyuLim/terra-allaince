package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DonggyuLim/Alliance-Rank/account"
	"github.com/DonggyuLim/Alliance-Rank/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Client
var dbName string
var collectionName string

func Connect() {
	url := utils.LoadENV("DBURL", "db.env")
	dbName = utils.LoadENV("DBNAME", "db.env")
	collectionName = utils.LoadENV("Collection", "db.env")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(url).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, clientOptions)
	// ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	db = client
	fmt.Println("============DB connect==================")
}

func Close() {
	err := db.Disconnect(context.TODO())
	utils.HandleErr("DB Disconnect", err)
	fmt.Println("=========Connection to MongoDB closed=============")
}
func GetCollection() *mongo.Collection {
	return db.Database(dbName).Collection(collectionName)
}

func Insert(account account.Account) error {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, account)
	if err != nil {
		return err
	}

	fmt.Println("Insert Complete", insertResult.InsertedID)
	return nil
}

func InsertMany(data []account.Account) {
	var a []interface{}
	for _, el := range data {
		a = append(a, el)
	}
	exp := 20 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	_, err := collection.InsertMany(ctx, a)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Insert End")
}

func FindOne(filter bson.D, a *account.Account) error {

	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	err := collection.FindOne(ctx, filter).Decode(a)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return mongo.ErrNoDocuments
		}
		log.Fatal(err)
	}
	return nil
}

func Find(key, value, desc string, limit int64) ([]account.Account, error) {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	findOptions := options.Find()

	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: desc, Value: -1}})
	var filter primitive.D
	if key == "" && value == "" {
		filter = bson.D{}
	} else {
		filter = bson.D{{Key: key, Value: value}}
	}

	cur, _ := collection.Find(ctx, filter, findOptions)
	var curs []account.Account
	err := cur.All(context.TODO(), &curs)
	return curs, err
}

func FindAndReplace(filter, update bson.D) {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)

	result := collection.FindOneAndReplace(ctx, filter, update)
	fmt.Println("DB update")
	fmt.Println(result.Err().Error())
}

func ReplaceOne(filter bson.D, account account.Account) {

	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)

	_, err := collection.ReplaceOne(ctx, filter, account)

	utils.PanicError(err)
}

func UpdateOne(filter, update bson.D) {
	// fmt.Println("Update")
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	collection := db.Database(dbName).Collection(collectionName)
	_, err := collection.UpdateOne(ctx, filter, update)
	utils.PanicError(err)
	// fmt.Println("Update End!")
}
