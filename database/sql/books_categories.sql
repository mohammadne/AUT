CREATE TABLE books_categories(
    id INT AUTO_INCREMENT PRIMARY KEY,

    book_id INT,
    category_id INT,
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(category_id) REFERENCES categories(id)
);

INSERT INTO books_categories(book_id, category_id) VALUES
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
