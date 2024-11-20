-- +goose Up
-- SQL para criar a tabela `orders`
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    order_date DATETIME NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL
);

-- +goose Down
-- SQL para desfazer a criação da tabela `orders`
DROP TABLE orders;
