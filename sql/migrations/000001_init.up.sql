CREATE TABLE orders (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    price decimal NOT NULL,
    tax decimal NOT NULL,
    finalPrice decimal NOT NULL
)
