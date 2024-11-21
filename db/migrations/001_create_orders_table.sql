-- +goose Up
-- SQL para criar a tabela `orders`
CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY, 
    price DECIMAL(10, 2) NOT NULL,
    tax DECIMAL(10, 2) NOT NULL,
    final_price DECIMAL(10, 2) NOT NULL
);

-- Garantir que o root tenha privilégios
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root' WITH GRANT OPTION;
FLUSH PRIVILEGES;


-- +goose Down
-- SQL para desfazer a criação da tabela `orders`
DROP TABLE orders;
