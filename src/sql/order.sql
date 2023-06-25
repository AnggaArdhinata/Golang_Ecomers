--DROP TABLE ORDER
-- DROP TABLE IF EXISTS orders;
--SELECT ALL ORDERS
-- SELECT * FROM orders
--INSERT ORDERS
-- INSERT INTO orders (cust_id, product_id, status, discount_code, created_at, updated_at) VALUES(5, 3, 'pending', 'IC015', 'now()', 'now()');
-- SELECT ORDERS WITH DISCOUNT
-- SELECT o.id AS order_id,
--     cust_id,
--     c.name AS customer_name,
--     p.name AS product_name,
--     cat.name AS category,
--     price,
--     o.discount_code,
--     CASE
--         WHEN discount_code = 'IC042'
--         AND cat.name = 'electronic' THEN price - (price * 5 / 100)
--         WHEN discount_code = 'IC003' THEN price - (price * 10 / 100)
--         WHEN discount_code = 'IC015' 
--         AND TO_CHAR(o.created_at, 'DY') = 'SAT' OR TO_CHAR(o.created_at, 'DY') = 'SUN' THEN price - (price * 10 / 100)
--         ELSE price
--     END AS final_price,
--     o.status,
--     ispaid,
--     TO_CHAR(o.created_at, 'Day-Mon-YYYY') AS order_date,
--     o.created_at,
--     o.updated_at
-- FROM orders AS o
--     INNER JOIN customer AS c ON o.cust_id = c.id
--     INNER JOIN product AS p ON O.product_id = p.id
--     INNER JOIN categories AS cat ON p.category_id = cat.id
--     ORDER BY o.id DESC
--SELECT PENDING ORDER
-- SELECT o.id, c.id AS cust_id, c.name, c.email, p.name AS product ,o.ispaid FROM orders AS o
-- INNER JOIN customer AS c ON o.cust_id = c.id
-- INNER JOIN product AS p ON O.product_id = p.id

--SELECT GENERATE CSV FILE
-- SELECT o.id AS order_id,
--     c.name AS customer_name,
--     TO_CHAR(o.created_at, 'Day-Mon-YYYY') AS order_date,
--     (SELECT SUM(DISTINCT p.price) AS total),
--     o.status
-- FROM orders AS o 
--     INNER JOIN customer AS c ON o.cust_id = c.id
--     INNER JOIN product AS p ON O.product_id = p.id
--     GROUP BY o.id, c.name, p.id
    
--UPDATE TABLE ORDER
-- UPDATE orders SET cust_id='1', product_id='2', ispaid= true, status='paid', discount_code='IC003', updated_at='now()' WHERE id= 4

-- SELECT SUM(p.price) AS total
-- FROM orders AS o
-- INNER JOIN product AS p ON O.product_id = p.id

-- SELECT o.id, c.id AS cust_id, c.name, c.email, p.name AS product, p.price AS price ,o.ispaid FROM orders AS o
-- 	INNER JOIN customer AS c ON o.cust_id = c.id
-- 	INNER JOIN product AS p ON O.product_id = p.id WHERE o.ispaid=false

-- SELECT * FROM orders;

-- UPDATE orders SET ispaid='true', status='paid' WHERE id=12


SELECT 
	o.id AS order_id,
    cust_id,
    c.name AS customer_name,
	c.email AS customer_email,
    p.name AS product_name,
    p.description,
    CASE
        WHEN discount_code = 'IC042'
        AND cat.name = 'electronic' THEN price - (price * 5 / 100)
        WHEN discount_code = 'IC003' THEN price - (price * 10 / 100)
        WHEN discount_code = 'IC015' 
        AND TO_CHAR(o.created_at, 'DY') = 'SAT' OR TO_CHAR(o.created_at, 'DY') = 'SUN' THEN price - (price * 10 / 100)
        ELSE price
    END AS final_price,
    status
	FROM orders AS o
    INNER JOIN customer AS c ON o.cust_id = c.id
    INNER JOIN product AS p ON O.product_id = p.id
    INNER JOIN categories AS cat ON p.category_id = cat.id
	ORDER BY o.id DESC