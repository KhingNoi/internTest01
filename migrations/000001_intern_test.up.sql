-- Create Products table
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

-- Create Users table
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50),
    avatar_url VARCHAR(255)
);


-- Create Orders table
CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50),
    total DECIMAL(10, 2),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- Create OrderDetails table
CREATE TABLE order_details (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    product_id INT,
    quantity INT,
    price DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);


-- Insert data into Products table
INSERT INTO products (id, name, description, price, region, weight, flavor_profile, grind_option, roast_level, image_url, stock)
VALUES
(1, 'Signature Blend', 'A rich, full-bodied coffee with notes of dark chocolate and black cherry. Grown on the slopes of a mist-covered mountain in Central America.', 12.99, 'Central America', 500, '["Dark Chocolate", "Black Cherry"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 3, 'https://iili.io/H8Y78Qt.webp', 5),
(2, 'Golden Sunrise', 'A smooth and bright coffee with a citrusy kick. Sourced from the rolling hills of Africa.', 10.99, 'Africa', 500, '["Citrus"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 2, 'https://iili.io/H8Y7WEg.webp', 9),
(3, 'Rainforest Rhapsody', 'An earthy and complex coffee with notes of toasted nuts and caramel. Sustainably grown in the rainforests of South America.', 14.99, 'South America', 500, '["Citrus"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 2, 'https://iili.io/H8Y7kTN.webp', 2),
(4, 'Harvest Moon', 'A smooth and earthy blend with notes of cocoa and hazelnut.', 9.99, 'Central America', 500, '["Cocoa", "Hazelnut"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 3, 'https://iili.io/H8Y7X4a.webp', 3),
(5, 'Wildfire', 'A bold and smoky blend with notes of dark chocolate and molasses.', 12.99, 'Africa', 500, '["Dark Chocolate", "Molasses"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 5, 'https://iili.io/H8Y7r4s.webp', 10),
(6, 'Walnut Wonder', 'A smooth and nutty coffee from the slopes of South America.', 9.99, 'South America', 500, '["Nutty", "Smooth"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 3, 'https://iili.io/H8Y7gGn.webp', 7),
(7, 'Breezy Beans', 'This coffee blend is made from beans sourced from Honduras and Costa Rica. It is a light roast coffee with a bright and citrusy flavor profile. It is perfect for pour-over and drip coffee brewing methods.', 11.99, 'Central America', 500, '["Citrusy"]', '["Whole Bean", "Filter"]', 1, 'https://iili.io/H8Y7lpV.webp', 6),
(8, 'Indo-Viet Roast', 'This coffee blend is made from beans sourced from Indonesia and Vietnam. It is a medium-dark roast coffee with a spicy and earthy flavor profile, with notes of cinnamon and clove. It is perfect for drip and French press brewing methods.', 13.99, 'Asia Pacific', 500, '["Spicy", "Earthy", "Cinnamon", "Clove"]', '["Whole Bean", "Filter", "French press"]', 4, 'https://iili.io/H8Y7wYv.webp', 8),
(9, 'Ethiopian Yirgacheffe', 'A light and fruity coffee with notes of blueberry and citrus. Grown in the highlands of Ethiopia, this coffee is sure to brighten up your morning.', 12.99, 'Africa', 500, '["Blueberry", "Citrus"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 2, 'https://iili.io/H8Y7VCF.webp', 4),
(10, 'Lazy Days', 'A medium-bodied coffee with a sweet and nutty flavor. Grown in the lush regions of Brazil, this coffee is perfect for a lazy afternoon.', 9.99, 'South America', 500, '["Hazelnut", "Chocolate", "Caramel"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 3, 'https://iili.io/H8Y7NvR.webp', 9),
(11, 'Andean Almond', 'A smooth and mellow coffee from the mountains of South America, with hints of almond and toffee.', 10.99, 'South America', 500, '["Almond", "Toffee"]', '["Whole Bean", "Cafetiere", "Filter"]', 3, 'https://iili.io/H8Y5Sgj.webp', 6),
(12, 'Savanna Noir', 'A bold and rich coffee from the African savanna, with notes of dark chocolate and blackcurrant.', 12.99, 'Africa', 500, '["Dark Chocolate", "Blackcurrant"]', '["Whole Bean", "Filter", "Espresso"]', 4, 'https://iili.io/H8Y7vjI.webp', 3),
(13, 'Coconut Kiss', 'A light and refreshing coffee from the shores of the Asia Pacific, with a subtle coconut flavor.', 9.99, 'Asia Pacific', 500, '["Coconut"]', '["Whole Bean", "Filter"]', 2, 'https://iili.io/H8Y7GQ1.webp', 10),
(14, 'Arabian Nights', 'A bold and spicy coffee from the Middle East, with notes of cardamom and cinnamon.', 13.99, 'Middle East', 500, '["Cardamom", "Cinnamon"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', NULL, 'https://iili.io/H8Y7ckQ.webp', 2),
(15, 'Midnight Mocha', 'Indulge in the rich, velvety flavors of this decadent mocha blend. Dark chocolate and espresso notes are combined with a touch of vanilla for a luxurious coffee experience.', 14.99, 'South America', 500, '["Dark Chocolate", "Espresso", "Vanilla"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 4, 'https://iili.io/H8Y7Opp.webp', 4),
(16, 'Himalayan Heights', 'Grown on the steep slopes of the Himalayan Mountains, this coffee is known for its bright acidity and delicate floral notes. This light roast is perfect for those who prefer a more delicate flavor profile.', 12.99, 'Asia Pacific', 500, '["Floral", "Citrus", "Honey"]', '["Whole Bean", "Cafetiere", "Filter", "Pour Over"]', 1, 'https://iili.io/H8Y7j3J.webp', 8),
(17, 'Sumatra Blend', 'Get your day started with the bold and earthy flavors of Sumatra. Grown on the lush tropical slopes of the Gayo Highlands, this coffee is known for its full body, low acidity, and notes of dark chocolate and smoke.', 8.99, 'Asia Pacific', 500, '["Dark Chocolate", "Smoke"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 5, 'https://iili.io/H8Y7UCX.webp', 5),
(18, 'Bali Bliss', 'Escape to the tropical paradise of Bali with this smooth and mellow blend. Grown on small family farms, this coffee is shade-grown and hand-picked for a rich and nuanced flavor profile. Notes of milk chocolate, hazelnut, and a hint of tropical fruit.', 11.99, 'Asia Pacific', 500, '["Milk Chocolate", "Hazelnut", "Tropical Fruit"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 3, 'https://iili.io/H8Y71TB.webp', 1),
(19, 'Central Perk', 'A medium roast coffee with a smooth, nutty flavor and a hint of caramel. Inspired by your favorite coffee shop!', 9.99, 'Central America', 500, '["Nutty", "Caramel"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 3, 'https://iili.io/H8Y7aYx.webp', 0),
(20, 'Chilean Charm', 'This coffee boasts a smooth and balanced flavor, with notes of chocolate, caramel, and a hint of fruit. It has a medium body and a subtle acidity that will leave you feeling refreshed and energized.', 12.99, 'South America', 500, '["Chocolate", "Caramel", "Fruit"]', '["Whole Bean", "Cafetiere", "Filter", "Espresso"]', 2, 'https://iili.io/H8Y7EhP.webp', 7);



-- Insert data into Users table with the same AvatarURL for all users
INSERT INTO users (username, first_name, last_name, email, phone, avatar_url) VALUES
('john_doe', 'John', 'Doe', 'john@example.com', '123-456-7890', 'https://api.iconify.design/subway:admin.svg?color=%232b4a9d'),
('jane_smith', 'Jane', 'Smith', 'jane@example.com', '987-654-3210', 'https://api.iconify.design/subway:admin.svg?color=%232b4a9d'),
('alice_green', 'Alice', 'Green', 'alice@example.com', '555-555-5555', 'https://api.iconify.design/subway:admin.svg?color=%232b4a9d'),
('bob_jones', 'Bob', 'Jones', 'bob@example.com', '111-222-3333', 'https://api.iconify.design/subway:admin.svg?color=%232b4a9d');


INSERT INTO orders (user_id, status, total) VALUES
(1, 'Pending', 25.99),
(2, 'Processing', 15.49),
(3, 'Completed', 30.99),
(4, 'Pending', 20.99);

INSERT INTO order_details (order_id, product_id, quantity, price) VALUES
(1, 1, 2, 19.98),
(1, 3, 1, 6.99),
(2, 2, 1, 10.99),
(3, 5, 3, 38.97),
(4, 4, 1, 9.99),
(4, 6, 2, 19.98);

