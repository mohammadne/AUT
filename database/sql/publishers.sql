CREATE TABLE publishers(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);

INSERT INTO publishers (name) 
VALUES ('Manning'), ('Pearson'), ('Oreilly'), ('Packt');
