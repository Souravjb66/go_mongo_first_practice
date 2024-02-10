package models
import(
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	
	Id  uint `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Age uint `json:"age" bason:"age"`
	
}