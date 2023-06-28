INSERT INTO customer ("name", email, "password", created_at, updated_at) VALUES('Muhammad Angga Ardhinata', 'ardhinataangga@gmail.com', 'angga123', 'now()', 'now()');
INSERT INTO categories ("name", created_at, updated_at) VALUES('electronic', 'now()', 'now()');
INSERT INTO product (name, category_id, price, description, image, created_at, updated_at) VALUES('Macbook Pro M2', 1, 21499000, 'MacBook Pro 13 inci lebih andal dari sebelumnya. Bertenaga super berkat chip M2 generasi berikutnya, MacBook Pro ini merupakan laptop pro Apple yang paling portabel, dengan kekuatan baterai hingga 20 jam', 'ini gambar', 'now()', 'now()');

INSERT INTO orders (cust_id, product_id, status, discount_code, created_at, updated_at) VALUES(1, 1, 'pending', 'IC015', 'now()', 'now()');