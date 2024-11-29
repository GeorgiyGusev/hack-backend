-- Оптимизированная таблица организаций
CREATE TABLE organizations
(
    id          UUID PRIMARY KEY,                     -- Уникальный идентификатор организации
    owner_id    UUID         NOT NULL,                -- Владелец организации
    photo_id    UUID         NULL,                    -- Фото организации (может быть NULL)
    title       VARCHAR(255) NOT NULL,                -- Название организации
    description TEXT,                                 -- Подробное описание организации
    phone       VARCHAR(20)  NOT NULL,                -- Контактный телефон
    email       VARCHAR(255) NOT NULL,                -- Электронная почта
    status      VARCHAR(50)  NOT NULL CHECK (
        status IN ('pending', 'approved', 'rejected') -- Статус организации
        ),
    longtitude  FLOAT        NOT NULL,                -- Географическая долгота
    latitude    FLOAT        NOT NULL,                -- Географическая широта
    created_at  TIMESTAMP DEFAULT NOW(),              -- Дата создания записи
    updated_at  TIMESTAMP DEFAULT NOW()               -- Дата последнего обновления записи
);

-- Индекс для ускорения поиска по статусу
CREATE INDEX idx_organizations_status ON organizations (status);

-- Индекс для быстрого поиска по географическому расположению
CREATE INDEX idx_organizations_location ON organizations (longtitude, latitude);

-- ЭТО НЕ ЭТИ ВАШИ ЧАТЫ ГПТ!! Я РИЛ НАПИСАЛ КОММЕНТЫ К МИГРАЦИИ )))