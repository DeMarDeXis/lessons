ALTER TABLE lessons DROP COLUMN IF EXISTS status;

DROP TYPE IF EXISTS status_enum;

DROP TABLE IF EXISTS teachers_checklist;