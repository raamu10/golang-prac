package model

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-api/views"
	//"reflect"
)

const COL_TODOLIST = "todolist"
const COL_PERSON = "person"

func AddTodoList(data views.PersonInput) error {
	//fmt.Println("In_addd:", data.TodoList)

	var todos []interface{}
	var list []views.Todo
	list = data.TodoList //Conver to the desired type of obj
	for _, elm := range list {
        todos = append(todos, bson.D{
			{"type", elm.Type},
			{"todo", elm.Todo},
		})
	}
	
	fmt.Println("todos-->", todos)

	collectionTodo := connectDB.Collection(COL_TODOLIST)
	collectionPerson := connectDB.Collection(COL_PERSON)
	
	addTodo, err := collectionTodo.InsertMany(context.TODO(), todos)
	if err != nil {
		log.Println("error in adding todo list")
	}
	
	fmt.Println("addTodo -->", addTodo)

	var insertedIDs []*primitive.ObjectID  //Conver to the desired type of obj
    for _, insertedID := range addTodo.InsertedIDs {
        if oid, id := insertedID.(primitive.ObjectID); id {
            insertedIDs = append(insertedIDs, &oid)
        }
	}
	
	fmt.Println("insertedIDs -->", insertedIDs)

	addPerson, err := collectionPerson.InsertOne(context.TODO(), bson.D{
		{ "name", data.Name },
		{ "email", data.Email},
		{ "todoList", insertedIDs},
	})

	if err != nil {
		log.Println("error in adding person todo")
	}
	
	fmt.Println("addPersons->",addPerson)

	return nil
}

func GetPersonData(id string) (views.Person, []views.Todo, error) {

	collectionPerson := connectDB.Collection(COL_PERSON)
	collectionTodo := connectDB.Collection(COL_TODOLIST)

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
	}

	var personData views.Person
	/* rows, err := collectionPerson.Find(context.TODO(), bson.D{
		{"_id", id},
	}) */

	personRows := collectionPerson.FindOne(context.TODO(), bson.M{"_id": objectId})

	personRows.Decode(&personData)
	fmt.Println("rows ->", personData)
	fmt.Println("rows_todoList ->", personData.TodoList)

	if err != nil {
		log.Println("error in getting person data")
	}

	var tododIds []*primitive.ObjectID  //Conver to the desired type of obj
    for _, insertedID := range personData.TodoList {

		//fmt.Println(insertedID)
		//fmt.Println(primitive.ObjectIDFromHex(insertedID))
        oid, err := primitive.ObjectIDFromHex(insertedID)
		//fmt.Println("oid:", oid)
		//fmt.Println("id:", err)
		if err != nil {
			log.Println("error in object id")
		}
		tododIds = append(tododIds, &oid)
	
	}

	fmt.Println("tododIds->", tododIds)

	todoRows, err:= collectionTodo.Find(context.TODO(), bson.M{"_id": bson.M{ "$in" : tododIds}})
	if err != nil {
		log.Println("error in todo list get")
	}

	var todoDataList []views.Todo //Assigning the value
	if err = todoRows.All(context.TODO(), &todoDataList); err != nil {
		log.Fatal(err)
	}

	fmt.Println("todoDataList ->", todoDataList)

	return personData, todoDataList, nil

}