package mongodb

import (
	"context"
	
	// "encoding/json"
	// "firstmongo/mongodb"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type DBconnection struct{
	Db *mongo.Client
}
var Mdb DBconnection
func ConnectDb(){
	client,err:=mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		fmt.Println("from connect ->>>>..>",err)

	}
	
	Mdb.Db=client
	
}

// func(db *DBconnection)MyDbConnect(){  //so we use pointer because by paramter we change the instance only for this method not for outside
// 	connector:=ConnectDb()
// 	db.Db=connector

// }


