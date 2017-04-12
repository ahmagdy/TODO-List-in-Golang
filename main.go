package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"strings"

	"os"

	"github.com/maxwellhealth/bongo"
)

//TODO : a Strucutre Representation of TODO List
type TODO struct {
	bongo.DocumentBase `bson:",inline"`
	Task               string
	IsDone             bool
}

func main() {
	config := &bongo.Config{
		ConnectionString: "localhost:27017",
		Database:         "TODOGO",
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected To DB")
	todoSlice := []*TODO{&TODO{Task: "Task 1", IsDone: false}, &TODO{Task: "Task 2", IsDone: false}}
	AddTODO(connection, todoSlice...)
	// DeleteingleTODO(connection, "58ed6b283464f927d49388ae")
	// FindAllToDos(connection)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("What Do you want to do?\n", "1 - Get All Elements\n", "2- Delete Element\n", "3- Add new TODO")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		FindAllToDos(connection)
	case 2:
		fmt.Println("Enter The ID of TODO")
		var id string
		fmt.Scan(&id)
		DeleteingleTODO(connection, id)
	case 3:
		fmt.Println("Enter Task name")
		var inputString string
		fmt.Scanln(&inputString)
		fmt.Println(inputString)
		if inputString == "" {
			os.Exit(0)
		}
		AddTODO(connection, &TODO{Task: inputString, IsDone: false})
	default:
		fmt.Println("Wrong Number")
	}
}

// AddTODO :  to Add Data To Collection
func AddTODO(connection *bongo.Connection, todos ...*TODO) {
	for _, todo := range todos {
		err := connection.Collection("TODOS").Save(todo)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Added Successfully ")
}

// FindAllToDos :  To Get All Rows in The Collection
func FindAllToDos(connection *bongo.Connection) {
	res := connection.Collection("TODOS").Find(bson.M{})
	some := new(TODO)
	for i := 0; res.Next(some); i++ {
		fmt.Println(i+1, some.Id, some.Task)
	}
	fmt.Println("Finish")
}

// DeleteingleTODO :  To Delete Single Element By ID
func DeleteingleTODO(connection *bongo.Connection, id string) {
	err := connection.Collection("TODOS").DeleteOne(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		panic(err)
	}

	fmt.Println("TODO Deleted Successfully !!!")
}
