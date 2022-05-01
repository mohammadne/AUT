CREATE TABLE roles(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);

INSERT INTO roles (name)
VALUES ('admin'), ('author'), ('normal');
