-- show books published by Manning (subquery version)

SELECT books.id, books.name
FROM books
WHERE publisher_id =
(
    SELECT id FROM publishers
    WHERE name='Manning'
)
