/* 01 - Criar Tabela de Usuários */

CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    user_type ENUM('owner', 'admin', 'user') NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

/* mock data */
SELECT * FROM users;

INSERT INTO users (username, email, password, user_type) 
VALUES ('islan gomes', 'islan_gomes@hotmail.com', '123456', 'owner');

