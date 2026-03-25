/* 04 - Tabela de Lembretes */

ALTER TABLE lembretes
ADD COLUMN last_sent_at DATETIME NULL,
ADD COLUMN next_send_at DATETIME NULL;

