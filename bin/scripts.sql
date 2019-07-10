DROP TABLE IF EXISTS products CASCADE;
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    category_id INT,
    name VARCHAR,
    body TEXT,
    image VARCHAR,
    price FLOAT,
    created_at INT NOT NULL default date_part('epoch', now()),
    updated_at INT NOT NULL default date_part('epoch', now()),
    features JSONB,
    is_active BOOL
);