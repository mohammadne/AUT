CREATE TABLE contact_informations(
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    phone VARCHAR(20),
    email VARCHAR(50),
    gender TinyINT,
    address VARCHAR(100),
);

INSERT INTO contact_informations (first_name, last_name, phone, email, gender, address) VALUES
    ('Mohammad', 'Nasr Esfahani', '09224505364', 'mohammadne@mail.ru', 1, 'Tehran'),
    
    ('Jeffrey', 'Nickoloff', '(666) 541-9578', 'jeffrey.nickoloff@email.com', 1, 'Ridgefield, Connecticut(CT), 06877'),
    ('Stephen', 'Kuenzli', '(907) 466-1659', 'stephen.kuenzli@email.com', 1, 'Centralia, Kansas(KS), 66415'),  
    ('Marko', 'Luksa', '(237) 533-5995', 'marko.luksa@email.com', 1, 'Cape Fair, Missouri(MO), 65624'),
    ('David', 'Clinton', '(276) 617-7138', 'david.clinton@email.com', 1, 'El Paso, Texas(TX), 79936'),
    ('John', 'Hennessy', '(434) 402-7985', 'john.hennessy@email.com', 1, 'Tracy, Minnesota(MN), 56175'),
    ('Thomas', 'H.Cormen', '(753) 365-1427', 'thomas.cormen@email.com', 1, 'Redlands, California(CA), 92374'),
    ('Charles', 'E.Leiserson', '(288) 512-9584', 'charles.leiserson@email.com', 1, 'Hudson, North Carolina(NC), 28638'),
    ('Ronald', 'L.Rivest', '(789) 767-0008', 'ronald.rivest@email.com', 1, 'Yukon, Oklahoma(OK), 73099'),
    ('Clifford', 'Stein', '(668) 392-2256', 'clifford.stein@email.com', 1, 'Westmont, Illinois(IL), 60559'),
    ('Bill', 'Laberis', '(985) 739-5169', 'bill.laberis@email.com', 1, 'Fountain Valley, California(CA), 92708'),
    ('Titus', 'Winters', '(935) 449-2497', 'titus.winters@email.com', 1, 'Alexandria, Virginia(VA), 22315'),
    ('Tom', 'Manshreck', '(982) 853-7713', 'tom.manshreck@email.com', 1, 'San Marcos, California(CA), 92069'),
    ('Hyrum', 'Wright', '(342) 847-0042', 'hyrum.wright@email.com', 0, 'Brillion, Wisconsin(WI), 54110'),
    ('Alex', 'Petrov', '(807) 633-6428', 'alex.petrov@email.com', 1, 'Jewett City, Connecticut(CT), 06351'),
    ('Jonathan', 'Baier', '(587) 987-6195', 'jonathan.baier@email.com', 1, 'Caldwell, Idaho(ID), 83607'),
    ('Anil', 'Batra', '(627) 998-7443', 'anil.batra@email.com', 1, 'Mc Lain, Mississippi(MS), 39456')

    ('ali', 'tavakoli', '09204545667', 'ali.tavakoli@email.com', 1, 'Tehran, razmandegan'),
    ('reza', 'ahmadi', '09204545467', 'reza.ahmadi@email.com', 1, 'Mashhad, Emam Reza'),
    ('narges', 'arasteh', '09135673444', 'narges.arasteh@email.com', 0, 'Tehran, zaferanie'),
    ('sara', 'manavi', '09201235667', 'sara.manavi@email.com', 0, 'Shiraz, hafez');
