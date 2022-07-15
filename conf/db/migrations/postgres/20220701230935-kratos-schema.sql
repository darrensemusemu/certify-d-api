
-- +migrate Up
-- add kratos customer user and schema
CREATE USER kratos_customer WITH PASSWORD 'kratos_customer';
CREATE SCHEMA kratos_customer AUTHORIZATION kratos_customer;
alter user kratos_customer set search_path to 'kratos_customer';

-- +migrate Down
-- remove kratos customer user and schema
DROP SCHEMA IF EXISTS kratos_customer;
DROP USER IF EXISTS kratos_customer;
