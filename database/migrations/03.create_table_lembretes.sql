/* 03 - Criar Tabela de Lembretes */

CREATE TABLE IF NOT EXISTS lembretes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    entidade_id BIGINT NOT NULL,
    nome_lembrete VARCHAR(255) NOT NULL,
    descricao TEXT,
    status ENUM('pendente', 'concluido', 'atrasado') DEFAULT 'pendente',
    data_vencimento DATE NOT NULL,
    dias_antecedencia INT DEFAULT 0,
    email_notificacao VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- 🔗 relacionamento
    CONSTRAINT fk_lembretes_entidade FOREIGN KEY (entidade_id) REFERENCES entidades (id) ON DELETE CASCADE
);