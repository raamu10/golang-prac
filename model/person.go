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
const COL_ADDRESS = "address"

func AddTodoList(data views.PersonInput) error {
	//fmt.Println("In_addd:", data.TodoList)

	//Connect to db collection
	collectionTodo := connectDB.Collection(COL_TODOLIST)
	collectionPerson := connectDB.Collection(COL_PERSON)
	collectionAddress := connectDB.Collection(COL_ADDRESS)

	//Todo
	var todos []interface{}
	var list []views.Todo
	list = data.TodoList //Conver to the desired type of obj
	for _, elm := range list {
        todos = append(todos, bson.D{
			{"type", elm.Type},
			{"todo", elm.Todo},
		})
	}
	//fmt.Println("todos-->", todos)
	addTodo, err := collectionTodo.InsertMany(context.TODO(), todos)
	if err != nil {
		log.Println("error in adding todo list")
	}
	//fmt.Println("addTodo -->", addTodo)
	var insertedIDs []*primitive.ObjectID  //Conver to the desired type of obj
    for _, insertedID := range addTodo.InsertedIDs {
        if oid, id := insertedID.(primitive.ObjectID); id {
            insertedIDs = append(insertedIDs, &oid)
        }
	}
	//fmt.Println("insertedIDs -->", insertedIDs)

	//Address
	var addressData views.Address
	addressData = data.Address

	addedAddress, err := collectionAddress.InsertOne(context.TODO(), bson.D{
		{ "street", addressData.Street },
		{ "city", addressData.City },
		{ "state", addressData.State },
		{ "country", addressData.Country },
		{ "zipcode", addressData.Zipcode },
	})

	if err != nil {
		log.Println("error in adding address data")
	}
	fmt.Println("addedAddress ->", addedAddress)

	//Get and conversion type for objectId
	var insertedAddressId *primitive.ObjectID
	if oid, ok := addedAddress.InsertedID.(primitive.ObjectID); ok {
		insertedAddressId = &oid
	}

	addPerson, err := collectionPerson.InsertOne(context.TODO(), bson.D{
		{ "name", data.Name },
		{ "email", data.Email},
		{ "todoList", insertedIDs},
		{ "address", insertedAddressId},
	})
	
	if err != nil {
		log.Println("error in adding person todo")
	}

	//Get and conversion type for objectId
	var personInsertId *primitive.ObjectID
	if oid, ok := addPerson.InsertedID.(primitive.ObjectID); ok {
		personInsertId = &oid
	}

	//Update Address with personId
	filter := bson.D{
		{"_id", insertedAddressId},
	}
	update := bson.D{
		{ "$set" , bson.D{
			{"personId", personInsertId},
			},
		},
	}

	res, err := collectionAddress.UpdateOne(context.TODO(),
		filter,
		update,
	)

	if err != nil {
		log.Println("error in updating person id in address")
	}

	fmt.Println("result->", res)
	fmt.Println("addPersons->",addPerson)

	return nil
}

/* 
 * @name GetPersonData
 * @description
 * Get the Person data by personID and return person data, todolist, address
 * retrun type views.Person, []views.Todo, views.Address, error
 */
func GetPersonData(id string) (views.Person, []views.Todo, views.Address, error) {

	collectionPerson := connectDB.Collection(COL_PERSON)
	collectionTodo := connectDB.Collection(COL_TODOLIST)
	collectionAddress := connectDB.Collection(COL_ADDRESS)

	// convert id string to ObjectId
	personId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
	}

	//person data
	var personData views.Person
	personRows := collectionPerson.FindOne(context.TODO(), bson.M{"_id": personId})
	personRows.Decode(&personData)

	//fmt.Println("rows ->", personData)
	//fmt.Println("rows_todoList ->", personData.TodoList)

	if err != nil {
		log.Println("error in getting person data")
	}

	//TODO data
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

	//fmt.Println("tododIds->", tododIds)

	todoRows, err:= collectionTodo.Find(context.TODO(), bson.M{"_id": bson.M{ "$in" : tododIds}})
	if err != nil {
		log.Println("error in todo list get")
	}

	var todoDataList []views.Todo //Assigning the value
	if err = todoRows.All(context.TODO(), &todoDataList); err != nil {
		log.Fatal(err)
	}
	fmt.Println("todoDataList ->", todoDataList)

	//Address data
	addressRow := collectionAddress.FindOne(context.TODO(), bson.M{"personId": personId})
	if  err != nil {
		log.Fatal(err)
		fmt.Println("Error in getting address with Person id")
	}

	//Decode address row
	var addressData views.Address
	addressRow.Decode(&addressData)

	fmt.Println("Address resp ->", addressData)

	return personData, todoDataList, addressData, nil

}

/**
 * Get the address data by person id
*/
func GetAddressByPersonId(id string) (views.Address, error) {
	//Address Collection
	collectionAddress := connectDB.Collection(COL_ADDRESS)

	// convert id string to ObjectId
	personId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
	}

	//Address data
	addressRow := collectionAddress.FindOne(context.TODO(), bson.M{"personId": personId})
	if  err != nil {
		log.Fatal(err)
		fmt.Println("Error in getting address with Person id")
	}

	//Decode address row
	var addressData views.Address
	addressRow.Decode(&addressData)

	return addressData, nil

}