package controllers

import (
	"encoding/json"
	"fmt"
	"github/tarun1369/golangXmongo/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		// is not valid object id
		w.WriteHeader(http.StatusNotFound) // 404 // header has status/error code
	}
	oid := bson.ObjectIdHex(id) // object id
	u := models.User{}        // empty user struct

	if err := uc.session.DB("golangXmongo").C("users").FindId(oid).One(&u); err != nil {
		// c for collection
		w.WriteHeader(404) // 404 as after finding id we are not able to find user
		return
	}

	uj, err := json.Marshal(u) // for json encoding

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") // postman ko bata rahe yaha
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	// r - postman se aayi hai
	// _ kyuki param ka kaam nhi

	// create user from json
	json.NewDecoder(r.Body).Decode(&u) // decode json and put it in u
	u.Id = bson.NewObjectId()

	uc.session.DB("golangXmongo").C("users").Insert(u) // insert user in mongodb

	// also sending data back to user in json format
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) DeleteUser(w http.ResponseWriter,r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("golangXmongo").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	} 

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted user",oid,"\n")

}
