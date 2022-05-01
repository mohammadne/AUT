CREATE TABLE orders(
    id INT AUTO_INCREMENT PRIMARY KEY,

    user_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

INSERT INTO books (user_id) VALUES
    (19), (21);