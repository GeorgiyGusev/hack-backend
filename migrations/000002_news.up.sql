-- Таблица новостей
CREATE TABLE news
(
    id              UUID PRIMARY KEY,        -- Уникальный идентификатор новости
    organization_id UUID         NOT NULL,   -- Ссылка на организацию
    title           VARCHAR(255) NOT NULL,   -- Заголовок новости
    description     TEXT         NOT NULL,   -- Полное описание новости
    publish_date    TIMESTAMP DEFAULT NOW(), -- Дата публикации
    media           UUID[],                  -- Список ID медиа файлов
    created_at      TIMESTAMP DEFAULT NOW(), -- Дата создания записи
    updated_at      TIMESTAMP DEFAULT NOW(), -- Дата последнего обновления записи
    FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE
);

-- Индекс для ускорения поиска новостей по организации
CREATE INDEX idx_news_organization_id ON news (organization_id);

-- Индекс для сортировки новостей по дате публикации
CREATE INDEX idx_news_publish_date ON news (publish_date);