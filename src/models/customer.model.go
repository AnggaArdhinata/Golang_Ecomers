package models

import (
	"net/http"
	"time"

	db "github.com/AnggaArdhinata/indochat/src/configs"
)

type Customer struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"update_at"`
}

func StoreCustomer(name string, email string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO customer (name, email, password, created_at, updated_at) VALUES($1 , $2 , $3 , 'now()', 'now()');"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(name, email, password)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = Msg{"Customer Data Added"}

	return res, nil
}

func GetAllCustomer() (Response, error) {
	var obj Customer
	var arrobj []Customer
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM customer ORDER BY id DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Email, &obj.Password, &obj.Created_At, &obj.Updated_At)

		obj.Password = "secret"

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func UpdateCustomer(id int, name string, email string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE customer SET name= $1, email= $2, password= $3, updated_at='now()' WHERE id= $4"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, email, password, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteCustomer(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM customer WHERE id=$1"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
