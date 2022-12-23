package views 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID	   primitive.ObjectID 	`bson:"_id" json:"id,omitempty"`
	Type   string    			`json:"type"`
	Todo   string 				`json:"todo"`
}

type Address struct {
	ID	     primitive.ObjectID	`bson:"_id" json:"id,omitempty"`
	PersonId primitive.ObjectID	`json:"personId,omitempty"`
	Street string				`json:"stree"`
	City    string 				`json:"city"`
	State	string				`json:"state"`
	Country string				`json:"country"`
	Zipcode int 				`json:"zipcode"`
}

type Person struct {
	ID		primitive.ObjectID  `bson:"_id" json:"id,omitempty"`
	Name	string			    `json:"name"`
	Email	string				`json:"email"`
	TodoList []string		    `json:"todoList"`
	Address Address				`json:address`
}

type PersonInput struct {
	ID		primitive.ObjectID  `bson:"_id" json:"id,omitempty"`
	Name	string			    `json:"name"`
	Email	string				`json:"email"`
	TodoList []Todo		        `json:"todoList"`
	Address Address				`json:address`
}