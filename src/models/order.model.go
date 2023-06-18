package models

import (
	"fmt"
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

	fmt.Println("get order berhasil!")

	return res, nil
}

func PendingPayment() []string {

	var order OrderJoin
	var email []string

	con := db.CreateCon()

	sqlStatement := `SELECT o.id, c.id AS cust_id, c.name, c.email, p.name AS product ,o.ispaid FROM orders AS o
	INNER JOIN customer AS c ON o.cust_id = c.id
	INNER JOIN product AS p ON O.product_id = p.id WHERE o.ispaid=false`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.Cust_Id, &order.Cust_Name, &order.Cust_Email, &order.Product_Name, &order.IsPaid)

		email = append(email, order.Cust_Email)
	}

	return email
}
