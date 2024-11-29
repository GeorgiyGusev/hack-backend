-- Удаление индексов таблицы organizations
DROP INDEX IF EXISTS idx_organizations_status;
DROP INDEX IF EXISTS idx_organizations_location;

-- Удаление таблицы organizations
DROP TABLE IF EXISTS organizations;