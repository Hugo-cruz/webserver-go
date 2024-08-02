-- Create the Devices table
CREATE TABLE IF NOT EXISTS Devices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    creation_time DATETIME NOT NULL
);
-- Insert sample data into the Devices table
INSERT INTO Devices (name, brand, creation_time)
VALUES
    ('Device1', 'BrandA', '2024-01-01 10:00:00'),
    ('Device2', 'BrandB', '2024-02-01 11:00:00'),
    ('Device3', 'BrandC', '2024-03-01 12:00:00');
