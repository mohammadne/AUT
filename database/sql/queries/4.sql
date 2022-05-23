-- show how many books of each category we have

SELECT categories.name, COUNT(*) AS 'number'
FROM categories
LEFT JOIN books_categories ON categories.id = books_categories.category_id
LEFT JOIN books ON books.id = books_categories.book_id
GROUP BY categories.id;
