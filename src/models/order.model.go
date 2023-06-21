package models

import (
	"net/http"

	"time"

	db "github.com/AnggaArdhinata/indochat/src/configs"
)

type Order struct {
	Id            int       `json:"id"`
	Cust_Id       int       `json:"cust_id"`
	Product_Id    int       `json:"product_id"`
	IsPaid        bool      `json:"isPaid"`
	Status        string    `json:"status"`
	Discount_Code string    `json:"discount_code"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
}

type OrderJoin struct {
	Id            int       `json:"id"`
	Cust_Id       int       `json:"cust_id"`
	Cust_Name     string    `json:"customer_name"`
	Cust_Email    string    `json:"customer_email"`
	Product_Name  string    `json:"product_name"`
	Category      string    `json:"category"`
	Price         int       `json:"price"`
	Discount_Code string    `json:"discount_code"`
	Final_Price   int       `json:"final_price"`
	Status        string    `json:"status"`
	IsPaid        bool      `json:"isPaid"`
	Order_Date    string    `json:"order_date"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
}

func GetOrder() (Response, error) {
	var order OrderJoin
	var arrOrderJoin []OrderJoin
	var res Response

	con := db.CreateCon()

	sqlStatement := `SELECT 
	o.id AS order_id,
    cust_id,
    c.name AS customer_name,
	c.email AS customer_email,
    p.name AS product_name,
    cat.name AS category,
    price,
    o.discount_code,
    CASE
        WHEN discount_code = 'IC042'
        AND cat.name = 'electronic' THEN price - (price * 5 / 100)
        WHEN discount_code = 'IC003' THEN price - (price * 10 / 100)
        WHEN discount_code = 'IC015' 
        AND TO_CHAR(o.created_at, 'DY') = 'SAT' OR TO_CHAR(o.created_at, 'DY') = 'SUN' THEN price - (price * 10 / 100)
        ELSE price
    END AS final_price,
    o.status,
    ispaid,
    TO_CHAR(o.created_at, 'Day-Mon-YYYY') AS order_date,
	o.created_at,
    o.updated_at
	FROM orders AS o
    INNER JOIN customer AS c ON o.cust_id = c.id
    INNER JOIN product AS p ON O.product_id = p.id
    INNER JOIN categories AS cat ON p.category_id = cat.id
	ORDER BY o.id DESC`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Cust_Id, &order.Cust_Name, &order.Cust_Email, &order.Product_Name,
			&order.Category, &order.Price, &order.Discount_Code, &order.Final_Price,
			&order.Status, &order.IsPaid, &order.Order_Date, &order.Created_At, &order.Updated_At)

		if err != nil {
			return res, err
		}

		arrOrderJoin = append(arrOrderJoin, order)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrOrderJoin

	return res, nil
}

type OrderPayload struct {
	Email   string
	Name    string
	Product string
}

func PendingPayment() ([]OrderPayload, error) {

	var order OrderJoin
	var data []OrderPayload

	con := db.CreateCon()

	sqlStatement := `SELECT o.id, c.id AS cust_id, c.name, c.email, p.name AS product ,o.ispaid FROM orders AS o
	INNER JOIN customer AS c ON o.cust_id = c.id
	INNER JOIN product AS p ON O.product_id = p.id WHERE o.ispaid=false`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return data, nil
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Cust_Id, &order.Cust_Name, &order.Cust_Email, &order.Product_Name, &order.IsPaid)

		data = append(data, OrderPayload{Email: order.Cust_Email, Name: order.Cust_Name, Product: order.Product_Name})
	}

	return data, nil
}

func GenerateCsv() (Response, error) {
	var order OrderJoin
	var arrOrder []OrderJoin
	var res Response

	con := db.CreateCon()

	sqlStatement := `SELECT o.id AS order_id,
    c.name AS customer_name,
    TO_CHAR(o.created_at, 'Day-Mon-YYYY') AS order_date,
    (SELECT SUM(DISTINCT p.price) AS total),
    o.status
FROM orders AS o 
    INNER JOIN customer AS c ON o.cust_id = c.id
    INNER JOIN product AS p ON O.product_id = p.id
    GROUP BY o.id, c.name, p.id`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, nil
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Cust_Name, &order.Order_Date, &order.Price, &order.Status)

		arrOrder = append(arrOrder, order)
	}

	res.Status = http.StatusOK
	res.Message = "succes"
	res.Data = arrOrder

	return res, nil

}

func StoreOrder(cust_id int, product_id int, discount_code string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO orders (cust_id, product_id, status, discount_code, created_at, updated_at) VALUES($1, $2, 'pending', $3, 'now()', 'now()');"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt.Exec(cust_id, product_id, discount_code)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = Msg{"order has created"}
	return res, nil
}

func UpdateOrder(id int, cust_id int, product_id int, ispaid bool, status string, discount_order string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE orders SET cust_id= $1 , product_id= $2, ispaid= $3, status= $4, discount_code= $5, updated_at='now()' WHERE id= $6"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(cust_id, product_id, ispaid, status, discount_order, id)
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

func DeleteOrder(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM orders WHERE id=$1"

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
