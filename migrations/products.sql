CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    sku INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    category TEXT,
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (sku, name, category, price) VALUES 
    (1, 'BV Lean leather ankle boots', 'boots', 89000),
    (2, 'BV Lean leather ankle boots', 'boots', 99000),
    (3, 'Ashlington leather ankle boots', 'boots', 71000),
    (4, 'Naima embellished suede sandals', 'sandals', 79500),
    (5, 'Nathane leather sneakers', 'sneakers', 59000);