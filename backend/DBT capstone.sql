DROP TABLE IF EXISTS "connections";
DROP TABLE IF EXISTS "message_providers";
DROP TABLE IF EXISTS "projects";
DROP TABLE IF EXISTS "users";


CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "email" varchar UNIQUE,
  "username" varchar UNIQUE,
  "password" varchar,
  "name" varchar
);


CREATE TABLE "projects" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "name" varchar
);

CREATE TABLE "message_providers" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "provider_type" varchar,
  "provider_label" varchar,
  "webhook" varchar
);


CREATE TABLE "connections" (
  "id" varchar PRIMARY KEY,
  "project_id" varchar unique,
  "message_provider_id" varchar unique
);

ALTER TABLE "connections" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "projects" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "connections" ADD FOREIGN KEY ("message_provider_id") REFERENCES "message_providers" ("id");

ALTER TABLE "message_providers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");




select * from users;
select * from connections;
select * from message_providers;
select * from projects;