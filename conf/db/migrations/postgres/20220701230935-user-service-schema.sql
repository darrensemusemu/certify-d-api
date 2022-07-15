
-- +migrate Up
-- add user_service user and schema
CREATE USER user_service WITH PASSWORD 'user_service';
CREATE SCHEMA user_service AUTHORIZATION user_service;
alter user user_service set search_path to 'user_service';

-- +migrate Down
-- remove user_service customer user and schema
DROP SCHEMA IF EXISTS user_service;
DROP USER IF EXISTS user_service;
