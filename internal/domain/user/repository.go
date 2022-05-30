package user

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `Userdb`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `Olegoni678`,
}

type Repository interface {
	FindAll() ([]User, error)
	FindById(id int64) (*User, error)
	CreateUser(user User) (int64, error)
	UpdateById(user User) (*User, error)
}

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	var users []User

	usersCol := sess.Collection("Users")
	err = usersCol.Find().All(&users)

	if err != nil {
		log.Fatal("usersCol.Find: ", err)
	}
	return users, nil
}

func (r *repository) FindById(id int64) (*User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var user User
	userColl := sess.Collection("Users")
	err = userColl.Find("Id", id).One(&user)

	if err != nil {
		log.Fatal("userCol.FindOne: ", err)
	}
	return &user, nil
}

func (r *repository) CreateUser(user User) (int64, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	_, err = sess.SQL().
		InsertInto("Users").
		Values(user).
		Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	}

	var lastIdUser User
	rows, err := sess.SQL().
		Query(`SELECT MAX("Id") FROM "Users"`)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	if !rows.Next() {
		log.Fatal("Expecting one row")
	}
	if err := rows.Scan(&lastIdUser.Id); err != nil {
		log.Fatal("Scan: ", err)
	}
	if err := rows.Close(); err != nil {
		log.Fatal("Close: ", err)
	}

	return lastIdUser.Id, nil
}

func (r *repository) UpdateById(user User) (*User, error) {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	UsersTable := sess.Collection("Users")
	res := UsersTable.Find(user.Id)

	FromUser, err := r.FindById(user.Id)
	if err != nil {
		fmt.Printf("Enter Id:")
	}
	fmt.Printf("User: %#v", FromUser)

	if user.Name != "" {
		FromUser.Name = user.Name
	}
	if user.Age >= 18 {
		FromUser.Age = user.Age
	}
	if user.City != "" {
		FromUser.City = user.City
	}
	if user.Country != "" {
		FromUser.Country = user.Country
	}

	// The result set is updated.
	if err := res.Update(FromUser); err != nil {
		fmt.Printf("Update: %v\n", err)
		fmt.Println("")
	}

	return FromUser, nil
}
