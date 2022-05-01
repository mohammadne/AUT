CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    
    role_id INT, 
    contact_information_id INT, 
    FOREIGN KEY(role_id) REFERENCES roles(id),
    FOREIGN KEY(contact_information_id) REFERENCES contact_informations(id)
);

INSERT INTO orders (role_id, contact_information_id) VALUES 
    (1, 1),
    (2, 2), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7), (2, 8), (2, 9), (2, 10), (2, 11), (2, 12), (2, 13), (2, 14), (2, 15), (2, 16), (2, 17),
    (3, 18), (3, 19), (3, 20), (3, 21);
