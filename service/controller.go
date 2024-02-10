package service

import (
	"context"
	"encoding/json"
	"firstmongo/models"
	"firstmongo/mongodb"
	"fmt"
	"strconv"

	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	
)



func GetUserById(w http.ResponseWriter,r *http.Request){
	var user models.User
	v1:=mux.Vars(r)
	id:=v1["id"]
	num1,err:=strconv.Atoi(id)
	if err!=nil{
		fmt.Println(err)
	}
	
	fmt.Println("id  :",num1)
	filter:=bson.M{"_id":num1}
	fmt.Println(filter)
	collection:=mongodb.Mdb.Db.Database("gouser").Collection("user")
	rr:=collection.FindOne(context.TODO(),filter).Decode(&user)
	if rr!=nil{
		fmt.Println("in get :",rr)
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)

}
func Storedata(w http.ResponseWriter,r *http.Request){
	var user models.User
	err:=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		fmt.Println(err)
	}
	collection:=mongodb.Mdb.Db.Database("gouser").Collection("user")
	one,err:=collection.InsertOne(context.TODO(),user)
	fmt.Println("one >>>>:",one)
	if err!=nil{
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(user)

}
func Display(w http.ResponseWriter,r *http.Request){
	temp,err:=template.ParseFiles("templates/index.html")
	if err!=nil{
		fmt.Println(err)
		http.Error(w,"not respond",http.StatusInternalServerError)
		return

	}
	temp.Execute(w,nil)
	
}
func GetData(w http.ResponseWriter,r *http.Request){
	// var num string=r.FormValue("number")
	// n1,err:=strconv.Atoi(num)
	r.ParseForm()
	n1:=r.Form.Get("number")
	fmt.Println(n1)
	
}
func UpdateUser(w http.ResponseWriter,r *http.Request){
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	filter:=bson.M{"_id":user.Id}
	update:=bson.M{"$set":
	bson.M{"name":user.Name,"age":user.Age}}
	collection:=mongodb.Mdb.Db.Database("gouser").Collection("user")
	collection.UpdateOne(context.TODO(),filter,update)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(user)
}