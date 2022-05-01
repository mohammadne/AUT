CREATE TABLE books_users(
    id INT AUTO_INCREMENT PRIMARY KEY,

    book_id INT,
    user_id INT,
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
)

INSERT INTO books_categories(book_id, user_id) VALUES
    (1, 2), (1, 3),
    (2, 1), (2, 2), (2, 3),
    (3, 3),
    (4, 3),
    (5, 5),
    (6, 1), (6, 2),
    (7, 1), (7, 2),
    (8, 4),
    (9, 1), (9, 2), (9, 3),
    (10, 4);
