# golang-prac

GoLang Practice repo
Desc: CRUD operation with GoLang and MongoDB

# Steps
1. Install GoLang
2. Create main.go file which has main package
# Create module init
3. go mod init 
3. Create folders
   controller/
   model/
   views/
4. controller - folder has routing/mux, api request
5. model - folder process the request and connect with DB to get data
6. views - folder has structs,interfaces for object creations

# Run the script command from home folder
go run main.go

#import mongodb packages
#mongodb connection
could check sample code in model/connect.go
import (
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)