package model

import (
	"fmt"
	"context"
	"log"
    "time"
    //"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db_name = "golangdb"
var connectDB *mongo.Database
var CTX context.Context
var CancelFunc context.CancelFunc

//var CTX *contex

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.
 
func connect(uri string)(*mongo.Client, context.Context,
                          context.CancelFunc, error) {
                           
    // ctx will be used to set deadline for process, here
    // deadline will of 30 seconds.
    ctx, cancel := context.WithTimeout(context.Background(),
                                       30 * time.Second)
     
    // mongo.Connect return mongo.Client method
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc){
            
    fmt.Println("Connection Closed!")
    // CancelFunc to cancel to context
    defer cancel()
     
    // client provides a method to close
    // a mongoDB connection.
    defer func(){
     
        // client.Disconnect method also has deadline.
        // returns error if any,
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error{
 
    // mongo.Client has Ping to ping mongoDB, deadline of
    // the Ping method will be determined by cxt
    // Ping method return error if any occurred, then
    // the error can be handled.
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}

func MongoConnector() (*mongo.Client) {
	fmt.Println("in MongoConnect!!")

	//Connect MongoDB
	// Get Client, Context, CancelFunc and
    // err from connect method.
    client, ctx, cancel, err := connect("mongodb://localhost:27017")
    //fmt.Println("ctx type:", reflect.TypeOf(ctx))
    CTX = ctx
    CancelFunc = cancel
    if err != nil {
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    //defer close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
    ping(client, ctx)

	//Databases
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Databases:")
    fmt.Println(databases)

    //database := client.Database(db_name)
    connectDB = client.Database(db_name)
    
    return client

	/* fmt.Println("db_name:",db_name)
    database := client.Database(db_name)
    connectDB = client.Database(db_name)
    theCollection := database.Collection("mycol")
    fmt.Println("theCollection type",reflect.TypeOf(theCollection))

    coldata, err := theCollection.Find(context.TODO(), bson.M{})
    
	if err != nil {
		log.Println(err)
	  
	}
	var data []bson.M
	if err = coldata.All(context.TODO(), &data); err != nil {
		log.Println(err)
    }
    fmt.Println("data:", data) 
	fmt.Println(data) */

}

func MongoClose(client *mongo.Client) {

    // Release resource when the main
    // function is returned.
    defer close(client, CTX, CancelFunc)

}