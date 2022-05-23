-- show authors of the book "Algorithm Design"

SELECT CONCAT(first_name,' ', last_name) AS author
FROM contact_informations
INNER JOIN users ON contact_informations.id = users.contact_information_id
INNER JOIN books_users ON users.id = books_users.author_id
INNER JOIN books ON books.id = books_users.book_id
WHERE books.name='Algorithm Design';
