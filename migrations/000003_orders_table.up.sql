CREATE TABLE Orders (
    OrderID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT,
    OrderDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Status VARCHAR(50),
    Total DECIMAL(10, 2),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);