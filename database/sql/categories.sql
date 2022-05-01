CREATE TABLE categories(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);

INSERT INTO categories (name) 
VALUES ('action'), ('adventure'), ('classic'), ('horror'), ('mystery');
