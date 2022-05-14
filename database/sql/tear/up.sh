#! /bin/bash

mysql -p <<'EOF'
use bookman;

-- LEVEL 1
source /data/sql/roles.sql;
source /data/sql/contact_informations.sql;
source /data/sql/categories.sql;
source /data/sql/publishers.sql;
source /data/sql/languages.sql;
source /data/sql/delivery_companies.sql;
source /data/sql/payment_statuses.sql;

-- LEVEL 2
source /data/sql/users.sql;
source /data/sql/books.sql;

-- LEVEL 3
source /data/sql/books_categories.sql;
source /data/sql/books_users.sql;
source /data/sql/comments.sql;
source /data/sql/orders.sql;

-- LEVEL 4
source /data/sql/order_items.sql;
source /data/sql/payments.sql;
source /data/sql/deliveries.sql;
EOF
