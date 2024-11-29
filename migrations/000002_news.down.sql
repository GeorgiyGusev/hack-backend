-- Удаление индексов таблицы news
DROP INDEX IF EXISTS idx_news_organization_id;
DROP INDEX IF EXISTS idx_news_publish_date;

-- Удаление таблицы news
DROP TABLE IF EXISTS news;