CREATE TABLE orders(
    id INT AUTO_INCREMENT PRIMARY KEY,
    count INT,

    user_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
