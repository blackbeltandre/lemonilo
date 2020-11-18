18 Nov'2020
Andreas Sahabat Lumban Batu


This api will work with basic auth.

Golang :
go version go1.14.4 windows/amd64


Basic auth credential 
username : techlead
password : coding


port : 8088
segment : localhost / 127.0.0.1


Method this :
	//CRUD
	r.HandleFunc("/user/get_user", user.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/update_user", user.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/delete_user", user.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user/get_id_user", user.GetIdUserHandler).Methods("GET")
	r.HandleFunc("/user/insert_user", user.CreateUserHandler).Methods("POST")

	//LOGIN
	r.HandleFunc("/user/get_authentication_user", user.GetAuthenticationUserHandler).Methods("GET")




