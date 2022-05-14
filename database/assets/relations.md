# Book-Store Database Design

## Details

| entities             | records                                                                 |
|----------------------|-------------------------------------------------------------------------|
| roles                | id(PK), name                                                            |
| contact_informations | id(PK), phone_number, email                                             |
| categories           | id(PK), name, description                                               |
| publishers           | id(PK), name, location                                                  |
| languages            | id(PK), name                                                            |
| delivery_companies   | id(PK), name                                                            |
| payment_statuses     | id(PK), name                                                            |
| users                | id(PK), role_id(FK), contact_informations_id(FK), first_name, last_name |
| books                | id(PK), name, publisher_id(FK), language_id(FK)                         |
| books_categories     | id(PK), books_id(FK), categories_id(FK)                                 |
| books_users          | id(PK), books_id(FK), members_id(FK)                                    |
| comments             | id(PK), description, books_id(FK), member_id(FK)                        |
| orders               | id(PK), members_id(FK), payments_id(FK), data_time                      |
| order_items          | id(PK), books_id(FK), orders_id(FK), count                              |
| payments             | id(PK), amount, status                                                  |
| deliveries           | id(PK), books_id(FK), orders_id(FK)                                     |
