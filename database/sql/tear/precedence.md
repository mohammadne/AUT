# Precedence

- level of each group doesn't matter

## level 1

roles
contact_informations
categories
publishers
languages
delivery_companies
payment_statuses

## level 2

users (roles, contact_informations)
books (publisher, language)

## level 3

books_categories (books, categories)
books_users (books, users)
comments (users, books, languages)
orders (users)

## level 4

order_items (orders, books)
payments (payment_statuses, orders)
deliveries (orders, delivery_companies)
