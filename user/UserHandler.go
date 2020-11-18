package user

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)
const USERNAME = "techlead"
const PASSWORD = "coding"

type Result struct {
	Message string      `json:"message"`
	Status  int         `:"status"`
	Result  interface{} `json:"result"`
}
//Checking authorize
func Auth(w http.ResponseWriter, r *http.Request) bool {
    username, password, ok := r.BasicAuth()
    if !ok {
        result := &Result{Result:"Credential is empty",Message:"Authorize Failed",Status:2}
        isEmpty, err := json.Marshal(result)
        if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
        }
        w.Header().Set("Content-Type", "application/json")
         w.Write(isEmpty)
        return false
    }
    isValid := (username == USERNAME) && (password == PASSWORD)
    if !isValid {
        result := &Result{Result:"Credential is not valid",Message:"Credential is not valid",Status:3}
        isFail, err := json.Marshal(result)
        if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
        }
        w.Header().Set("Content-Type", "application/json")
         w.Write(isFail)
        return false
    }
    return true
    }
type User struct {
	UserID string `json:"userid"`
	Email string `json:"email"`
	Address string `json:"address"`
	Password string `json:"password"`
}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r){ return }
	email,_ := r.URL.Query()["email"]
	password,_ := r.URL.Query()["password"]
	address,_ := r.URL.Query()["address"]
	userid,_ := r.URL.Query()["userid"]

	 err := store.UpdateUser(email[0],address[0],password[0],userid[0])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := &Result{Result:nil,Message:"Data Success Updated",Status:1}
	userListBytes, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userListBytes)
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r){ return }
	userid,ok := r.URL.Query()["userid"]
	if ok != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userList, err := store.DeleteIdUser(userid[0])
	result := &Result{Result:userList,Message:"Data Success Deleted",Status:1}
	userListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userListBytes)
}
func GetAuthenticationUserHandler(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r){ return }
	if store == nil{
		result := &Result{Result:"SOMETHING WRONG",Message:"PLEASE REGISTER DB TO MAIN CLASS !",Status:0}
		userListBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(userListBytes)
		return
	}
	email,ok := r.URL.Query()["email"]
	if ok != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	password, ok := r.URL.Query()["password"]
	if ok != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	} 
	userList, err := store.GetAuthenticationUser(email[0],password[0])
	result := &Result{Result:userList,Message:"Success Display Data",Status:1}
	userListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userListBytes)
}
func GetIdUserHandler(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r){ return }
	if store == nil{
		result := &Result{Result:"SOMETHING WRONG",Message:"PLEASE REGISTER DB TO MAIN CLASS !",Status:0}
		userListBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(userListBytes)
		return
	}
	userid,ok := r.URL.Query()["userid"]
	fmt.Println(userid)
	if ok != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userList, err := store.GetIdUser(userid[0])
	result := &Result{Result:userList,Message:"Success Display Data",Status:1}
	userListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userListBytes)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r){ return }
	if store == nil{
		result := &Result{Result:"SOMETHING WRONG",Message:"PLEASE REGISTER DB TO MAIN CLASS !",Status:0}
		userListBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(userListBytes)
		return
	}
	userList, err := store.GetUser()
	result := &Result{Result:userList,Message:"Success Display Data",Status:1}
	userListBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userListBytes)
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r){ return }
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := User{}
	user.UserID = r.Form.Get("userid")
	user.Email = r.Form.Get("email")
	user.Password = r.Form.Get("password")
	user.Address = r.Form.Get("address")
	
	err = store.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
	}
	result := &Result{Result:"Process Success",Message:"Data Added",Status:1}
	status_ok, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(status_ok)
	// w.WriteHeader(http.StatusFound)
}
