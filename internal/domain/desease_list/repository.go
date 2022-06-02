package desease_list

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

type DeseaseListRepository interface {
	FindAll() ([]DeseaseList, error)
	FindById(id int64) (*DeseaseList, error)
	CreateDesease(desease_list DeseaseList) (int64, error)
	//UpdateById(desease_list DeseaseList) (*DeseaseList, error)
}

type repository struct {
	// Some internal data
}

func NewDeseaseListRepository() DeseaseListRepository {
	return &repository{}
}

func (r *repository) FindAll() ([]DeseaseList, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	var users []DeseaseList

	usersCol := sess.Collection("Users")
	err = usersCol.Find().All(&users)

	if err != nil {
		log.Fatal("usersCol.Find: ", err)
	}
	return users, nil
}

func (r *repository) FindById(id int64) (*DeseaseList, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var user DeseaseList
	userColl := sess.Collection("Users")
	err = userColl.Find("Id", id).One(&user)

	if err != nil {
		log.Fatal("userCol.FindOne: ", err)
	}
	return &user, nil
}

func (r *repository) CreateDesease(desease_list DeseaseList) (int64, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	_, err = sess.SQL().
		InsertInto("Users").
		Values(desease_list).
		Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	}

	var lastIdUser DeseaseList
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

//func (r *repository) UpdateById(desease_list DeseaseList) (*DeseaseList, error) {
//
//	sess, err := postgresql.Open(settings)
//	if err != nil {
//		log.Fatal("Open: ", err)
//	}
//	defer sess.Close()
//
//	UsersTable := sess.Collection("Users")
//	res := UsersTable.Find(desease_list.Id)
//
//	FromDeseaseList, err := r.FindById(desease_list.Id)
//	if err != nil {
//		fmt.Printf("Enter Id:")
//	}
//	fmt.Printf("User: %#v", FromDeseaseList)
//
//	if desease_list.Name != "" {
//		FromDeseaseList.Name = desease_list.Name
//	}
//	if desease_list.Age >= 18 {
//		FromDeseaseList.Age = desease_list.Age
//	}
//	if desease_list.City != "" {
//		FromDeseaseList.City = desease_list.City
//	}
//	if desease_list.Country != "" {
//		FromDeseaseList.Country = desease_list.Country
//	}
//
//	// The result set is updated.
//	if err := res.Update(FromDeseaseList); err != nil {
//		fmt.Printf("Update: %v\n", err)
//		fmt.Println("")
//	}
//
//	return FromDeseaseList, nil
//}
