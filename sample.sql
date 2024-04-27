-- Create table
CREATE TABLE IF NOT EXISTS Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50),
    age INT
);

-- Insert sample data
INSERT INTO Users (name, email, age) VALUES 
('John Doe', 'john@example.com', 30),
('Jane Smith', 'jane@example.com', 25),
('Michael Johnson', 'michael@example.com', 35);