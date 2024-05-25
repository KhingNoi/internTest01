CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    region VARCHAR(100),
    weight DECIMAL(10, 2),
    flavor_profile JSON,
    grind_option JSON,
    roast_level INT NULL,
    image_url VARCHAR(255),
    stock INT
);