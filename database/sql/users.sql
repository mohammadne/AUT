CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    
    role_id INT, 
    FOREIGN KEY(role_id) REFERENCES roles(id)
);
