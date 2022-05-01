CREATE TABLE payments(
    id INT AUTO_INCREMENT PRIMARY KEY,

    orders_id INT,
    payment_status_id INT,
    FOREIGN KEY(orders_id) REFERENCES orders(id),
    FOREIGN KEY(payment_status_id) REFERENCES payment_statuses(id)
);

INSERT INTO payments (orders_id, payment_status_id) VALUES
    (1, 1), (2, 2);