INSERT INTO customer ("name", email, "password", created_at, updated_at) VALUES('muslih', 'muslihfreshfreuit@gmail.com', 'angga123', 'now()', 'now()');
INSERT INTO categories ("name", created_at, updated_at) VALUES('electronic', 'now()', 'now()');
INSERT INTO product (name, category_id, price, description, image, created_at, updated_at) VALUES('Mixue', 1, 5000, 'Dimana - mana ada mixue', 'ini gambar', 'now()', 'now()');
INSERT INTO orders (cust_id, product_id, status, discount_code, created_at, updated_at) VALUES(2, 1, 'pending', 'IC015', 'now()', 'now()');