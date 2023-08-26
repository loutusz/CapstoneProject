DROP TABLE IF EXISTS "connections";
DROP TABLE IF EXISTS "message_providers";
DROP TABLE IF EXISTS "projects";
DROP TABLE IF EXISTS "users";


CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "email" varchar UNIQUE,
  "username" varchar UNIQUE,
  "password" varchar,
  "name" varchar
);


CREATE TABLE "projects" (
  "project_id" varchar PRIMARY KEY,
  "project_user_id" varchar,
  "name" varchar
);

CREATE TABLE "message_providers" (
  "message_provider_id" varchar PRIMARY KEY,
  "message_provider_user_id" varchar,
  "provider_type" varchar,
  "provider_label" varchar,
  "webhook" varchar
);


CREATE TABLE "connections" (
  "connections_id" varchar PRIMARY KEY,
  "connections_project_id" varchar unique,
  "connection_message_provider_id" varchar unique
);

ALTER TABLE "connections" ADD FOREIGN KEY ("connections_project_id") REFERENCES "projects" ("project_id");

ALTER TABLE "projects" ADD FOREIGN KEY ("project_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "connections" ADD FOREIGN KEY ("connection_message_provider_id") REFERENCES "message_providers" ("message_provider_id");

ALTER TABLE "message_providers" ADD FOREIGN KEY ("message_provider_user_id") REFERENCES "users" ("user_id");




select * from users;
select * from connections;
select * from message_providers;
select * from projects;