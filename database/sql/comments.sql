CREATE TABLE comments(
    id INT AUTO_INCREMENT PRIMARY KEY,
    text VARCHAR(500),

    user_id INT, 
    book_id INT,
    language_id INT,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(language_id) REFERENCES languages(id)
)
