CREATE TABLE order_items(
    id INT AUTO_INCREMENT PRIMARY KEY,
    count INT,

    book_id INT,
    order_id INT,
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(order_id) REFERENCES orders(id)
);

INSERT INTO order_items (count, book_id, order_id) VALUES
    (1, 2, 1), (2, 5, 1)
    (1, 8, 2);