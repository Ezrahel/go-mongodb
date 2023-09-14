package main

import (
	"net/http"

	"github.com/ezrahel/mongodb-go/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	//r.PUT()
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8000", r)

}

func getSession() *mgo.Session {
	//s, err := mgo.Dail("mongodb://localhost:27107")
	s, err := mgo.Dial("mongodb+srv://ezrahel:Taepryung@cluster0.ncdi5ro.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	return s
}