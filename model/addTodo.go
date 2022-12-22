package model

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"todo-api/views"
)

func AddTodo(name string, todo string) error {

	collection := connectDB.Collection("mycol")

	addData, err := collection.InsertOne(context.TODO(), bson.D{
		{ "name", name },
		{ "todo", todo},
	})

	fmt.Println(err)

	if err != nil {
		log.Println("There was an errr in trying to migrate the data into the database")
	}
	
	fmt.Println(addData)
	return nil
}

func GetTodo() ([]views.AddTodo, error ) {
	collection := connectDB.Collection("mycol")

	rows, err := collection.Find(context.TODO(), bson.M{})
	log.Println("rows", rows)

	var todoList []views.AddTodo

	if err = rows.All(context.TODO(), &todoList); err != nil {
		log.Fatal(err)
	}

	fmt.Println(todoList)

	if err != nil {
		log.Println("Errot to get data from the database")
	}

	return todoList, nil
}