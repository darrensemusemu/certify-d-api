
-- +migrate Up
-- Create table and relationships
CREATE TABLE "upload_service"."store" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (public.uuid_generate_v4()),
  "ref" text NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "upload_service"."file" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (public.uuid_generate_v4()),
  "title" text NOT NULL,
  "url" text NOT NULL,
  "type" text NOT NULL,
  "number_of_pages" int DEFAULT 1,
  "store_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "upload_service"."file" ADD FOREIGN KEY ("store_id") REFERENCES "upload_service"."store" ("id");

-- +migrate Down
DROP TABLE  "upload_service"."file" CASCADE;

DROP TABLE  "upload_service"."store" CASCADE;