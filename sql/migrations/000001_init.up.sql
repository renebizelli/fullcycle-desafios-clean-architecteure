CREATE TABLE IF NOT EXISTS  orders (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    price decimal(10,2) NOT NULL,
    tax decimal(10,2) NOT NULL,
    final_price decimal(10,2) NOT NULL
)
