CREATE TABLE IF NOT EXISTS  orders (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    price decimal NOT NULL,
    tax decimal NOT NULL,
    final_price decimal NOT NULL
)
