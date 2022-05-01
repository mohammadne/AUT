CREATE TABLE payment_statuses(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);

INSERT INTO payment_statuses (name) VALUES 
    ('success'), ('failed');
