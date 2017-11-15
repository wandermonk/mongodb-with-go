package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	database := session.DB("test")
	collection := database.C("test")
	fmt.Print(collection.Name)
	session.Close()

}
