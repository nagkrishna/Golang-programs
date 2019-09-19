package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"../models"
	"github.com/julienschmidt/httprouter"
)

//UserController export
type UserController struct {
	session *mgo.Session
}

//NewUserController export
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//Index export
func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html>
	<head>
		<meta charset='utf-8'>
		<meta http-equiv='X-UA-Compatible' content='IE=edge'>
		<title>sample</title>
		<meta name='viewport' content='width=device-width, initial-scale=1'>
		<link rel='stylesheet' type='text/css' media='screen' href='main.css'>
		<script src='main.js'></script>
	</head>
	<body>
		<a href=/user/9"> GO TO: http://localhost:8080/user/9</a>
	</body>
	</html>`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

//GetUser export
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id == "101" {
		u := []models.User{}
		if err := uc.session.DB("test").C("users1").Find(bson.M{}).All(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		for _, v := range u {
			uj, _ := json.Marshal(v)
			fmt.Fprintf(w, "%s\n", uj)
		}
	} else if bson.IsObjectIdHex(id) {
		fmt.Println(id)
		oid := bson.ObjectIdHex(id)
		fmt.Println(oid)
		u := models.User{}
		if err := uc.session.DB("test").C("users1").FindId(oid).One(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", uj)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

//CreateUser export
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("test").C("users1").Insert(u)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteUser export
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("test").C("users1").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deleted user", oid, "\n")
}

//UpdateUser export
func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// id := p.ByName("id")
	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// // oid := bson.ObjectIdHex(id)
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	// uid := bson.M{"_id": u.Id}
	if err := uc.session.DB("test").C("users1").Update(u.Id, bson.M{"name": u.Name, "gender": u.Gender, "age": u.Age}); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// if err := uc.session.DB("test").C("users").Update(uid, bson.M{"name": u.Name, "age": u.Age}); err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w, "updated user:", u.Id, "\n")
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
