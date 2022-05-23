CREATE TABLE comments(
    id INT AUTO_INCREMENT PRIMARY KEY,
    text VARCHAR(500),

    user_id INT, 
    book_id INT,
    language_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(language_id) REFERENCES languages(id)
);

INSERT INTO comments(text, user_id, book_id, language_id) VALUES
    ("best docker book that I have seen so far", 18, 1, 1),
    ("most complete algorithm book", 18, 5, 1),

    ("best book for learning sql", 19, 10, 1),
    ("کتاب عالی برای یادگیری دیتابیس", 19, 8, 2),

    ("المرحبا", 20, 2, 3),
    ("تبارک الله", 20, 9, 3),
    ("good cloud book", 20, 6, 1),
    ("my second good cloud book", 20, 7, 1),

    ("best computer system book", 21, 4, 1),
    ("best linux book", 21, 3, 1);