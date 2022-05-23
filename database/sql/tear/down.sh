#! /bin/bash

mysql -p bookman <<'EOF'
DROP TABLE order_items;
DROP TABLE payments;
DROP TABLE deliveries;

DROP TABLE books_categories;
DROP TABLE books_users;
DROP TABLE comments;
DROP TABLE orders;

DROP TABLE users;
DROP TABLE books;

DROP TABLE roles;
DROP TABLE contact_informations;
DROP TABLE categories;
DROP TABLE publishers;
DROP TABLE languages;
DROP TABLE delivery_companies;
DROP TABLE payment_statuses;
EOF
