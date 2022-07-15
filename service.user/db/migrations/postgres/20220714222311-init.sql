-- +migrate Up
-- Create tables and relationships

CREATE TYPE "user_service"."permission_ref" AS ENUM (
  'super_admin',
  'admin'
);

CREATE TABLE "user_service"."user" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (public.uuid_generate_v4()),
  "role_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "user_service"."role" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (public.uuid_generate_v4()),
  "slug" text UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "user_service"."role_permission" (
  "id" uuid NOT NULL DEFAULT (public.uuid_generate_v4()),
  "role_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp,
  "deleted_at" timestamp,
  PRIMARY KEY ("id", "role_id", "permission_id")
);

CREATE TABLE "user_service"."permission" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (public.uuid_generate_v4()),
  "slug" user_service.permission_ref UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "user_service"."user" ADD FOREIGN KEY ("role_id") REFERENCES "user_service"."role" ("id");

ALTER TABLE "user_service"."role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "user_service"."role" ("id");

ALTER TABLE "user_service"."role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "user_service"."permission" ("id");

-- Seed values
INSERT INTO "user_service"."role" ("slug") VALUES('customer');

-- +migrate Down
-- Drop table & relationships
DROP TABLE "user_service"."role_permission" CASCADE;

DROP TABLE "user_service"."permission" CASCADE;

DROP TABLE "user_service"."role" CASCADE;

DROP TABLE "user_service"."user" CASCADE;

DROP TYPE "user_service"."permission_ref" CASCADE;
