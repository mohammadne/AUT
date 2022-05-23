-- Show number of authors of each book

SELECT books.name AS 'book name', COUNT(*) AS number
FROM books
INNER JOIN books_users ON books.id = books_users.book_id
GROUP BY books.name;
