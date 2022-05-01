CREATE TABLE books_users(
    id INT AUTO_INCREMENT PRIMARY KEY,

    book_id INT,
    author_id INT,
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(author_id) REFERENCES users(id)
)

-- Docker in action -> Jeffrey Nickoloff, Stephen Kuenzli
-- Kubernetes in action -> Marko Luksa
-- Linux in action -> David Clinton

-- Computer System Architecture -> John Hennessy
-- Algorithm Design ->  Thomas H.Cormen, Charles E.Leiserson, Ronald L.Rivest, Clifford Stein

-- What is the Cloud -> Bill Laberis
-- Cloud Computing -> Titus Winters, Tom Manshreck, Hyrum Wright
-- Database Internals -> Alex Petrov

-- Complete Kubernetes Stack -> Jonathan Baier
-- SQL Fundamentals -> Anil Batra

INSERT INTO books_categories(book_id, author_id) VALUES
    (1, 2), (1, 3),
    (2, 4),
    (3, 5),
    (4, 6),
    (5, 7), (5, 8), (5, 9), (5, 10),
    (6, 11),
    (7, 12), (7, 13), (7, 14),
    (8, 15),
    (9, 16),
    (10, 17);
