package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	log.Println("create a session with mongodb")

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("peoples")
	err = c.Insert(&Person{"Gosho", 27}, &Person{"Pesho", 10})

	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Gosho"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("New phone who dis: ", result.Name)
}
