-- DB作成
CREATE DATABASE gohub_dev; 

-- 作成したDBへ切り替え
\c gohub_dev

-- スキーマ作成
CREATE SCHEMA gohub_dev_schema;

-- ロールの作成
CREATE ROLE gohub WITH LOGIN PASSWORD 'gohub_dev_password';

-- 権限追加
GRANT ALL PRIVILEGES ON SCHEMA gohub_dev_schema, table_name TO gohub;
