package main
import(
	"context"
	"net/http"
	"github.com/gorilla/mux"
	"firstmongo/service"
	"firstmongo/mongodb"
)

func main(){
	mongodb.ConnectDb()
	Mydb:=mongodb.Mdb.Db
	defer Mydb.Disconnect(context.TODO())
	router:=mux.NewRouter()
	router.HandleFunc("/getUser/{id}",service.GetUserById).Methods("GET")
	router.HandleFunc("/storeUser",service.Storedata).Methods("POST")
	router.HandleFunc("/post",service.GetData).Methods("POST")
	router.HandleFunc("/index",service.Display).Methods("GET")
	router.HandleFunc("/updateUser",service.UpdateUser).Methods("POST")
	http.ListenAndServe(":8080",router)


}

