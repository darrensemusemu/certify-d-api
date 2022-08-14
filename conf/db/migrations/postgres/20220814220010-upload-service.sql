
-- +migrate Up
-- add upload_service and schema
CREATE USER upload_service WITH PASSWORD 'upload_service';
CREATE SCHEMA upload_service AUTHORIZATION upload_service;
alter user upload_service set search_path to 'upload_service';

-- +migrate Down
-- remove upload_service amd schema
DROP SCHEMA IF EXISTS upload_service;
DROP USER IF EXISTS upload_service;
