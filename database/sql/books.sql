CREATE TABLE books(
    id INT AUTO_INCREMENT PRIMARY KEY,

    author INT,
    category_id INT,
    publisher_id INT,
    language_id INT,
    FOREIGN KEY(author) REFERENCES users(id),
    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(language_id) REFERENCES languages(id)
)
