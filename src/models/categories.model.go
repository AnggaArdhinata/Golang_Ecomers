package models

import (
	"net/http"
	"time"

	db "github.com/AnggaArdhinata/indochat/src/configs"
)

type Categories struct {
	Id         int       `json:"id"`
	Name       string    `json:"category"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"update_at"`
}

func StoreCategory(name string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `INSERT INTO categories (name, created_at, updated_at) VALUES ($1 , 'now()', 'now()');`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(name)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = Msg{"category added"}

	return res, nil
}

func GetAllCategory() (Response, error) {
	var obj Categories
	var arrobj []Categories
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM categories ORDER BY id DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Created_At, &obj.Updated_At)

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

func UpdateCategory(id int, name string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE categories SET name= $1, updated_at='now()' WHERE id= $2"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, id)
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

func DeleteCategory(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM categories WHERE id=$1"

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