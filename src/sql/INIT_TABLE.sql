SET TIMEZONE ='Asia/Jakarta';
CREATE TABLE IF NOT EXISTS customer (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR,
    password VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP  
);

CREATE TABLE IF NOT EXISTS product (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    category_id INT,
    price INT,
    description TEXT,
    image VARCHAR,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT product_fk_1 FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    cust_id INT NOT NULL,
    product_id INT NOT NULL,
    ispaid BOOLEAN default false,
    status VARCHAR,
    discount_code VARCHAR,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT orderfk_1 FOREIGN KEY (cust_id) REFERENCES customer(id) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT orderfk_2 FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE SET NULL ON UPDATE CASCADE
);