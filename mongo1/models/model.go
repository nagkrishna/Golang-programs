package models

import "gopkg.in/mgo.v2/bson"

//Userinfo struct
type Userinfo struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Age  int           `json:"age" bson:"age"`
}

//UserInterest struct
type UserInterest struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Sport string        `json:"sport" bson:"sport"`
}
