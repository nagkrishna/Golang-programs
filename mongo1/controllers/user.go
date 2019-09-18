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
type OTP struct {
	Id  bson.ObjectId `json:"id" bson:"_id"`
	Otp int           `json:"otp" bson:"otp"`
}

var u models.Userinfo
var i models.UserInterest

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
	otp := uc.CreateOtp()
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
func (uc UserController) CreateOtp() int {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(100000)
	Id := bson.NewObjectId()
	otp := OTP{Id, random}
	uc.session.DB("test").C("OTP").Insert(otp)
	return random
}
