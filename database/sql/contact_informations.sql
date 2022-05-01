CREATE TABLE contact_informations(
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    phone VARCHAR(20),
    email VARCHAR(50),
    gender TinyINT,
    address VARCHAR(100),
);

INSERT INTO contact_informations (first_name, last_name, phone, email, gender, address) 
VALUES ('ali', 'tavakoli', '09204545667', 'ali.tavakoli@email.com', 1, 'tehran, razmandegan'),
       ('reza', 'ahmadi', '09204545467', 'reza.ahmadi@email.com', 1, 'mashhad, emam'),
       ('mohammad', 'nasr', '09224505364', 'mohammad.nasr@email.com', 1, 'isfahan, monarjonban'),
       ('narges', 'arasteh', '09135673444', 'narges.arasteh@email.com', 0, 'tehran, zaferanie'),
       ('sara', 'manavi', '09201235667', 'sara.manavi@email.com', 0, 'shiraz, hafez');
