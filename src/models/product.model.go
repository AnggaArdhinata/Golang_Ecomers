package models

import (
	"net/http"
	"time"

	db "github.com/AnggaArdhinata/indochat/src/configs"
)

type Product struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Category_Id int          `json:"category_id"`
	Price       int          `json:"price"`
	Description string       `json:"description"`
	Image       string       `json:"image"`
	Created_At  time.Time    `json:"created_at"`
	Updated_At  time.Time    `json:"updated_at"`
}

func StoreProduct(name string, category_id int, price int, description string, image string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO product (name, category_id, price, description, image, created_at, updated_at) VALUES($1, $2, $3, $4, $5, 'now()', 'now()')"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(name, category_id, price, description, image)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = Msg{"successfully add product"}

	return res, nil
}

func GetAllProduct() (Response, error) {
	var obj Product
	var arrobj []Product
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM product ORDER BY id DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Category_Id, &obj.Price, &obj.Description, &obj.Image, &obj.Created_At, &obj.Updated_At)

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

func UpdateProduct(id int, name string, category_id int, price int, description string, image string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE product SET name= $1, category_id= $2, price= $3, description= $4, image= $5, updated_at='now()' WHERE id= $6"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, category_id, price, description, image, id)
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

func DeleteProduct(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM product WHERE id=$1"

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
