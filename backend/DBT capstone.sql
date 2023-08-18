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
  "provider_id" varchar PRIMARY KEY,
  "project_id" varchar,
  "provider_type" varchar,
  "provider_label" varchar,
  "channel" varchar
);

ALTER TABLE "projects" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "message_providers" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

-- INSERT INTO "users" (id, email, password, name)
-- 		VALUES (1234, 'name@email.com', 'password', 'name');
		
select * from users;
select * from message_providers;
select * from projects;