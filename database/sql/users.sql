CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    
    role_id INT, 
    contact_information_id INT, 
    FOREIGN KEY(role_id) REFERENCES roles(id),
    FOREIGN KEY(contact_information_id) REFERENCES contact_informations(id)
);

INSERT INTO orders (role_id, contact_information_id)
VALUES (3, 1),
       (2, 2),
       (1, 3),
       (3, 4),
       (2, 5);
