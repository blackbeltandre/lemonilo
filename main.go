package main

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"lemonilo/user"
	"flag"
	"fmt"

)

var(
	dbPort = "3306"
	port = "8088"
	dbUser = "root"
	dbPass = ""
	schema = "lemonilo"
	host = "localhost"
)

func init(){
 flag.StringVar(&dbPort,"dbport","3306","untuk db port")
 flag.StringVar(&port,"port","8088","untuk port")
 flag.StringVar(&dbUser,"dbuser","root","untuk user db")
 flag.StringVar(&dbPass,"dbpass","","untuk password db")
 flag.StringVar(&schema,"schema","lemonilo","untuk schema db")
 flag.StringVar(&host,"host","localhost","untuk host db")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()


	//CRUD
	r.HandleFunc("/user/get_user", user.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/update_user", user.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/delete_user", user.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user/get_id_user", user.GetIdUserHandler).Methods("GET")
	r.HandleFunc("/user/insert_user", user.CreateUserHandler).Methods("POST")


	//LOGIN
	r.HandleFunc("/user/get_authentication_user", user.GetAuthenticationUserHandler).Methods("GET")

	return r
}
func main() {
	flag.Parse()
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/lemonilo")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	storeUser := &user.DbStore{Db: db}
	user.Regis(storeUser)
	
	r := newRouter()
	fmt.Println("Starting server at port : ", port,"- Lemonilo Backend Active")
	http.ListenAndServe(fmt.Sprintf(":%v",port), r)

}
