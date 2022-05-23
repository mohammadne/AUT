#! /bin/bash

mysql -p bookman <<'EOF'
source /data/sql/roles.sql;
source /data/sql/contact_informations.sql;
source /data/sql/categories.sql;
source /data/sql/publishers.sql;
source /data/sql/languages.sql;
source /data/sql/delivery_companies.sql;
source /data/sql/payment_statuses.sql;

source /data/sql/users.sql;
source /data/sql/books.sql;

source /data/sql/books_categories.sql;
source /data/sql/books_users.sql;
source /data/sql/comments.sql;
source /data/sql/orders.sql;

source /data/sql/order_items.sql;
source /data/sql/payments.sql;
source /data/sql/deliveries.sql;
EOF
