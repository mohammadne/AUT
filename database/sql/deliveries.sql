CREATE TABLE deliveries(
    id INT AUTO_INCREMENT PRIMARY KEY,

    orders_id INT,
    delivery_company_id INT,
    FOREIGN KEY(orders_id) REFERENCES orders(id),
    FOREIGN KEY(delivery_company_id) REFERENCES delivery_companies(id)
);

INSERT INTO deliveries (orders_id, delivery_company_id) VALUES
    (1, 1);