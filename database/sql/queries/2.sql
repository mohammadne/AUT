-- show books published by Manning (join version)

SELECT books.id, books.name
FROM books
LEFT JOIN publishers ON books.publisher_id = publishers.id
WHERE publishers.name='Manning';
