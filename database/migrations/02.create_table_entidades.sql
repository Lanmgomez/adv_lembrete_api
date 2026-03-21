/* 02 - Criar Tabela de Entidades */

CREATE TABLE IF NOT EXISTS entidades (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nome_entidade VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

/* mock data */
INSERT INTO entidades (nome_entidade) VALUES ('Câmara de Tabira');

SELECT * FROM entidades;


