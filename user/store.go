package user

import (
	"database/sql"
	"fmt"
)
type Store interface {
	CreateUser(user *User) error
	GetUser() ([]*User, error)
	GetAuthenticationUser(email,password string) (*User, error)
	GetIdUser(userid string) (*User, error)
	DeleteIdUser(userid string) (*User, error)
	UpdateUser(email,address,password,userid string)  error
}
func (store *DbStore) UpdateUser(email,address,password,userid string) error{
	 sql := "update asset_user set email = ?, address = ?, password = ? where userid = ?"
	 stmt, err := store.Db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(email,address,password,userid)
	if err != nil {
		return err
	}

	return nil 
}
func (store *DbStore) DeleteIdUser(userid string) (*User, error) {
	user := &User{}
	err := store.Db.QueryRow("delete FROM asset_user where userid = ? ",userid).Scan(&user.UserID,&user.Email, &user.Address, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
type DbStore struct {
	Db *sql.DB
}
func Regis(storeDB *DbStore){
	store = storeDB
}
var store Store
func (store *DbStore) CreateUser(user *User) error {
	_, err := store.Db.Query(
		"INSERT INTO asset_user(email,address,password) VALUES (?,?,?)",
		user.Email, user.Address, user.Password)
	return err
}
func (store *DbStore) GetUser() ([]*User, error) {
	rows, err := store.Db.Query("SELECT * FROM asset_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	userList := []*User{}
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.UserID, &user.Email, &user.Address,&user.Password); 
		err != nil {
			
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}
func (store *DbStore) GetAuthenticationUser(email ,password string) (*User, error) {
	user := &User{}
	err := store.Db.QueryRow("SELECT email,password,userid,address FROM asset_user where email = ? and password = ?",email,password).Scan( &user.Email, &user.Password,&user.UserID,&user.Address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
func (store *DbStore) GetIdUser(userid string) (*User, error) {
	user := &User{}
	err := store.Db.QueryRow("SELECT userid,email,password,address FROM asset_user where userid = ? ",userid).Scan( &user.UserID,&user.Password,&user.Email,&user.Address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
