package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"math/rand"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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

type User struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Age   int           `json:"age" bson:"age"`
	Sport string        `json:"sport" bson:"sport"`
}
type Otp struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Otp  int           `json:"otp" bson:"otp"`
}

var u models.Userinfo
var i models.UserInterest
var o Otp

//GetUser export
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	otp := p.ByName("otp")
	if err := uc.session.DB("test").C("otp").Find(bson.M{"otp": otp}).One(&o); err != nil {
		fmt.Println(err)
	}
	if o.Name == name {
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
		<h1>You are Successfully Logged In.</h1>
		//<a href=/user/9"> GO TO: http://localhost:8080/user/9</a>
	</body>
	</html>`
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(s))
	}
}

//CreateUserinfo export
func (uc UserController) CreateUserinfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	var wg sync.WaitGroup
	wg.Add(2)
	go uc.CreateUser1(name, &wg)
	go uc.CreateUser2(name, &wg)
	wg.Wait()
	Id := bson.NewObjectId()
	s := User{Id, u.Name, u.Age, i.Sport}
	uc.session.DB("test").C("user1").Insert(s)
	otp := uc.CreateOtp(u.Name)
	uj, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
	fmt.Println("OTP is: ", otp)
}

//CreateUser1 export
func (uc UserController) CreateUser1(name string, wg *sync.WaitGroup) {
	if err := uc.session.DB("test").C("userinfo").Find(bson.M{"name": name}).One(&u); err != nil {
		fmt.Println("Error!")
	}
	wg.Done()
}

//CreateUser2 export
func (uc UserController) CreateUser2(name string, wg *sync.WaitGroup) {
	if err := uc.session.DB("test").C("userinterests").Find(bson.M{"name": name}).One(&i); err != nil {
		fmt.Println("Error!")
	}
	wg.Done()
}

//CreateOtp export
func (uc UserController) CreateOtp(name string) int {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(100000)
	Id := bson.NewObjectId()
	otp := Otp{Id, name, random}
	uc.session.DB("test").C("otp").Insert(otp)
	return random
}
