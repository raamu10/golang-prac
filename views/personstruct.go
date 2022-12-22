package views 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID	   primitive.ObjectID 	`bson:"_id" json:"id,omitempty"`
	Type   string    			`json:"type"`
	Todo   string 				`json:"todo"`
}

type Person struct {
	ID		primitive.ObjectID  `bson:"_id" json:"id,omitempty"`
	Name	string			    `json:"name"`
	Email	string				`json:"email"`
	TodoList []string		     `json:"todoList"`
}

type PersonInput struct {
	ID		primitive.ObjectID  `bson:"_id" json:"id,omitempty"`
	Name	string			    `json:"name"`
	Email	string				`json:"email"`
	TodoList []Todo		        `json:"todoList"`
}