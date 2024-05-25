CREATE TABLE Products (
    ProductID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255),
    Description TEXT,
    Price DECIMAL(10, 2),
    Region VARCHAR(100),
    Weight DECIMAL(10, 2),
    FlavorProfile JSON,
    GrindOption JSON,
    RoastLevel INT NULL,
    ImageURL VARCHAR(255),
    Stock INT
);