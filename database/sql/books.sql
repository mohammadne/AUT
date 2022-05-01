CREATE TABLE books(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),

    publisher_id INT,
    language_id INT,
    FOREIGN KEY(publisher_id) REFERENCES publishers(id),
    FOREIGN KEY(language_id) REFERENCES languages(id)
)

INSERT INTO books (name, publisher_id, language_id) 
VALUES ('Docker in action', 1, 1), ('Kubernetes in action', 1, 1), ('Linux in action', 1, 3),
       ('Computer System Architecture', 2, 1), ('Algorithm Design', 2, 2),
       ('What is the Cloud', 3, 1), ('Cloud Computing', 3, 3), ('Database Internals', 3, 1),
       ('Complete Kubernetes Stack', 4, 1), ('SQL Fundamentals', 4, 2),;
