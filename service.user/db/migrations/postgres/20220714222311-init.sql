-- +migrate Up
-- Create tables and relationships
CREATE TYPE "user_service"."permission_ref" AS ENUM (
  'super_admin',
  'admin'
);

CREATE TABLE "user_service"."user" (
  "id" uuid PRIMARY KEY NOT NULL,
  "role_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "user_service"."role" (
  "id" uuid PRIMARY KEY NOT NULL,
  "slug" text UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "user_service"."role_permission" (
  "id" uuid,
  "role_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  PRIMARY KEY ("id", "role_id", "permission_id")
);

CREATE TABLE "user_service"."permission" (
  "id" uuid PRIMARY KEY NOT NULL,
  "slug" user_service.permission_ref UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

ALTER TABLE "user_service"."user" ADD FOREIGN KEY ("role_id") REFERENCES "user_service"."role" ("id");

ALTER TABLE "user_service"."role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "user_service"."role" ("id");

ALTER TABLE "user_service"."role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "user_service"."permission" ("id");

-- +migrate Down
-- Drop table & relationships
DROP TABLE "user_service"."role_permission" CASCADE;

DROP TABLE "user_service"."permission" CASCADE;

DROP TABLE "user_service"."role" CASCADE;

DROP TABLE "user_service"."user" CASCADE;

DROP TYPE "user_service"."permission_ref" CASCADE;
