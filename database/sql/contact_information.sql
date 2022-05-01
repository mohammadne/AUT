CREATE TABLE contact_information(
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    phone VARCHAR(20),
    email VARCHAR(50),
    gender TinyINT,
    address VARCHAR(100),
);
