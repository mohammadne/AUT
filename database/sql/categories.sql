CREATE TABLE categories(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);

INSERT INTO categories (name) 
VALUES ('network'), ('cloud'), ('linux'), ('database'), ('algorithm');
